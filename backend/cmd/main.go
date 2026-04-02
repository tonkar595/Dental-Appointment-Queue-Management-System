package main

import (
	db "backend/internal/db"
	"backend/internal/routers"
	"backend/internal/wire"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	db.DatabaseConnected()
	// db.AutoMigrate(db.DB)
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS, PATCH",
		AllowCredentials: true,
	}))

	container := wire.NewContainer(db.DB)

	routers.SetupRoutes(app, container)

	app.Listen(":8080")

}
