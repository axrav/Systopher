package middleware

import (
	"fmt"

	"github.com/axrav/Systopher/backend/helpers"
	"github.com/axrav/Systopher/backend/types"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func ServerMiddleware(c *fiber.Ctx) error {
	email := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["email"].(string)
	servers := helpers.GetServers(email)
	if len(servers) == 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "No servers found",
		})
	}

	c.Locals("servers", servers)
	return c.Next()
}

func WebSocketMiddleware(c *fiber.Ctx) error {
	fmt.Println(c.Query("userid"))
	email := helpers.GetEmailFromId(c.Query("userid"))
	if email == "" {
		return c.Status(500).JSON(fiber.Map{
			"message": "Invalid user id",
		})
	}
	servers := helpers.GetServers(email)
	if len(servers) == 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "No servers found",
		})
	}
	c.Locals("servers", servers)
	return c.Next()
}

func VerifyMiddleware(c *fiber.Ctx) error {
	user := new(types.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Wrong data",
		})
	} else {
		if user.Email == "" || user.Password == "" {
			return c.Status(400).JSON(fiber.Map{
				"message": "Missing email or password",
			})
		}
	}
	isVerified := helpers.GetVerified(user.Email)
	if !isVerified {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unverified or user not exists",
		})
	}
	c.Locals("loginUser", user)
	return c.Next()
}
