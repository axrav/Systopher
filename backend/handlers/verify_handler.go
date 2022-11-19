package handlers

import (
	"fmt"

	"github.com/axrav/Systopher/backend/errors"
	"github.com/axrav/Systopher/backend/helpers"
	"github.com/axrav/Systopher/backend/models"
	"github.com/gofiber/fiber/v2"
)

func Verify(c *fiber.Ctx) error {
	resp := new(models.OTPResponse)
	if err := c.BodyParser(resp); err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(errors.InvalidData.Merror())
	} else {
		if resp.Otp == "" || resp.Email == "" {
			return c.Status(400).JSON(errors.InvalidData.Merror())
		}
		out := helpers.GetVerified(resp.Email)
		if out {
			return c.Status(409).JSON(errors.AlreadyVerified.Merror())
		}
		verified := helpers.VerifyOtp(resp.Email, resp.Otp)
		if verified {
			if err := helpers.SetVerify(resp.Email); err != nil {
				fmt.Println(err)
				return c.Status(500).JSON(errors.InternalServerError.Merror())
			}
			token, err := helpers.GenerateJWT(resp.Email, false, "browse")
			if err != nil {
				fmt.Println(err)
				return c.Status(500).JSON(errors.InternalServerError.Merror())
			}
			return c.JSON(fiber.Map{
				"message": "Verified",
				"token":   token,
				"user":    helpers.GetUserData(resp.Email),
			})
		} else {
			return c.Status(400).JSON(errors.InvalidOtp.Merror())
		}
	}
}

func ResendOTP(c *fiber.Ctx) error {
	email := new(models.Email)
	if err := c.BodyParser(email); err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(errors.InvalidData.Merror())
	} else {
		if email.Email == "" {
			return c.Status(400).JSON(errors.InvalidData.Merror())
		}
		out := helpers.GetVerified(email.Email)
		if out {
			return c.Status(409).JSON(errors.AlreadyVerified.Merror())
		}

		sent, err := helpers.SendOtp(email.Email)
		if err != nil {
			fmt.Println(err)
			return c.Status(500).JSON(errors.InternalServerError.Merror())
		}
		output := helpers.SaveOtp(email.Email, sent)

		if output {
			return c.JSON(fiber.Map{
				"message": "OTP sent",
			})
		}
	}
	return c.Status(500).JSON(errors.InternalServerError.Merror())
}
