package handlers

import (
	"fmt"

	"github.com/axrav/Systopher/backend/errors"
	"github.com/axrav/Systopher/backend/helpers"
	"github.com/axrav/Systopher/backend/models"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	users, err := helpers.GetAllUsers()
	if err != nil {
		fmt.Println("error", err)
		return c.Status(500).JSON(errors.InternalServerError)
	}
	return c.Status(200).JSON(fiber.Map{
		"users": users,
	})

}

func AddAdmin(c *fiber.Ctx) error {
	email := new(models.Email)
	if err := c.BodyParser(email); err != nil {
		fmt.Println("error", err)
		return c.Status(500).JSON(errors.InternalServerError)
	}
	err := helpers.AddAdmin(email.Email)
	if err != nil {
		fmt.Println("error", err)
		return c.Status(500).JSON(errors.InternalServerError)
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "Admin added",
	})
}
