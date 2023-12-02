package routes

import (
	"amikom-hacktalent-be/controllers"
	"amikom-hacktalent-be/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	app.Post("/checkin", controllers.CheckIn)
	app.Use(middlewares.EmailCheck)
	
	app.Post("/schedule", controllers.CreateSchedule)

	app.Get("/schedule", controllers.GetDetailSchedule)
	app.Get("/schedule", controllers.GetListSchedule)

	app.Patch("/schedule", controllers.UpdateSchedule)
	app.Delete("/schedule", controllers.DeleteSchedule)
}
