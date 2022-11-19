package middleware

import (
	"fmt"

	"github.com/axrav/Systopher/servopher/helpers"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	header := c.Get("X-API-KEY")
	fmt.Println("Key: ", header)
	if header == "" {
		return c.Status(401).JSON(fiber.Map{"message": "Unauthorized"})
	}
	if header == helpers.ApiKey {
		return c.Next()
	} else {
		return c.Status(401).JSON(fiber.Map{"message": "Unauthorized"})

	}
}
