package routes

import (
	"os"

	"github.com/axrav/Systopher/servopher/handler"
	"github.com/axrav/Systopher/servopher/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	app.Use(logger.New())
	app.Use(middleware.AuthMiddleware)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"message": "Online"})
	})
	app.Get("/stats", handler.GetStats)
	app.Listen(":" + os.Getenv("PORT"))
}
