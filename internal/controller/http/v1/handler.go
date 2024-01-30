package v1

import (
	"github.com/go-chi/chi/v5"
	"komek/config"
	"komek/internal/controller/http/middleware"
	"komek/pkg/logger"
)

type (
	Handler struct {
		l      logger.ILogger
		cfg    *config.Config
		userUC UserUseCase
	}

	HandlerParams struct {
		Logger logger.ILogger
		Cfg    *config.Config
		UserUC UserUseCase
	}
)

func NewHandler(p *HandlerParams) *Handler {
	return &Handler{
		p.Logger,
		p.Cfg,
		p.UserUC,
	}
}

func (h *Handler) Register(r *chi.Mux) {
	r.Use(middleware.Logging(h.l))

	h.userRoutes(r)
}
