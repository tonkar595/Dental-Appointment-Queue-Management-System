package controller

import (
	"backend/internal/dto"
	"backend/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type DentalAddController struct {
	service *service.DentalAddService
}

func NewDentalAddController(s *service.DentalAddService) *DentalAddController {
	return &DentalAddController{service: s}
}

func (c *DentalAddController) Create(ctx *fiber.Ctx) error {
	var req dto.CreateServiceRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "invalid request body"})
	}

	if err := c.service.CreateService(req); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(201).JSON(fiber.Map{"message": "service created successfully"})
}
func (c *DentalAddController) GetAll(ctx *fiber.Ctx) error {
	services, err := c.service.GetAllServices()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "Could not fetch services"})
	}
	return ctx.JSON(services)
}

// GET /api/services/:id
func (c *DentalAddController) GetByID(ctx *fiber.Ctx) error {
	// แปลง ID จาก String เป็น uint
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid service ID"})
	}

	service, err := c.service.GetServiceByID(uint(id))
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "Service not found"})
	}

	return ctx.JSON(service)
}
