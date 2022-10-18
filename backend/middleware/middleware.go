package middleware

import (
	"github.com/axrav/SysAnalytics/backend/helpers"
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

func VerifyMiddleware(c *fiber.Ctx) error {
	email := c.FormValue("email")
	isVerified := helpers.GetVerified(email)
	if !isVerified {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unverified or user not exists",
		})
	}
	return c.Next()
}
