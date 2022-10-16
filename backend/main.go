package main

import (
	"github.com/axrav/SysAnalytics/backend/db"
	"github.com/axrav/SysAnalytics/backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../.env")
	db.InitPostgres()
	db.InitRedis()
	app := fiber.New()

	app.Use(logger.New())
	// use cors middleware
	app.Use(cors.New(cors.ConfigDefault))
	routes.SetupRoutes(app)

}
