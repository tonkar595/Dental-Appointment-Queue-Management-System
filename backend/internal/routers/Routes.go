package routers

import (
	"backend/internal/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, controller *controller.UserController) {
	api := app.Group("/api")
	users := api.Group("/users")

	users.Post("/create", controller.Create)
	users.Get("/get-all", controller.GetAll)
}
