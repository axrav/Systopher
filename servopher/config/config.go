package config

import (
	"fmt"
	"os"

	"github.com/axrav/Systopher/servopher/helpers"
	"github.com/joho/godotenv"
)

func Load() {
	if helpers.EnvExists() {
		fmt.Println("Env exists")
	} else {
		helpers.TakeUserInputAndSave()
	}
	godotenv.Load(helpers.GetEnvPath())
	godotenv.Load("/etc/os-release")
	helpers.ApiKey = os.Getenv("TOKEN")
}
