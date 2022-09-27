package routes

import (
	"unified-hiring-portal/controllers"

	"github.com/gofiber/fiber/v2"
)

func GetApiRoutes(app *fiber.App) {
	app.Get("/", controllers.HelloApiPage)
	app.Get("/jobs", controllers.GetAllJobs)
	app.Post("/jobs/:jobid", controllers.AddJobApplicant)
	app.Get("/identity/token", controllers.GenerateToken)
}
