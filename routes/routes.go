package routes

import (
	"unified-hiring-portal/controllers"

	"github.com/gofiber/fiber/v2"
)

func GetRoutes(app *fiber.App) {

	app.Get("/", controllers.Hello)

	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
	app.Get("/user", controllers.User)
	app.Get("/jobs", controllers.GetAllJobsPostedByUser)
	app.Post("/jobs/postjob", controllers.AddJobPost)
	app.Post("/jobs/:jobid", controllers.UpdateJob)
	app.Delete("/jobs/:jobid", controllers.DeleteJob)
}

func GetApiRoutes(app *fiber.App) {
	app.Get("/jobs", controllers.GetAllJobs)
	app.Post("/jobs/:jobid", controllers.AddJobApplicant)
}
