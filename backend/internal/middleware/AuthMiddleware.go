package middleware

import (
	"errors"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(c *fiber.Ctx) error {
	// 1. ดึง Token จาก Cookie
	tokenString := c.Cookies("jwt")

	if tokenString == "" {
		authHeader := c.Get("Authorization")
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			tokenString = authHeader[7:]
		}
	}

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

		secret := os.Getenv("JWT_SECRET")
		// fmt.Println("JWT_TOKEN : ", secret)
		if secret == "" {

			secret = "default_secret_fallback"
			fmt.Println("⚠️ WARNING: JWT_SECRET is empty! Check your .env file.")
		}
		return []byte(secret), nil
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

	// 4. เก็บข้อมูลไว้ใน Locals
	c.Locals("user_id", claims["user_id"])
	c.Locals("role", claims["role"])

	return c.Next()
}
func RoleChecker(requiredRole string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 1. ดึงข้อมูล Role ที่เราเก็บไว้ใน Locals (จาก DeserializeUser)
		userRole := c.Locals("role")

		// 2. ตรวจสอบสิทธิ์
		if userRole != requiredRole {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Access denied: you do not have the required permissions",
			})
		}

		return c.Next()
	}
}
