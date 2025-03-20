package config

import (
	"log"
	_ "os"

	"github.com/joho/godotenv"
)

// LoadEnvVariables carga las variables de entorno desde el .env
func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}
}
