package controllers

import (
	"unified-hiring-portal/models"

	"github.com/gofiber/fiber/v2"
)

func GenerateToken(c *fiber.Ctx) error {

	var client models.Client

	if err := c.BodyParser(&client); err != nil {
		return err
	}

	return nil

}
