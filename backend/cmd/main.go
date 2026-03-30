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
		AllowOrigins:     "http://localhost:3000", // เปลี่ยนเป็น Domain ของ Frontend คุณ
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS, PATCH",
		AllowCredentials: true, // สำคัญมาก! เพื่อให้รับ-ส่ง Cookies ได้
	}))

	container := wire.NewContainer(db.DB)

	routers.SetupRoutes(app, container.UserController)

	app.Listen(":8080")

}
