package helpers

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func EnvExists() bool {
	path := os.Getenv("HOME") + "/.systopher/config.env"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func TakeUserInputAndSave() {
	var uniqueToken string
	var port string
	color.Blue("Enter your Generated Token: ")
	fmt.Scanln(&uniqueToken)
	color.Blue("Enter your Port: ")
	fmt.Scanln(&port)
	err := os.MkdirAll(os.Getenv("HOME")+"/.systopher", 0755)
	if err != nil {
		color.Red("Error creating directory")
	}
	file, err := os.Create(os.Getenv("HOME") + "/.systopher/config.env")
	if err != nil {
		color.Red("Error creating file")
	}
	defer file.Close()
	file.WriteString("TOKEN=" + uniqueToken + "\n" + "PORT=" + port)
	color.Green("Config file created successfully!! in ~/.systopher/config.env")

}

func GetEnvPath() string {
	return os.Getenv("HOME") + "/.systopher/config.env"
}
