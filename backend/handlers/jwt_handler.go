package handlers

import "github.com/gofiber/fiber/v2"

func ErrorHandler(c *fiber.Ctx, err error) error {
	return c.Status(401).JSON(fiber.Map{
		"message": "Unauthorized",
	})
}
