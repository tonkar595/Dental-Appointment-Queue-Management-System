package controller

import (
	"backend/internal/dto"
	"backend/internal/service"

	"github.com/gofiber/fiber/v2"
)

type ClinicController struct {
	service *service.ClinicService
}

func NewClinicController(s *service.ClinicService) *ClinicController {
	return &ClinicController{service: s}
}

// API สำหรับ Admin/หมอ ตั้งค่าวันพิเศษ
func (c *ClinicController) SetCustomSchedule(ctx *fiber.Ctx) error {
	var req dto.UpdateClinicScheduleRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := c.service.UpdateCustomSchedule(req); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "บันทึกการตั้งค่าเรียบร้อยแล้ว"})
}

// API สำหรับเช็คสถานะวัน (ใช้ได้ทั้งหน้าบ้านคนไข้และหมอ)
func (c *ClinicController) CheckAvailability(ctx *fiber.Ctx) error {
	dateStr := ctx.Query("date") // /api/clinic/check?date=2026-04-10
	if dateStr == "" {
		return ctx.Status(400).JSON(fiber.Map{"error": "Date is required"})
	}

	res := c.service.GetEffectiveSchedule(dateStr)
	return ctx.JSON(res)
}
