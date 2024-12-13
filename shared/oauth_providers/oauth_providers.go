package oauthproviders

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
)

type Provider struct {
	Config       *oauth2.Config
	UserInfoURL  string
	EmailInfoURL string
}

type Providers map[string]Provider

var Instance Providers

func Init() {
	Instance = Providers{
		"google": {
			Config: &oauth2.Config{
				ClientID:     envs.googleClientID,
				ClientSecret: envs.googleClientSecret,
				RedirectURL:  "https://learn-oo-api.onrender.com/api/v1/users/auth/oauth/google/callback",
				Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
				Endpoint:     google.Endpoint,
			},
			UserInfoURL: "https://www.googleapis.com/oauth2/v2/userinfo?alt=json",
		},
		"facebook": {
			Config: &oauth2.Config{
				ClientID:     envs.facebookClientID,
				ClientSecret: envs.facebookClientSecret,
				RedirectURL:  "https://learn-oo-api.onrender.com/api/v1/users/auth/oauth/facebook/callback/",
				Scopes:       []string{"public_profile", "email"},
				Endpoint:     facebook.Endpoint,
			},
			UserInfoURL: "https://graph.facebook.com/me?fields=id,name,email,picture",
		},
	}
}

func IsSupportedProvider(provider string) bool {
	_, exists := Instance[provider]
	return exists
}
