package handlers

import (
	"github.com/axrav/Systopher/backend/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func User(c *fiber.Ctx) error {
	email := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["email"].(string)
	data := helpers.GetUserData(email)
	return c.Status(200).JSON(data)
}
