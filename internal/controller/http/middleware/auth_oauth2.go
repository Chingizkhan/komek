package middleware

import (
	"komek/internal/service/oauth/oidc"
	"komek/pkg/http/api_util"
	"komek/pkg/token"
	"log"
	"net/http"
)

func AuthOauth2(secret []byte) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			// check token exists
			tok := token.FromHttpRequest(r)
			log.Println("token:", tok)
			if tok == "" {
				// return login page
				log.Println("return login page")
				err := oidc.LoginFlow(w, r, secret)
				if err != nil {
					api_util.RenderErrorResponse(w, "can not show login flow", http.StatusInternalServerError)
					return
				}
				return
			}

			// introspect token
			if err := oidc.Introspect(r.Context(), tok); err != nil {
				api_util.RenderErrorResponse(w, err.Error(), http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
