package wire

import (
	"backend/internal/controller"
	"backend/internal/repository"
	"backend/internal/service"

	"gorm.io/gorm"
)

type Container struct {
	UserController *controller.UserController
}

func NewContainer(db *gorm.DB) *Container {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)
	return &Container{
		UserController: userController,
	}
}
