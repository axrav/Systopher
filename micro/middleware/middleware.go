package middleware

import (
	"github.com/axrav/Systopher/micro/helpers"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	header := c.GetRespHeader("X-API-KEY")
	if header == "" {
		return c.Status(401).JSON(fiber.Map{"message": "Unauthorized"})
	}
	if header == helpers.ApiKey {
		return c.Next()
	} else {
		return c.Status(401).JSON(fiber.Map{"message": "Unauthorized"})

	}
}
