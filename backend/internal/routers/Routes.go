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
	//==========================Patient===============================================
	// patientGroup := api.Group("/patient", middleware.RoleChecker("Patient"))

}
