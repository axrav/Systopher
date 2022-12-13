package middleware

import (
	"github.com/axrav/Systopher/servopher/helpers"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	header := c.Get("X-API-KEY")
	if header == helpers.ApiKey {
		return c.Next()
	} else {
		return c.Status(401).JSON(fiber.Map{"message": "Unauthorized"})

	}
}
