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
	}

	// Validate input data
	if user.Username == "" || user.Email == "" || user.Password == "" {
		return c.Status(400).JSON(errors.InvalidData.Merror())
	}

	// Comprehensive validation using UserCheckers
	if validationError := helpers.UserCheckers(user); validationError.Err != nil {
		return c.Status(400).JSON(validationError)
	}

	// Hash password
	hash, err := helpers.HashPassword(user.Password)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(errors.InternalServerError.Merror())
	}

	// Generate UUID for user
	u_id := helpers.GenerateUUIDForUser()

	// Create user in database
	if err = helpers.CreateUser(user.Email, hash, user.Username, u_id); err != nil {
		if strings.HasSuffix(err.Error(), "\"users_email_key\"") {
			return c.Status(409).JSON(errors.EmailTaken.Merror())
		}
		if strings.HasSuffix(err.Error(), "\"users_username_key\"") {
			return c.Status(409).JSON(errors.UsernameTaken.Merror())
		}
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(errors.InternalServerError.Merror())
	}

	// Send OTP for email verification
	sent, err := helpers.SendOtp(user.Email)
	if err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(errors.InternalServerError.Merror())
	}

	// Save OTP and set user ID in Redis
	success := helpers.SaveOtp(user.Email, sent)
	if !success {
		return c.Status(500).JSON(errors.InternalServerError.Merror())
	}

	err = helpers.SetUserId(u_id, user.Email)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(errors.InternalServerError.Merror())
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "User created, verify to continue",
	})
}
