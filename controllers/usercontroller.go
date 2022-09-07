package controllers

import (
	"fmt"
	"net/http"
	"unified-hiring-portal/database"
	"unified-hiring-portal/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllJobsPostedByUser(c *fiber.Ctx) error {

	email, err := retrieve_email(c)

	if email == "" {
		return err
	}
	// var job []models.Job
	var employerID uint
	result := database.DB.Select("ID").Where("Email = ?", email).First(&employerID)

	return c.JSON(result.RowsAffected)

}

func AddJobPost(c *fiber.Ctx) error {
	email, err := retrieve_email(c)

	if email == "" {
		return err
	}

	var employer models.User
	database.DB.Select("id").Where("email = ?", email).First(&employer)
	fmt.Println("Employer Email :", email, employer.Email, employer.ID)

	var job models.Job

	err = c.BodyParser(&job)

	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request Failed"})
		return err
	}
	// var employer models.User
	// database.DB.Where("ID = ?", employerID).First(&employer)
	// job.Employer = employer
	job.EmployerID = employer.ID

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

	return nil
}

func DeleteJob(c *fiber.Ctx) error {

	return nil
}
