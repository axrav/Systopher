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
	app := fiber.New()
	db.Init()
	// helpers.DataChannel = make(chan types.ServerData)
	//helpers.ServerChannel = make(chan []string)
	//go helpers.GetData(helpers.ServerChannel, helpers.DataChannel)
	app.Use(logger.New())
	// use cors middleware
	app.Use(cors.New(cors.ConfigDefault))
	routes.SetupRoutes(app)

}
