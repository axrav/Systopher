package main

import (
	"github.com/axrav/Systopher/backend/db"
	"github.com/axrav/Systopher/backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	// HealthCheck godoc
	// @Summary Show the status of server.
	// @Description get the status of server.
	// @Tags root
	// @Accept */*
	// @Produce json
	// @Success 200 {object} map[string]interface{}
	// @Router / [get]
	godotenv.Load("../.env")

	db.InitPostgres()
	db.InitRedis()
	app := fiber.New()

	app.Use(logger.New())
	// use cors middleware
	app.Use(cors.New(cors.ConfigDefault))
	routes.SetupRoutes(app)

}
