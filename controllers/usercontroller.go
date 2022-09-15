package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"unified-hiring-portal/database"
	"unified-hiring-portal/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllJobsPostedByUser(c *fiber.Ctx) error {

	uid, err := retrieve_id(c)

	if uid == 0 {
		return err
	}

	var employer models.User
	err = database.DB.Where("ID = ?", uid).Preload("JobsPosted").Find(&employer).Error

	if err != nil {
		return c.JSON(fiber.Map{
			"Error": err,
		})
	}

	return c.JSON(employer)

}

func AddJobPost(c *fiber.Ctx) error {
	uid, err := retrieve_id(c)

	if uid == 0 {
		return err
	}

	var job models.Job

	err = c.BodyParser(&job)

	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request Failed"})
		return err
	}

	job.EmployerID = uid

	fmt.Println(job)
	err = database.DB.Create(&job).Error

	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create a Job"})
		return err
	}

	return c.JSON(fiber.Map{
		"Message": "Successfully created the Job",
		"JobID":   job,
	})

}

func UpdateJob(c *fiber.Ctx) error {
	uid, err := retrieve_id(c)

	if uid == 0 {
		return err
	}

	jobid := c.Params("jobid")

	if jobid == "" {
		return c.JSON(fiber.Map{
			"Error": "Empty JobID, Cannot Update the Job",
		})
	}

	temp, _ := strconv.ParseUint(jobid, 10, 64)

	jid := uint(temp)

	var updated_job models.Job

	err = c.BodyParser(&updated_job)

	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request Failed"})
		return err
	}

	job := models.Job{}

	err = database.DB.Where("id = ? AND employer_id = ?", jid, uid).First(&job).Error

	if err != nil {
		return c.JSON(fiber.Map{
			"Error": err,
		})
	}

	updated_job.ID = jid

	err = database.DB.Model(job).Updates(updated_job).Error

	if err != nil {
		return c.JSON(fiber.Map{
			"Error": err,
		})
	}

	return c.JSON(fiber.Map{
		"Message": "Successfully Updated the Job",
	})
}
func DeleteJob(c *fiber.Ctx) error {

	uid, err := retrieve_id(c)

	if uid == 0 {
		return err
	}

	jobid := c.Params("jobid")

	if jobid == "" {
		return c.JSON(fiber.Map{
			"Error": "Empty JobID, Cannot delete the Job",
		})
	}

	temp, _ := strconv.ParseUint(jobid, 10, 64)

	jid := uint(temp)

	job := models.Job{}

	err = database.DB.Where("id = ? AND employer_id = ?", jid, uid).First(&job).Error

	if err != nil {
		return c.JSON(fiber.Map{
			"Error": err,
		})
	}

	err = database.DB.Delete(&models.Job{}, jid).Error

	if err != nil {
		return c.JSON(fiber.Map{
			"Error": err,
		})
	}

	return c.JSON(fiber.Map{
		"Message": "Successfully Deleted the Job",
	})
}
