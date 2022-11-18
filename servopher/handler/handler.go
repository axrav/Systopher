package handler

import (
	"github.com/axrav/Systopher/servopher/helpers"
	"github.com/gofiber/fiber/v2"
)

func GetStats(c *fiber.Ctx) error {
	stats := helpers.SendServerData()
	return c.JSON(stats)
}
