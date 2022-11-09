package handlers

import (
	"github.com/axrav/Systopher/backend/helpers"
	"github.com/axrav/Systopher/backend/types"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func ChangePassword(c *fiber.Ctx) error {
	email := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["email"].(string)
	password := new(types.Password)
	if err := c.BodyParser(password); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Wrong data",
		})
	} else {
		if password.Password == "" || password.NewPassword == "" {
			return c.Status(400).JSON(fiber.Map{
				"message": "Missing old or new password",
			})
		}
	}

	err := helpers.ChangePassword(email, password.Password, password.NewPassword)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "Password changed",
	})
}

func ForgetPassword(c *fiber.Ctx) error {
	return nil
}
