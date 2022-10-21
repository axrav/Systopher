package main

import (
	"fmt"
	"os"

	"github.com/axrav/Systopher/micro/helpers"
	"github.com/axrav/Systopher/micro/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if helpers.EnvExists() {
		fmt.Println("Env exists")
	} else {
		helpers.TakeUserInputAndSave()
	}
	godotenv.Load(helpers.GetEnvPath())
	godotenv.Load("/etc/os-release")
	helpers.ApiKey = os.Getenv("TOKEN")

	routes.SetupRoutes(fiber.New())
}
