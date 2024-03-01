package v1

import (
	"github.com/Chingizkhan/sso_client"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"komek/config"
	customMiddleware "komek/internal/controller/http/middleware"
	"komek/internal/service/oauth_service"
	"komek/pkg/logger"
	"net/http"
	"time"
)

type (
	Handler struct {
		l                 logger.ILogger
		cfg               *config.Config
		user              User
		banking           Banking
		cookieSecret      []byte
		oauthServerClient OauthServerClient
		sso               sso_client.Client
	}

	HandlerParams struct {
		Logger            logger.ILogger
		Cfg               *config.Config
		User              User
		Banking           Banking
		CookieSecret      []byte
		OauthServerClient OauthServerClient
		Sso               sso_client.Client
	}

	OauthServerClient interface {
		Introspect(tok string) (oauth_service.IntrospectResponse, error)
	}
)

func NewHandler(p *HandlerParams) *Handler {
	return &Handler{
		l:                 p.Logger,
		cfg:               p.Cfg,
		user:              p.User,
		banking:           p.Banking,
		cookieSecret:      p.CookieSecret,
		oauthServerClient: p.OauthServerClient,
		sso:               p.Sso,
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
	h.bankingRoutes(r)

	r.With(h.sso.AuthOauth2).Get("/test", h.test)
	r.With(h.sso.AuthClientCredentials).Get("/test-cc", h.test)
	r.Get("/callback", h.callback)
}

func (h *Handler) test(w http.ResponseWriter, r *http.Request) {
	h.Resp(w, "okay", http.StatusOK)
}
