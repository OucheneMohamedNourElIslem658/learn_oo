package payment

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	SecretKey string
	BaseURL   string
}

var envs = initAPI()

func initAPI() Config {
	if err := godotenv.Load("../../.env");err != nil {
		log.Fatalf("Error loading .env file")
	}

	return Config{
		SecretKey: os.Getenv("CHARGILY_SECRET_KEY"),
		BaseURL:   "https://pay.chargily.net/test/api/v2",
	}
}
