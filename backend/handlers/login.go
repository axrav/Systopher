package handlers

import (
	"fmt"

	"github.com/axrav/Systopher/backend/helpers"
	"github.com/axrav/Systopher/backend/types"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	user := c.Locals("loginUser").(*types.LoginUser)
	if check, _ := helpers.CompareHashAndPassword(user.Password, user.Email); check {
		token, err := helpers.GenerateJWT(user.Email, user.Remember, "browse")
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error",
			})
		}
		data := helpers.GetUserData(user.Email)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Logged in",
			"token":   token,
			"user":    data,
		})
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Wrong password/User Not found",
		})
	}

}
