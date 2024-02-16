package oauth

import (
	"golang.org/x/oauth2"
)

var (
	Config oauth2.Config
)

func init() {
	Config = oauth2.Config{
		ClientID:     "321ffb75-c57b-4f11-866c-63fd1a561ddb",
		ClientSecret: "mysecret",
		RedirectURL:  "http://localhost:8888/callback",
		Scopes:       []string{"offline", "users.write", "users.read", "users.edit", "users.delete"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "http://localhost:9010/oauth2/auth",
			TokenURL: "http://localhost:9010/oauth2/token",
		},
	}
}
