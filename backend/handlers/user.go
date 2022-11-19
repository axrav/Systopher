package handlers

import (
	"fmt"

	"github.com/axrav/Systopher/backend/errors"
	"github.com/axrav/Systopher/backend/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func User(c *fiber.Ctx) error {
	email := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["email"].(string)
	data := helpers.GetUserData(email)
	return c.Status(200).JSON(data)
}

func CheckUser(c *fiber.Ctx) error {
	username := c.Query("username")
	fmt.Println("username", username)
	err := helpers.CheckUserNameExists(username)
	if err.Err != nil {
		return c.Status(409).JSON(errors.UsernameTaken.Merror())
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "user doesnt exists",
	})
}
