package handlers

import (
	"fmt"
	"strings"

	"github.com/axrav/SysAnalytics/backend/db"
	"github.com/axrav/SysAnalytics/backend/helpers"
	"github.com/axrav/SysAnalytics/backend/types"
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
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	} else {
		sent := helpers.SendOtpAndSave(email)
		if sent {
			if _, err = db.Db.Exec("INSERT INTO users (email, password, isverified) VALUES ($1, $2, $3)", email, hash, false); err != nil {
				if strings.HasSuffix(err.Error(), "\"users_email_key\"") {
					return c.Status(403).JSON(fiber.Map{
						"message": "user already exists",
					})

				}
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Unable to create user",
				})
			}
			return c.JSON(fiber.Map{
				"message": "Signed up, verify to continue",
			})
		}
		return c.Status(500).JSON(fiber.Map{"message": "Internal Server Error"})
	}
}

func Verify(c *fiber.Ctx) error {
	resp := new(types.OTPResponse)
	if err := c.BodyParser(resp); err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Wrong data",
		})
	} else {
		if resp.Otp == "" || resp.Email == "" {
			return c.Status(400).JSON(fiber.Map{
				"message": "Missing email or otp",
			})
		}
		out := helpers.GetVerified(resp.Email)
		if out {
			return c.Status(409).JSON(fiber.Map{
				"message": "the user is already verified",
			})
		}
		verified := helpers.VerifyOtp(resp.Email, resp.Otp)
		if verified {
			if _, err := db.Db.Exec("UPDATE users SET isverified = true WHERE email = $1", resp.Email); err != nil {
				fmt.Println(err)
				return c.Status(500).JSON(fiber.Map{
					"message": "Internal server error",
				})
			}
			return c.JSON(fiber.Map{
				"message": "Verified",
			})
		} else {
			return c.Status(400).JSON(fiber.Map{
				"message": "Wrong otp",
			})
		}
	}
}
