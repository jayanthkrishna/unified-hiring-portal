package main

import (
	"unified-hiring-portal/database"
	"unified-hiring-portal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	api := fiber.New()
	database.Connect()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	api.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.GetApiRoutes(api)

	routes.GetRoutes(app)

	app.Listen(":8000")
	api.Listen(":8001")
}
