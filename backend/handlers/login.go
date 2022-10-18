package handlers

import (
	"fmt"

	"github.com/axrav/Systopher/backend/helpers"
	"github.com/axrav/Systopher/backend/types"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	user := c.Locals("loginUser").(*types.User)
	if check, _ := helpers.CompareHashAndPassword(user.Password, user.Email); check {
		token, err := helpers.GenerateJWT(user.Email)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error",
			})
		}
		c.Cookie(&fiber.Cookie{
			Name:  "login",
			Value: token,
		})
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Logged in",
			"token":   token,
		})
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Wrong password/User Not found",
		})
	}

}
