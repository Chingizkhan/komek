package oidc

import (
	"fmt"
	"komek/internal/service/oauth"
	"komek/pkg/cookies"
	"komek/pkg/state"
	"log"
	"net/http"
	"time"
)

var (
	CookieName = "oauth_state"
)

func LoginFlow(w http.ResponseWriter, r *http.Request, secret []byte) error {
	st, err := state.Generate()
	if err != nil {
		return fmt.Errorf("state.Generate - %w", err)
	}
	log.Println("generated state:", string(st))

	expiration := time.Now().Add(10 * time.Minute)
	cookie := http.Cookie{
		Name:     CookieName,
		Value:    string(st),
		Expires:  expiration,
		Secure:   true,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	err = cookies.WriteSigned(w, cookie, secret)
	if err != nil {
		return fmt.Errorf("cookies.WriteSigned - %w", err)
	}

	loginUrl := oauth.Config.AuthCodeURL(string(st))
	//api_util.RenderResponse(w, loginUrl, http.StatusOK)

	http.Redirect(w, r, loginUrl, http.StatusTemporaryRedirect)

	return nil
}
