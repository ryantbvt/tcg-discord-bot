package framework

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DiscToken string `json:"disc_token"`
}

func LoadEnv() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	// Add below for new tokens
	config := &Config{
		DiscToken: os.Getenv("DISCORD_TOKEN"),
	}

	log.Println("Successfully loaded secrets")
	return config
}
