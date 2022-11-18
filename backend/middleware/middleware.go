package middleware

import (
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
	user := new(types.LoginUser)
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
			"message": "Unverified or user doesnt' exists",
		})
	}
	c.Locals("loginUser", user)
	return c.Next()
}

func AdminMiddleware(c *fiber.Ctx) error {
	email := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["email"].(string)
	isAdmin := helpers.GetAdmin(email)
	if !isAdmin {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	return c.Next()

}

func SignupChecks(c *fiber.Ctx) error {
	user := new(types.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Wrong data",
		})
	} else {
		if user.Email == "" || user.Password == "" || user.Username == "" {
			return c.Status(400).JSON(fiber.Map{
				"message": "Missing email, password or name",
			})
		}
	}
	err := helpers.UserCheckers(user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	c.Locals("user", user)
	return c.Next()

}
