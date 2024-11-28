package filestorage

import (
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	publicKey   string
	privateKey   string
	endpointURL string
}

var envs = initAPI()

func initAPI() config {
	godotenv.Load()
	return config{
		publicKey: os.Getenv("IMAGEKIT_PUBLIC_KEY"),
		privateKey: os.Getenv("IMAGEKIT_PRIVATE_KEY"),
		endpointURL: os.Getenv("IMAGEKIT_ENDPOINT_URL"),
	}
}
