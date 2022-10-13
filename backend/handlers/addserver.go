package handlers

import (
	"fmt"

	"github.com/axrav/SysAnalytics/backend/helpers"
	"github.com/axrav/SysAnalytics/backend/types"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func AddServer(c *fiber.Ctx) error {
	newServer := types.Server{}
	if err := c.BodyParser(&newServer); err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Wrong data",
		})
	}
	newServer.Owner = c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["email"].(string)
	fmt.Println(newServer.Owner)
	// check if server already exists in database or not
	added, err := helpers.CheckServerAndAdd(&newServer)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	if added {
		return c.Status(200).JSON(fiber.Map{
			"message": "Server added successfully",
		})
	} else {
		return c.Status(400).JSON(fiber.Map{
			"message": "Server already exists",
		})
	}

}
