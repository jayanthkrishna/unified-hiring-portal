package controllers

import (
	"unified-hiring-portal/database"
	"unified-hiring-portal/models"

	"github.com/gofiber/fiber/v2"
)

func AddJobApplicant(c *fiber.Ctx) error {

	return nil
}

func GetAllJobs(c *fiber.Ctx) error {

	jobs := []models.Job{}

	err := database.DB.Preload("Employer").Find(&jobs).Error

	if err != nil {
		return c.JSON(fiber.Map{
			"Error": err,
		})
	}

	return c.JSON(jobs)

}
