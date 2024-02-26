package app

import (
	"github.com/Chingizkhan/sso_client"
	"github.com/go-chi/chi/v5"
	"golang.org/x/oauth2"
	"komek/config"
	"komek/internal/controller/http/v1"
	"komek/internal/repos/tx"
	"komek/internal/repos/user_repo"
	"komek/internal/service/hasher"
	"komek/internal/service/locker"
	"komek/internal/service/oauth_service"
	"komek/internal/service/token"
	"komek/internal/service/transactional"
	"komek/internal/usecase/banking_uc"
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
	txRepo := tx.NewTX(pg)
	tokenMaker, err := token.NewPasetoMaker("12312312312312312312312312312312")
	if err != nil {
		l.Error("app - Run - token.NewPasetoMaker:", logger.Err(err))
		os.Exit(1)
	}

	sso := sso_client.New(sso_client.Config{
		CookieSecret:   "secret",
		CookieLifetime: 10 * time.Minute,
		OauthAddr:      "http://localhost:8081",
		Oauth2Config: oauth2.Config{
			ClientID:     "38b36b9d-48a8-40fd-9911-ee4462428c58",
			ClientSecret: "mysecret",
			Endpoint: oauth2.Endpoint{
				AuthURL:  "http://localhost:9010/oauth2/auth",
				TokenURL: "http://localhost:9010/oauth2/token",
			},
			RedirectURL: "http://localhost:8082/callback",
			Scopes:      []string{"offline", "users.write", "users.read", "users.edit", "users.delete"},
		},
	})

	go startCron()

	//for _, v := range []int{1, 2} {
	//	v := v
	//	go func() {
	//		testMultiplePay(v)
	//	}()
	//}

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
	userUC := user_uc.New(userRepo, transactionalRepo, hash, tokenMaker)
	bankingUC := banking_uc.New(txRepo)

	// start http server
	r := chi.NewRouter()
	handler := v1.NewHandler(&v1.HandlerParams{
		Logger:            l,
		Cfg:               cfg,
		User:              userUC,
		Banking:           bankingUC,
		CookieSecret:      []byte(cfg.Cookie.Secret),
		OauthServerClient: oauthServerClient,
		Sso:               sso,
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

//type (
//	TransferRequest struct {
//		FromPhone string  `json:"from_phone"`
//		ToPhone   string  `json:"to_phone"`
//		Amount    float64 `json:"amount"`
//		OrderType string  `json:"order_type"`
//		IsDev     bool    `json:"is_dev"`
//	}
//
//	TransferResponse struct {
//		Id string `json:"id"`
//	}
//)
//
//const (
//	path     = "http://localhost:8083/api/v1/integration"
//	transfer = "/wallet/transfer"
//)

//func testMultiplePay(v int) {
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
//	defer cancel()
//
//	js, err := json.Marshal(TransferRequest{
//		FromPhone: "11111111113",
//		ToPhone:   "11111111114",
//		Amount:    10,
//		OrderType: "",
//		IsDev:     true,
//	})
//	req, err := http.NewRequestWithContext(ctx, http.MethodPost, path+transfer, bytes.NewBuffer(js))
//	if err != nil {
//		log.Println("new request:", err)
//		return
//	}
//	response, err := http.DefaultClient.Do(req)
//	if err != nil {
//		log.Println("defaultClient.Do:", err)
//		return
//	}
//	defer response.Body.Close()
//
//	responseJS, err := io.ReadAll(response.Body)
//	if err != nil {
//		log.Println("can not read response:", err)
//		return
//	}
//
//	var transferResponse TransferResponse
//	err = json.Unmarshal(responseJS, &transferResponse)
//	if err != nil {
//		log.Println("yoyo", string(responseJS))
//		log.Println("can not unmarshal response:", err)
//		return
//	}
//
//	log.Printf("%d - transferResponse: %#v", v, transferResponse)
//}
