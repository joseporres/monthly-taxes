package config

// get environment variable
import (
	"github.com/joho/godotenv"
)


func LoadEnvFile() {
	//load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}