package controllers

import (
	"fmt"
	"net/http"
	"unified-hiring-portal/database"
	"unified-hiring-portal/models"

	"github.com/gofiber/fiber/v2"
)

func HelloApiPage(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"Message": "Hello World",
	})
}
func AddJobApplicant(c *fiber.Ctx) error {
	clientID, err := retrieve_Client_id(c)

	if err != nil {
		c.JSON(fiber.Map{
			"Error": err,
		})
	}

	fmt.Println("Client ID :", clientID)

	jobid := c.Params("jobid")

	applicant := models.Applicant{}

	err = c.BodyParser(&applicant)

	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"Message": "Request failed. Unprocessed",
		})
	}

	existingApplicant := models.Applicant{}
	err = database.DB.Where("Email = ?", applicant.Email).First(&existingApplicant).Error

	if err != nil {

		err = database.DB.Create(&applicant).Error

		if err != nil {
			return c.JSON(fiber.Map{
				"Error": "Could not create Applicant",
			})
		}
	} else {
		applicant.ID = existingApplicant.ID

	}

	application_object := map[string]interface{}{
		"job_id":       jobid,
		"applicant_id": applicant.ID,
	}

	err = database.DB.Table("job_applications").Create(&application_object).Error

	if err != nil {
		return c.JSON(fiber.Map{
			"Error": err,
		})
	}

	return c.JSON(fiber.Map{
		"Message": "Success",
	})

}

func GetAllJobs(c *fiber.Ctx) error {

	clientID, err := retrieve_Client_id(c)

	if err != nil {
		c.JSON(fiber.Map{
			"Error": err,
		})
	}

	fmt.Println("Client ID :", clientID)

	jobs := []models.Job{}

	err = database.DB.Preload("Employer").Find(&jobs).Error

	if err != nil {
		return c.JSON(fiber.Map{
			"Error": err,
		})
	}

	return c.JSON(jobs)

}

func GetAllJobsByTags(c *fiber.Ctx) error {
	clientID, err := retrieve_Client_id(c)

	if err != nil {
		c.JSON(fiber.Map{
			"Error": err,
		})
	}

	fmt.Println("Client ID :", clientID)

	var tags []models.Tag

	err = c.BodyParser(&tags)

	if err != nil {
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	var results []map[string]interface{}

	err = database.DB.Table("job_tags").Distinct("job_id").Where(tags).Find(&results).Error

	if err != nil {
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	return c.JSON(fiber.Map{
		"result": results,
	})

}

func Pdf_func(c *fiber.Ctx) error {
	formheader, err := c.FormFile("resume")

	if err != nil {
		c.JSON(fiber.Map{
			"Error": fiber.StatusBadRequest,
		})
	}

	err = c.SaveFile(formheader, formheader.Filename)

	if err != nil {
		c.JSON(fiber.Map{
			"error": err,
		})
	}

	return nil
}
