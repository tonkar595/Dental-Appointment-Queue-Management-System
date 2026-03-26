package controller

import (
	models "backend/internal/model"
	"backend/internal/service"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService}
}

func (c *UserController) Create(ctx *fiber.Ctx) error {
	var user models.User
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := c.userService.CreateUser(&user); err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(201).JSON(fiber.Map{
		"user":    user,
		"message": "Successfully created user",
		"status":  201,
	})
}
func (c *UserController) GetAll(ctx *fiber.Ctx) error {
	users, err := c.userService.GetAllUsers()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(fiber.Map{
		"user":    users,
		"message": "Success",
		"status":  200,
	})
}
func (c *UserController) GetByID(ctx *fiber.Ctx) error {
	return nil
}
func (c *UserController) Update(ctx *fiber.Ctx) error {
	return nil
}
func (c *UserController) Delete(ctx *fiber.Ctx) error {
	return nil
}
