package controller

import (
	"backend/internal/dto"
	"backend/internal/service"

	"github.com/gofiber/fiber/v2"
)

type AppointmentController struct {
	service *service.AppointmentService
}

func NewAppointmentController(s *service.AppointmentService) *AppointmentController {
	return &AppointmentController{service: s}
}

// POST: /api/appointments
func (c *AppointmentController) Create(ctx *fiber.Ctx) error {
	var req dto.CreateAppointmentRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := c.service.CreateAppointment(req); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(201).JSON(fiber.Map{"message": "จองนัดหมายสำเร็จ"})
}
