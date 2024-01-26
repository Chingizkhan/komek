package middleware

import (
	"fmt"
	"github.com/gorilla/mux"
	"komek/pkg/logger"
	"net/http"
)

func Logging(l logger.ILogger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			l.Info(fmt.Sprintf("[%s]: %s", r.Method, r.RequestURI)) //info("url", r.RequestURI),
			//info("url", r.RequestURI),

			next.ServeHTTP(w, r)
		})
	}
}
