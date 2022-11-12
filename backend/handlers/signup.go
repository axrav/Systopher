package handlers

import (
	"fmt"
	"strings"

	"github.com/axrav/Systopher/backend/helpers"
	"github.com/axrav/Systopher/backend/types"
	"github.com/gofiber/fiber/v2"
)

func Signup(c *fiber.Ctx) error {
	user := new(types.User)
	if err := c.BodyParser(user); err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Wrong data",
		})
	} else {
		if user.Username == "" || user.Email == "" || user.Password == "" {
			return c.Status(400).JSON(fiber.Map{
				"message": "Missing Username/email or password",
			})
		}
		if hash, err := helpers.HashPassword(user.Password); err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal server error",
			})
		} else {
			sent, err := helpers.SendOtp(user.Email)
			if err != nil {
				fmt.Println(err)
				return c.Status(500).JSON(fiber.Map{
					"message": "Internal server error",
				})
			}
			success := helpers.SaveOtp(user.Email, sent)
			if success {
				u_id := helpers.GenerateUserId()
				if err = helpers.CreateUser(user.Email, hash, user.Username, u_id); err != nil {
					if strings.HasSuffix(err.Error(), "\"users_email_key\"") {
						return c.Status(409).JSON(fiber.Map{
							"message": "user already exists",
						})

					}
					if strings.HasSuffix(err.Error(), "\"users_username_key\"") {
						return c.Status(409).JSON(fiber.Map{
							"message": "username is taken",
						})
					}
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"message": "Unable to create user" + err.Error(),
					})
				}
				err := helpers.SetUserId(u_id, user.Email)
				if err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"message": "Unable to create user" + err.Error(),
					})
				}
				return c.JSON(fiber.Map{
					"message": "Signed up, verify to continue",
				})
			}
			return c.Status(500).JSON(fiber.Map{"message": "Internal Server Error"})
		}
	}
}
