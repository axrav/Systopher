package middleware

import (
	"github.com/axrav/Systopher/backend/errors"
	"github.com/axrav/Systopher/backend/helpers"
	"github.com/axrav/Systopher/backend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func ServerMiddleware(c *fiber.Ctx) error {
	email := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["email"].(string)
	servers := helpers.GetServers(email)
	if len(servers) == 0 {
		return c.Status(400).JSON(errors.NotFound.Merror())
	}

	c.Locals("servers", servers)
	return c.Next()
}

func WebSocketMiddleware(c *fiber.Ctx) error {
	email := helpers.GetEmailFromId(c.Query("token"))
	if email == "" {
		return c.Status(500).JSON(errors.InvalidToken.Merror())
	}
	servers := helpers.GetServers(email)
	if len(servers) == 0 {
		return c.Status(400).JSON(errors.NotFound.Merror())
	}
	c.Locals("servers", servers)
	return c.Next()
}

func VerifyMiddleware(c *fiber.Ctx) error {
	user := new(models.LoginUser)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).JSON(errors.InternalServerError.Merror())
	} else {
		if user.Email == "" || user.Password == "" {
			return c.Status(400).JSON(errors.InternalServerError.Merror())
		}
	}
	isVerified := helpers.GetVerified(user.Email)
	if !isVerified {
		return c.Status(401).JSON(errors.InvalidUser.Merror())
	}
	c.Locals("loginUser", user)
	return c.Next()
}

func AdminMiddleware(c *fiber.Ctx) error {
	email := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["email"].(string)
	isAdmin := helpers.GetAdmin(email)
	if !isAdmin {
		return c.Status(401).JSON(errors.InvalidToken.Merror())
	}
	return c.Next()

}

func SignupChecks(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).JSON(errors.InvalidData.Merror())
	} else {
		if user.Email == "" || user.Password == "" || user.Username == "" {
			return c.Status(400).JSON(errors.InvalidData.Merror())
		}
	}
	err := helpers.UserCheckers(user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Next()

}
