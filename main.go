package main

import (
	"jwt-auth-go/database"
	"jwt-auth-go/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	database.Connect()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.GetRoutes(app)

	app.Listen(":8000")
}
