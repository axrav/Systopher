package handlers

import (
	"fmt"

	"github.com/axrav/Systopher/backend/db"
	"github.com/axrav/Systopher/backend/errors"
	"github.com/axrav/Systopher/backend/helpers"
	"github.com/gofiber/fiber/v2"
)

func GenerateToken(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"token": helpers.GenerateId("SYSTO-"),
	})
}

func AddServer(c *fiber.Ctx) error {
	newServer := new(db.Server)
	if err := c.BodyParser(newServer); err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Wrong data",
		})
	} else {
		newServer.Owner = c.Locals("userdata").(db.User)
		if newServer.Ip == "" || newServer.Port == "" || newServer.Token == "" {
			return c.Status(400).JSON(errors.InvalidData.Merror())
		}
		// check if server already exists in database or not
		err := helpers.CheckServerAndAdd(newServer)
		if err != nil {
			fmt.Println(err)
			return c.Status(502).JSON(fiber.Map{
				"message": err.Error()})

		}
		return c.Status(200).JSON(fiber.Map{
			"message": "added server successfully",
		})
	}
}
func DeleteServer(c *fiber.Ctx) error {
	server := new(db.Server)
	if err := c.BodyParser(server); err != nil {
		return c.Status(500).JSON(errors.InvalidData.Merror())
	}
	if server.Ip == "" {
		return c.Status(400).JSON(errors.InvalidData.Merror())
	}
	server.Owner = c.Locals("userdata").(db.User)
	err := helpers.CheckServerAndDelete(server)
	if err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(errors.InternalServerError.Merror())
	}
	return c.Status(404).JSON(errors.NotFound.Merror())
}
