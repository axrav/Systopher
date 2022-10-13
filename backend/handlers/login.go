package handlers

import (
	"github.com/axrav/SysAnalytics/backend/helpers"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	if email == "" || password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Missing email or password",
		})
	} else {
		if check, _ := helpers.CompareHashAndPassword(password, email); check {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"message": "Logged in",
			})
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Wrong password/User Not found",
			})
		}

	}
}
