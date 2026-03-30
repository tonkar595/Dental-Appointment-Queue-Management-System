package middleware

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func DeserializeUser(c *fiber.Ctx) error {
	// 1. ดึง Token จาก Cookie
	tokenString := c.Cookies("jwt")

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "you are not logged in",
		})
	}

	// 2. ตรวจสอบ Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// ตรวจสอบ Signing Method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte("your_secret_key"), nil // ต้องตรงกับใน AuthService
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid token or expired",
		})
	}

	// 3. แกะข้อมูลจาก Claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token claims"})
	}

	// 4. เก็บข้อมูลไว้ใน Locals เพื่อให้ Controller อื่นเรียกใช้ได้ง่ายๆ
	c.Locals("user_id", claims["user_id"])
	c.Locals("role", claims["role"])

	return c.Next()
}
