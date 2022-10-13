package handlers

import (
	"fmt"

	"github.com/axrav/SysAnalytics/backend/helpers"
	"github.com/axrav/SysAnalytics/backend/types"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func AddServer(c *fiber.Ctx) error {
	newServer := new(types.Server)
	if err := c.BodyParser(newServer); err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Wrong data",
		})
	} else {
		newServer.Owner = c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["email"].(string)
		if newServer.Ip == "" || newServer.Port == "" || newServer.Owner == "" {
			return c.Status(400).JSON(fiber.Map{
				"message": "Missing IP/Port/Owner",
			})
		}
		// check if server already exists in database or not
		added, err := helpers.CheckServerAndAdd(newServer)
		if err != nil {
			fmt.Println(err)
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
}
func DeleteServer(c *fiber.Ctx) error {
	server := new(types.Server)
	if err := c.BodyParser(server); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Wrong data",
		})
	}
	if server.Ip == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Missing IP",
		})
	}
	server.Owner = c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["email"].(string)
	deleted, err := helpers.CheckServerAndDelete(server)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	if deleted {
		return c.Status(200).JSON(fiber.Map{
			"message": "Server deleted successfully",
		})
	} else {
		return c.Status(404).JSON(fiber.Map{
			"message": "Server not found",
		})
	}
}
