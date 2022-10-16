package routes

import (
	"os"

	"github.com/axrav/SysAnalytics/backend/handlers"
	"github.com/axrav/SysAnalytics/backend/middleware"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/gofiber/websocket/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/status", func(c *fiber.Ctx) error {
		return c.SendString("Online")
	})
	auth := app.Group("/auth")
	auth.Post("/login", middleware.VerifyMiddleware, handlers.Login)
	auth.Post("/signup", handlers.Signup)
	auth.Post("/verify", handlers.Verify)
	// protected routes
	server := app.Group("/server")
	server.Use(jwtware.New(jwtware.Config{
		SigningKey:   []byte(os.Getenv("JWT_SECRET")),
		ErrorHandler: handlers.ErrorHandler,
	}))
	server.Post("/addserver", handlers.AddServer)
	server.Delete("/deleteserver", middleware.ServerMiddleware, handlers.DeleteServer)

	// protected routes
	// websocket route for server
	wsgroup := server.Group("/ws", middleware.ServerMiddleware)
	wsgroup.Get("/:api", websocket.New(handlers.ServerWS))

	app.Listen(":" + os.Getenv("SERVER_PORT"))

}
