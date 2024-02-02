package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"komek/config"
	customMiddleware "komek/internal/controller/http/middleware"
	"komek/internal/service/oauthServer"
	"komek/pkg/logger"
	"time"
)

type (
	Handler struct {
		l                 logger.ILogger
		cfg               *config.Config
		userUC            UserUseCase
		cookieSecret      []byte
		oauthServerClient OauthServerClient
	}

	HandlerParams struct {
		Logger            logger.ILogger
		Cfg               *config.Config
		UserUC            UserUseCase
		CookieSecret      []byte
		OauthServerClient OauthServerClient
	}

	OauthServerClient interface {
		Introspect(tok string) (oauthServer.IntrospectResponse, error)
	}
)

func NewHandler(p *HandlerParams) *Handler {
	return &Handler{
		l:            p.Logger,
		cfg:          p.Cfg,
		userUC:       p.UserUC,
		cookieSecret: p.CookieSecret,
	}
}

func (h *Handler) Register(r *chi.Mux, timeout time.Duration) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(timeout))
	r.Use(customMiddleware.Cors)
	r.Use(customMiddleware.Logging(h.l))

	h.userRoutes(r)
	r.Get("/callback", h.callback(h.cookieSecret))
}
