package controller

import (
	"backend/internal/dto"
	"backend/internal/service"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	service *service.AuthService
}

func NewAuthController(s *service.AuthService) *AuthController {
	return &AuthController{service: s}
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	var req dto.LoginRequest // ใช้ DTO ที่แยกไฟล์ไว้

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "bad request"})
	}

	token, err := c.service.Login(req.Identity, req.Password)
	if err != nil {
		return ctx.Status(401).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(dto.LoginResponse{
		Token: token,
		Type:  "Bearer",
	})
}
