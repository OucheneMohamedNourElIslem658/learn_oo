package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	database "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/database"
	email "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/email"
	oauthproviders "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/oauth_providers"
	utils "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/utils"
)

func init()  {
	database.Init()
	email.Init()
	oauthproviders.Init()
	utils.InitValidators()
}

func main() {
	godotenv.Load();
	host := os.Getenv("HOST")
	
	server := NewServer(fmt.Sprintf("%v:8000", host))
	server.Run()
}