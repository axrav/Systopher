package main

import (
	"github.com/axrav/Systopher/servopher/config"
	"github.com/axrav/Systopher/servopher/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.Load()
	routes.SetupRoutes(fiber.New())
}
