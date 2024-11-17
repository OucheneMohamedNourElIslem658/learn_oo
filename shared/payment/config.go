package payment

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	SecretKey string
	BaseURL   string
}

var envs = initAPI()

func initAPI() Config {
	godotenv.Load()
	return Config{
		SecretKey: os.Getenv("CHARGILY_SECRET_KEY"),
		BaseURL:   "https://pay.chargily.net/test/api/v2",
	}
}
