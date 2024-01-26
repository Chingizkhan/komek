package middleware

import (
	"context"
	"github.com/gorilla/mux"
	"komek/internal/controller/http/api_util"
	"komek/internal/domain/jwt"
	"komek/pkg/logger"
	"komek/pkg/rsa"
	"net/http"
	"strings"
)

type FuncRetrospect func(ctx context.Context, accessToken string) (*gocloak.IntroSpectTokenResult, error)

func JWT(l logger.ILogger, publicKey string, retrospect FuncRetrospect) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			jwtToken := jwt.FromRequestHeader(r)

			pk, err := rsa.ParseKeycloakRSAPublicKey(publicKey)

			err = jwtToken.Parse(pk)
			if err != nil {
				if strings.Contains(err.Error(), "token is expired") {
					l.Error("jwt middleware parse:", logger.Err(err))
					api_util.RenderErrorResponse(w, "token is expired", http.StatusConflict)
					return
				}
				l.Error("jwt middleware parse:", logger.Err(err))
				api_util.RenderErrorResponse(w, "wrong parsing jwt", http.StatusConflict)
				return
			}

			ctx := r.Context()

			r = jwtToken.SetClaims(ctx, r)

			token, err := retrospect(ctx, jwtToken.Raw)
			if err != nil {
				l.Error("jwt middleware retrospect:", logger.Err(err))
				api_util.RenderErrorResponse(w, "wrong jwt", http.StatusConflict)
				return
			}
			if !*token.Active {
				api_util.RenderErrorResponse(w, "token is not active", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
