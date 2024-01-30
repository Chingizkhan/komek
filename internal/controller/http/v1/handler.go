package v1

import (
	"github.com/go-chi/chi/v5"
	"komek/config"
	"komek/internal/controller/http/middleware"
	"komek/pkg/logger"
)

type (
	Handler struct {
		l   logger.ILogger
		cfg *config.Config
	}
)

func NewHandler(l logger.ILogger, cfg *config.Config) *Handler {
	return &Handler{l, cfg}
}

func (h *Handler) Register(r *chi.Mux) {
	r.Use(middleware.Logging(h.l))

	h.userRoutes(r)
}
