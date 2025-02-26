package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	openAPIMiddleware "github.com/go-openapi/runtime/middleware"
	"komek/config"
	"komek/internal/controller"
	customMiddleware "komek/internal/controller/http/middleware"
	"komek/internal/service/oauth_service"
	"komek/internal/service/token"
	"komek/internal/usecase/client"
	"komek/pkg/logger"
	"net/http"
	"time"
)

type (
	Handler struct {
		l                 logger.ILogger
		cfg               *config.Config
		user              controller.User
		banking           controller.Banking
		client            *client.UseCase
		tokenMaker        token.Maker
		cookieSecret      []byte
		oauthServerClient OauthServerClient
		//sso               sso_client.Client
	}

	HandlerParams struct {
		Logger            logger.ILogger
		Cfg               *config.Config
		User              controller.User
		Banking           controller.Banking
		Client            *client.UseCase
		TokenMaker        token.Maker
		CookieSecret      []byte
		OauthServerClient OauthServerClient
		//Sso               sso_client.Client
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
		client:            p.Client,
		tokenMaker:        p.TokenMaker,
		cookieSecret:      p.CookieSecret,
		oauthServerClient: p.OauthServerClient,
		//sso:               p.Sso,
	}
}

func (h *Handler) Register(r *chi.Mux, timeout time.Duration) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(timeout))
	r.Use(customMiddleware.Cors)
	r.Use(customMiddleware.Logging(h.l))

	opts := openAPIMiddleware.RedocOpts{SpecURL: "swagger.yaml"}
	sh := openAPIMiddleware.Redoc(opts, nil)

	r.Handle("/", http.RedirectHandler("/docs", http.StatusMovedPermanently))
	r.Handle("/docs", sh)
	r.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	h.userRoutes(r)
	h.bankingRoutes(r)
	h.clientRoutes(r)

	//r.With(h.sso.AuthOauth2).Get("/test", h.test)
	//r.With(h.sso.AuthClientCredentials).Get("/test-cc", h.test)
	//r.Get("/callback", h.callback)
}

func (h *Handler) test(w http.ResponseWriter, r *http.Request) {
	h.Resp(w, "okay", http.StatusOK)
}
