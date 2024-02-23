package app

import (
	"github.com/go-chi/chi/v5"
	"komek/config"
	"komek/internal/controller/http/v1"
	"komek/internal/repos/store"
	"komek/internal/repos/user_repo"
	"komek/internal/service/banking"
	"komek/internal/service/hasher"
	"komek/internal/service/locker"
	"komek/internal/service/oauth_service"
	"komek/internal/service/transactional"
	"komek/internal/usecase/user_uc"
	"komek/pkg/httpserver"
	"komek/pkg/logger"
	"komek/pkg/postgres"
	"komek/pkg/redis"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// todo: connect with sso server
// todo: finish user endpoints
// todo: token use case
// todo: add auth middleware

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

	cache, err := redis.New(cfg.Redis.Addr, cfg.Redis.Password)
	if err != nil {
		l.Error("app - Run - redis.New:", logger.Err(err))
		os.Exit(1)
	}

	oauthServerClient := oauth_service.New(time.Second*10, cfg.Oauth2Raw.ServiceAddr)
	userRepo := user_repo.New(pg)
	transactionalRepo := transactional.New(pg)
	hash := hasher.New()
	lock := locker.New(cache.Client, cfg.LockTimeout)
	txRepo := store.NewTX(pg)

	err = lock.Lock("data")
	if err != nil {
		log.Println("lock.Lock:", err)
		os.Exit(1)
	}
	err = lock.Unlock("data")
	if err != nil {
		log.Println("lock.Unlock:", err)
		os.Exit(1)
	}

	// get usecases
	userUC := user_uc.New(userRepo, transactionalRepo, hash)
	bankingUC := banking.New(txRepo)

	// start http server
	r := chi.NewRouter()
	handler := v1.NewHandler(&v1.HandlerParams{
		Logger:            l,
		Cfg:               cfg,
		User:              userUC,
		Banking:           bankingUC,
		CookieSecret:      []byte(cfg.Cookie.Secret),
		OauthServerClient: oauthServerClient,
	})
	handler.Register(r, cfg.HTTP.Timeout)
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
