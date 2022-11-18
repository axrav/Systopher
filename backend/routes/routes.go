package routes

import (
	"os"

	"github.com/axrav/Systopher/backend/handlers"
	"github.com/axrav/Systopher/backend/middleware"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/gofiber/websocket/v2"
)

func SetupRoutes(app *fiber.App) {
	// status check
	app.Get("/status", func(c *fiber.Ctx) error {
		return c.SendString("Online")
	})

	// auth routes
	auth := app.Group("/auth")
	auth.Post("/login", middleware.VerifyMiddleware, handlers.Login)
	auth.Post("/signup", handlers.Signup)
	auth.Post("/resendotp", handlers.ResendOTP)
	auth.Post("/verify", handlers.Verify)
	auth.Post("/reset", handlers.ForgetPassword)
	forget := app.Group("/forget")
	forget.Use(jwtware.New(jwtware.Config{
		SigningKey:   []byte(os.Getenv("FORGET_SECRET")),
		ErrorHandler: handlers.ErrorHandler,
	}))
	forget.Post("/", handlers.GenerateNewPassword)
	// protected routes
	api := app.Group("/api")
	api.Use(jwtware.New(jwtware.Config{
		SigningKey:   []byte(os.Getenv("JWT_SECRET")),
		ErrorHandler: handlers.ErrorHandler,
	}))

	// server routes
	server := api.Group("/server")
	server.Post("/addserver", handlers.AddServer)
	server.Get("/generateToken", handlers.GenerateToken)

	server.Delete("/deleteserver", middleware.ServerMiddleware, handlers.DeleteServer)

	// user routes
	user := api.Group("/user")
	user.Put("/change", handlers.ChangePassword)
	user.Get("/", handlers.User)

	// websocket route for server
	app.Get("/ws", middleware.WebSocketMiddleware, websocket.New(handlers.ServerWS))

	// admin routes
	admin := api.Group("/admin")
	admin.Use(middleware.AdminMiddleware)
	admin.Get("/users", handlers.GetUsers)
	admin.Post("/addadmin", handlers.AddAdmin)
	app.Listen(":" + os.Getenv("SERVER_PORT"))

}
