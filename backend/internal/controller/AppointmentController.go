package controller

import (
	"backend/internal/dto"
	"backend/internal/service"
	"strconv"

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

// GET /api/appointments?date=2026-04-10
func (c *AppointmentController) GetAppointments(ctx *fiber.Ctx) error {
	date := ctx.Query("date") // รับวันที่จาก query param ถ้าไม่มีจะดึงทั้งหมด

	appointments, err := c.service.GetAppointments(date)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "ไม่สามารถดึงข้อมูลนัดหมายได้"})
	}

	return ctx.JSON(appointments)
}

// GET /api/appointments/patient/:id
func (c *AppointmentController) GetPatientHistory(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "รหัสคนไข้ไม่ถูกต้อง"})
	}

	history, err := c.service.GetPatientHistory(uint(id))
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(history)
}

func (c *AppointmentController) Patch(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")
	var req dto.UpdateAppointmentRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(400).JSON(dto.MessageResponse{
			Status:  "error",
			Message: "รูปแบบข้อมูลไม่ถูกต้อง",
		})
	}

	err := c.service.PatchAppointment(uint(id), req)
	if err != nil {
		return ctx.Status(500).JSON(dto.MessageResponse{
			Status:  "error",
			Message: "ไม่สามารถอัปเดตนัดหมายได้: " + err.Error(),
		})
	}

	return ctx.Status(200).JSON(dto.MessageResponse{
		Status:  "success",
		Message: "อัปเดตนัดหมายเรียบร้อยแล้ว",
	})
}

func (c *AppointmentController) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(400).JSON(dto.MessageResponse{
			Status:  "error",
			Message: "รหัสนัดหมายไม่ถูกต้อง",
		})
	}

	err = c.service.DeleteAppointment(uint(id))
	if err != nil {
		return ctx.Status(500).JSON(dto.MessageResponse{
			Status:  "error",
			Message: "ลบนัดหมายไม่สำเร็จ: " + err.Error(),
		})
	}

	return ctx.Status(200).JSON(dto.MessageResponse{
		Status:  "success",
		Message: "ลบนัดหมายเรียบร้อยแล้ว",
	})
}
