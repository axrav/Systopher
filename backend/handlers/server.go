package handlers

import (
	"fmt"

	"github.com/axrav/SysAnalytics/backend/helpers"
	"github.com/axrav/SysAnalytics/backend/types"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"token": helpers.GenerateServerToken(),
	})
}

func AddServer(c *fiber.Ctx) error {
	newServer := new(types.Server)
	if err := c.BodyParser(newServer); err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Wrong data",
		})
	} else {
		newServer.Owner = c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["email"].(string)
		if newServer.Ip == "" || newServer.Port == "" || newServer.Owner == "" || newServer.Token == "" {
			return c.Status(400).JSON(fiber.Map{
				"message": "Missing IP/Port/Owner/Token",
			})
		}
		// check if server already exists in database or not
		err := helpers.CheckServerAndAdd(newServer)
		if err != nil {
			fmt.Println(err)
			return c.Status(502).JSON(fiber.Map{
				"message": err.Error()})

			// might use them later for debugging purposes

			// if strings.HasSuffix(err.Error(), "unable to connect to server") || strings.HasSuffix(err.Error(), "running?") {
			// 	return c.Status(502).JSON(fiber.Map{
			// 		"message": "Unable to connect to server",
			// 	})

			// } else if strings.HasPrefix(err.Error(), "timeout") {
			// 	return c.Status(504).JSON(fiber.Map{
			// 		"message": "Timeout, PORT not exposed",
			// 	})

			// } else {
			// 	return c.Status(500).JSON(fiber.Map{
			// 		"message": "Internal Server Error",
			// 	})
			// }
		}
		return c.Status(200).JSON(fiber.Map{
			"message": "added server successfully",
		})
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
	err := helpers.CheckServerAndDelete(server)
	if err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	return c.Status(404).JSON(fiber.Map{
		"message": "Server not found",
	})
}
