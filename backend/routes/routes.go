package routes

import (
	"os"

	"github.com/axrav/SysAnalytics/backend/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Online")
	})
	app.Post("/login", handlers.Login)
	app.Post("/signup", handlers.Signup)
	// app.Post("/addserver", handlers.AddServer)
	// app.Delete("/deleteserver", handlers.DeleteServer)
	app.Listen(":" + os.Getenv("SERVER_PORT"))

}
