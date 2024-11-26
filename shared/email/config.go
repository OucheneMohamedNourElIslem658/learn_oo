package email

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	email     string
	password string
}

var mailerConfig = initAPI()

func initAPI() config {
	if err := godotenv.Load();err != nil {
		log.Fatalf("Error loading .env file")
	}
	return config{
		email:     os.Getenv("EMAIL"),
		password: os.Getenv("PASSWORD"),
	}
}