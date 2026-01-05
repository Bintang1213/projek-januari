package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func IsAdmin(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	role := claims["role"].(string)

	if role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Akses ditolak! Anda bukan Admin.",
		})
	}
	return c.Next()
}