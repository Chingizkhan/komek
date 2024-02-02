package oauth

import (
	"golang.org/x/oauth2"
)

var (
	Config oauth2.Config
)

func init() {
	Config = oauth2.Config{
		ClientID:     "1ac9df27-ade0-4f03-b434-b17337b6411e",
		ClientSecret: "mysecret",
		RedirectURL:  "http://localhost:8888/callback",
		Scopes:       []string{"offline", "users.write", "users.read", "users.edit", "users.delete"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "http://localhost:9010/oauth2/auth",
			TokenURL: "http://localhost:9010/oauth2/token",
		},
	}
}
