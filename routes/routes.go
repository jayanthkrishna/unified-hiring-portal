package routes

import (
	"jwt-auth-go/controllers"

	"github.com/gofiber/fiber/v2"
)

func GetRoutes(app *fiber.App) {

	app.Get("/", controllers.Hello)

	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
}
