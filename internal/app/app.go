package app

import (
	"github.com/go-chi/chi/v5"
	"komek/config"
	"komek/internal/controller/http/v1"
	"komek/pkg/httpserver"
	"komek/pkg/logger"
	"komek/pkg/postgres"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config, l *logger.Logger) {
	pg, err := postgres.New(
		cfg.PG.DSN(),
		postgres.MaxPoolSize(cfg.PG.PoolMax),
	)
	if err != nil {
		l.Error("app - Run - postgres.New:", logger.Err(err))
		os.Exit(1)
	}
	defer pg.Close()

	// get usecases
	//wordUC := worduc.New(wordRepo.New(pg))
	//userUC := useruc.New(im)
	//tokenUC := tokenuc.New(im)

	// start http server
	r := chi.NewRouter()
	handler := v1.NewHandler(l, cfg)
	handler.Register(r)
	httpServer := httpserver.New(
		r,
		httpserver.Port(cfg.HTTP.Port),
	)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal:", slog.String("signal", s.String()))
	case err := <-httpServer.Notify():
		l.Error("app - Run - http_server.Notify:", logger.Err(err))
	}

	err = httpServer.Shutdown()
	if err != nil {
		l.Error("app - Run - httpServer.Shutdown:", logger.Err(err))
		return
	}
}
