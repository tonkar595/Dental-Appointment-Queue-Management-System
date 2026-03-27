package main

import (
	db "backend/internal/db"
	"backend/internal/routers"
	"backend/internal/wire"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db.DatabaseConnected()
	// db.AutoMigrate(db.DB)
	app := fiber.New()

	container := wire.NewContainer(db.DB)

	routers.SetupRoutes(app, container.UserController)

	app.Listen(":8080")

}
