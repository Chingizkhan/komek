package middleware

import (
	"context"
	response "komek/internal/controller/http"
	tokenService "komek/internal/service/token"
	"komek/pkg/token"
	"net/http"
)

var (
	AuthorizationPayloadKey = "authorization_payload"
)

func Auth(tokenMaker tokenService.Maker) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			// check and validate token
			accessToken, err := token.Get(r)
			if err != nil {
				response.Err(w, err.Error(), http.StatusUnauthorized)
				return
			}
			payload, err := tokenMaker.VerifyToken(accessToken)
			if err != nil {
				response.Err(w, err.Error(), http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), AuthorizationPayloadKey, payload)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
