package main

import (
	database "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/database"
	email "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/email"
	oauthproviders "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/oauth_providers"
	utils "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/utils"
)

func init() {
	database.Init()
	email.Init()
	oauthproviders.Init()
	utils.InitValidators()
}

func main() {
	server := NewServer("0.0.0.0:8000")
	server.Run()
}
