package app

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"komek/config"
	"komek/internal/controller/http/v1"
	"komek/internal/repos/user_repo"
	"komek/pkg/httpserver"
	"komek/pkg/logger"
	"komek/pkg/postgres"
	"log"
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

	ctx := context.Background()

	tx, err := pg.Pool.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		return
	}
	defer tx.Rollback(ctx)

	userRepo := user_repo.New(pg)
	uu, err := uuid.Parse("9034fda7-543e-48da-a463-973c70dbbecd")
	resp, err := userRepo.UpdateLogin(ctx, tx, uu, "login")
	respPass, err := userRepo.UpdatePasswordHash(ctx, tx, uu, "passHash")
	_ = resp
	_ = respPass
	if err != nil {
		log.Println("err update", err)
		return
	}

	tx.Commit(ctx)

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
