package handlers

import (
	"github.com/axrav/Systopher/backend/errors"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	return c.Status(401).JSON(errors.InvalidToken.Merror())
}
