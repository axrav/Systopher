package handlers

import (
	"github.com/axrav/Systopher/backend/types"
	"github.com/gofiber/fiber/v2"
)

func Forgetpassword(c *fiber.Ctx) error {
	email := new(types.Email)
	if err := c.BodyParser(email); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	if email.Email == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Missing email",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Email sent",
	})

}
