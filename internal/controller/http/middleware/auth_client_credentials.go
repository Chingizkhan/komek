package middleware

import (
	"context"
	"fmt"
	"komek/internal/service/oauth/client_credentials"
	"komek/pkg/cookies"
	"komek/pkg/http/api_util"
	"komek/pkg/token"
	"log"
	"net/http"
	"time"
)

func AuthClientCredentials() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			// check token exists
			accessToken := token.FromHttpRequest(r)
			if accessToken == "" {
				api_util.RenderErrorResponse(w, "access token is not set", http.StatusUnauthorized)
				return
			}

			// introspect token
			if err := client_credentials.Introspect(r.Context(), accessToken); err != nil {
				api_util.RenderErrorResponse(w, err.Error(), http.StatusUnauthorized)
				log.Println("introspecting token error", err)
				return
			}
			log.Println("introspecting token success")
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}

func authenticate(w http.ResponseWriter, clientID, clientSecret string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	accessToken, err := client_credentials.Auth(ctx, clientID, clientSecret)
	if err != nil {
		log.Println("client_credentials.Auth err:", err)
		api_util.RenderErrorResponse(w, "can not authenticate by client credentials", http.StatusUnauthorized)
		return
	}
	err = writeAccessToken(w, accessToken)
	if err != nil {
		log.Println("writeAccessToken:", err)
		api_util.RenderErrorResponse(w, "can not give token back", http.StatusInternalServerError)
		return
	}
}

func findAccessToken(r *http.Request) (string, error) {
	accessToken := token.FromHttpRequest(r)
	if accessToken != "" {
		log.Println("found access token in header", accessToken)
		return accessToken, nil
	}

	return accessToken, nil
}

func writeAccessToken(w http.ResponseWriter, accessToken string) error {
	expiration := time.Now().Add(10 * time.Minute)
	cookie := http.Cookie{
		Name:     "Access-Token",
		Value:    accessToken,
		Expires:  expiration,
		Secure:   true,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	err := cookies.WriteSigned(w, cookie, []byte("secret_cookie"))
	if err != nil {
		return fmt.Errorf("cookies.WriteSigned: %w", err)
	}
	return nil
}
