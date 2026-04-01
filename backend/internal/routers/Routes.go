package routers

import (
	"backend/internal/wire"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, c *wire.Container) {
	api := app.Group("/api")
	// users := api.Group("/users")
	api.Post("/login", c.AuthController.Login)
	api.Post("/register", c.AuthController.Register)
	api.Post("/logout", c.AuthController.Logout)

	// users.Post("/create", controller.Create)
	// users.Get("/get-all", controller.GetAll)
}
