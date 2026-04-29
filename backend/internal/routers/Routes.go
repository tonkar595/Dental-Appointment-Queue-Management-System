package routers

import (
	"backend/internal/middleware"
	"backend/internal/wire"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, c *wire.Container) {
	api := app.Group("/api")

	api.Post("/login", c.AuthController.Login)
	api.Post("/register", c.AuthController.Register)
	api.Post("/logout", c.AuthController.Logout)
	//==========================Dentist===============================================
	dentistGroup := api.Group("/dentist", middleware.AuthMiddleware, middleware.RoleChecker("Dentist"))
	dentistGroup.Post("/createdService", c.DentalAddController.Create)
	dentistGroup.Get("/show-all", c.DentalAddController.GetAll)
	dentistGroup.Get("/show/:id", c.DentalAddController.GetByID)
	dentistGroup.Patch("/service/:id", c.DentalAddController.Update)
	dentistGroup.Patch("/service/:id/status", c.DentalAddController.ToggleStatus)

	//===========================================================================================
	dentistGroup.Post("/clinic/schedule", c.ClinicController.SetCustomSchedule)
	dentistGroup.Get("/clinic/availability", c.ClinicController.CheckAvailability)
	dentistGroup.Post("/appointments/created", c.AppointmentController.Create)
	dentistGroup.Get("/appointments/getall", c.AppointmentController.GetAppointments)
	dentistGroup.Get("/appointments/patient/:id", c.AppointmentController.GetPatientHistory)
	dentistGroup.Patch("/appointments/updated/:id", c.AppointmentController.Patch)
	dentistGroup.Delete("/appointments/deleted/:id", c.AppointmentController.Delete)

	//==========================Patient===============================================
	// patientGroup := api.Group("/patient", middleware.RoleChecker("Patient"))
	// patientGroup.Post("/")

}
