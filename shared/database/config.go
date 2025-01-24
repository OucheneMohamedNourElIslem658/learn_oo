package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

var envs = initAPI()

func initAPI() config {
	if err := godotenv.Load();err != nil {
		fmt.Println("here 1")
		log.Fatalf("Error loading .env file")
	}
	fmt.Println(os.Getenv("DB_USER"))
	return config{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
	}
}

func (config config) getDatabaseDSN() string {
	return fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v",
		config.Host,
		config.User,
		config.Password,
		config.Name,
		config.Port,
	)
}
