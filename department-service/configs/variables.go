package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// function to check if .env can be accessible
func LoadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Unable to load the environmental variables: %s\n", err)
	} else {
		log.Println("Environmental variables are loaded successfully ...")
	}
}

// function to fetch the data from the .env
func GetEnv(key string, placeholder string) string {
	value := os.Getenv(key)
	if value == "" {
		return placeholder
	} else {
		return value
	}
}
