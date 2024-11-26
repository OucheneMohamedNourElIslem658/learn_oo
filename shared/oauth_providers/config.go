package oauthproviders

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	googleClientID       string
	googleClientSecret   string
	githubClientID       string
	githubClientSecret   string
	facebookClientID     string
	facebookClientSecret string
	twitterClientID      string
	twitterClientSecret  string
	microsoftClientID    string
	microsoftClientSecret string
	callbackURL          string
}

var envs = initAPI()

func initAPI() config {
	if err := godotenv.Load();err != nil {
		fmt.Println("here")
		log.Fatalf("Error loading .env file")
	}

	return config{
		googleClientID:       os.Getenv("GOOGLE_CLIENT_ID"),
		googleClientSecret:   os.Getenv("GOOGLE_CLIENT_SECRET"),
		githubClientID:       os.Getenv("GITHUB_CLIENT_ID"),
		githubClientSecret:   os.Getenv("GITHUB_CLIENT_SECRET"),
		facebookClientID:     os.Getenv("FACEBOOK_CLIENT_ID"),
		facebookClientSecret: os.Getenv("FACEBOOK_CLIENT_SECRET"),
		twitterClientID:      os.Getenv("TWITTER_CLIENT_ID"),
		twitterClientSecret:  os.Getenv("TWITTER_CLIENT_SECRET"),
		microsoftClientID:    os.Getenv("MICROSOFT_CLIENT_ID"),
		microsoftClientSecret: os.Getenv("MICROSOFT_CLIENT_SECRET"),
		callbackURL:          os.Getenv("CALLBACK_URL"),
	}
}