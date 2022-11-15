package handlers

import (
	"github.com/axrav/Systopher/backend/helpers"
	"github.com/axrav/Systopher/backend/types"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	users, err := helpers.GetAllUsers()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"users": users,
	})

}

func AddAdmin(c *fiber.Ctx) error {
	email := new(types.Email)
	if err := c.BodyParser(email); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	err := helpers.AddAdmin(email.Email)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "Admin added",
	})
}
