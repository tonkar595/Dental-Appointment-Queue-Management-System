package controller

import (
	"backend/internal/dto"
	"backend/internal/service"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	service *service.AuthService
}

func NewAuthController(s *service.AuthService) *AuthController {
	return &AuthController{service: s}
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	var req dto.LoginRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "bad request"})
	}

	// 1. เรียก Service เพื่อยืนยันตัวตนและรับ Token
	token, err := c.service.Login(req.Identity, req.Password)
	if err != nil {
		return ctx.Status(401).JSON(fiber.Map{"error": err.Error()})
	}

	// 2. ตั้งค่า Cookie
	ctx.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24), // หมดอายุใน 24 ชม.
		HTTPOnly: true,                           // สำคัญ! ป้องกัน JS อ่าน Token
		Secure:   false,                          // ถ้าใช้ HTTPS จริงๆ ใน Production ให้ตั้งเป็น true
		SameSite: "Lax",                          // ป้องกัน CSRF ในระดับหนึ่ง
		Path:     "/",
	})

	// 3. ตอบกลับ (ไม่ต้องส่ง Token ไปใน Body แล้วก็ได้ หรือจะส่งไปทั้งคู่ก็ไม่ผิดครับ)

	return ctx.JSON(dto.LoginResponse{
		Message: "login successful",
		Token:   token,
		Type:    "Bearer",
	})
}

func (c *AuthController) Register(ctx *fiber.Ctx) error {
	var req dto.RegisterRequest

	// 1. Parse JSON จากหน้าบ้าน
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	// 2. เรียก Service เพื่อทำการสมัครสมาชิก
	if err := c.service.Register(req); err != nil {
		msg := err.Error()

		// ถ้าเป็น Error เรื่องข้อมูลซ้ำ ให้ตอบ 409 Conflict
		if strings.Contains(msg, "already exists") {
			return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error": msg,
			})
		}

		// Error อื่นๆ ให้ตอบ 500
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "internal server error",
		})
	}

	// 3. ตอบกลับเมื่อสำเร็จ
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  fiber.StatusCreated,
		"message": "registration successful",
	})
}

func (c *AuthController) Logout(ctx *fiber.Ctx) error {
	ctx.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), // ตั้งให้ย้อนหลังเพื่อลบทิ้ง
		HTTPOnly: true,
	})
	return ctx.JSON(fiber.Map{
		"message": "logged out successfully",
		"status":  201,
	})
}
