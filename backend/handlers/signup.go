package handlers

import (
	"fmt"

	"github.com/axrav/SysAnalytics/backend/db"
	"github.com/axrav/SysAnalytics/backend/helpers"
	"github.com/gofiber/fiber/v2"
)

func Signup(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	if email == "" || password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Missing email or password",
		})
	}
	if hash, err := helpers.HashPassword(password); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	} else {
		if _, err = db.Db.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", email, hash); err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Unable to create user",
			})
		}

		return c.JSON(fiber.Map{
			"message": "Signed up successfully",
		})
	}
}
