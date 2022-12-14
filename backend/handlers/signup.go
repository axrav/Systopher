package handlers

import (
	"fmt"
	"strings"

	"github.com/axrav/Systopher/backend/errors"
	"github.com/axrav/Systopher/backend/helpers"
	"github.com/axrav/Systopher/backend/models"
	"github.com/gofiber/fiber/v2"
)

func Signup(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(errors.InternalServerError.Merror())
	} else {
		if user.Username == "" || user.Email == "" || user.Password == "" {
			return c.Status(400).JSON(errors.InvalidData.Merror())
		}
		if hash, err := helpers.HashPassword(user.Password); err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(errors.InternalServerError.Merror())
		} else {
			u_id := helpers.GenerateId("USER-")
			if err = helpers.CreateUser(user.Email, hash, user.Username, u_id); err != nil {
				if strings.HasSuffix(err.Error(), "\"users_email_key\" (SQLSTATE 23505)") {
					return c.Status(409).JSON(errors.EmailTaken.Merror())

				}
				if strings.HasSuffix(err.Error(), "\"users_username_key\" (SQLSTATE 23505)") {
					return c.Status(409).JSON(errors.UsernameTaken.Merror())
				}
				return c.Status(fiber.StatusInternalServerError).JSON(errors.InternalServerError.Merror())
			} else {
				sent, err := helpers.SendOtp(user.Email)
				if err != nil {
					fmt.Println(err)
					return c.Status(500).JSON(errors.InternalServerError.Merror())
				}
				success := helpers.SaveOtp(user.Email, sent)
				if success {
					err := helpers.SetUserId(u_id, user.Email)
					if err != nil {
						return c.Status(fiber.StatusInternalServerError).JSON(errors.InternalServerError.Merror())
					}
					return c.Status(200).JSON(fiber.Map{
						"message": "User created, verify to continue",
					})
				}
				return c.Status(500).JSON(errors.InternalServerError.Merror())
			}
		}
	}
}
