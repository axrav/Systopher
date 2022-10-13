package routes

import (
	"os"

	"github.com/axrav/SysAnalytics/backend/handlers"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/status", func(c *fiber.Ctx) error {
		return c.SendString("Online")
	})
	auth := app.Group("/auth")
	auth.Post("/login", handlers.Login)
	auth.Post("/signup", handlers.Signup)

	// protected routes
	server := app.Group("/server")
	server.Use(jwtware.New(jwtware.Config{
		SigningKey:   []byte(os.Getenv("JWT_SECRET")),
		ErrorHandler: handlers.ErrorHandler,
	}))
	server.Post("/addserver", handlers.AddServer)
	server.Delete("/deleteserver", handlers.DeleteServer)
	app.Listen(":" + os.Getenv("SERVER_PORT"))

}
