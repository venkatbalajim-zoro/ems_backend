package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(data string) {
	envMap, err := godotenv.Unmarshal(data)
	for key, value := range envMap {
		os.Setenv(key, value)
	}
	if err != nil {
		log.Fatalf("Unable to load the environmental variables: %s\n", err)
	} else {
		log.Println("Environmental variables are loaded successfully ...")
	}
}

func GetEnv(key string, placeholder string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	} else {
		return placeholder
	}
}
