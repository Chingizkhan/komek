package app

import (
	"github.com/go-chi/chi/v5"
	"komek/config"
	"komek/internal/controller/grpc"
	"komek/internal/controller/http/v1"
	accountRepo "komek/internal/domain/account/repository"
	accountSrv "komek/internal/domain/account/service"
	clientRepo "komek/internal/domain/client/repository"
	clientSrv "komek/internal/domain/client/service"
	fundraise_repo "komek/internal/domain/fundraise/repository"
	fundraise_service "komek/internal/domain/fundraise/service"
	operation_repo "komek/internal/domain/operation/repository"
	operation_service "komek/internal/domain/operation/service"
	transaction_repo "komek/internal/domain/transaction/repository"
	transaction_service "komek/internal/domain/transaction/service"
	userRepo "komek/internal/domain/user/repository"
	userSrv "komek/internal/domain/user/service"
	"komek/internal/repo/session_repo"
	banking "komek/internal/service/banking/service"
	"komek/internal/service/banking_old"
	"komek/internal/service/hasher"
	"komek/internal/service/identity"
	"komek/internal/service/locker"
	"komek/internal/service/oauth_service"
	"komek/internal/service/token"
	"komek/internal/service/transactional"
	"komek/internal/usecase/banking_uc"
	"komek/internal/usecase/client"
	"komek/internal/usecase/fundraise"
	"komek/internal/usecase/user"
	"komek/pkg/grpcserver"
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
	userRepository := userRepo.New(pg)
	clientRepository := clientRepo.New(pg)
	accountRepository := accountRepo.New(pg)
	sessionRepo := session_repo.New(pg)
	transactionalRepo := transactional.New(pg)
	im := identity.NewIdentityManager("localhost:8181", "komek", "", "")
	operationRepo := operation_repo.New(pg)
	transactionRepo := transaction_repo.New(pg)
	fundraiseRepo := fundraise_repo.New(pg)

	bankingOldService, err := banking_old.New(cfg.BankingService.Addr, cfg.BankingService.EnableTLS)
	if err != nil {
		l.Error("app - Run - bankingService:", logger.Err(err))
		os.Exit(1)
	}
	_ = bankingOldService

	hash := hasher.New()
	lock := locker.New(cache.Client, cfg.LockTimeout)
	tokenMaker, err := token.NewPasetoMaker("12312312312312312312312312312312")
	if err != nil {
		l.Error("app - Run - token.NewPasetoMaker:", logger.Err(err))
		os.Exit(1)
	}

	//sso := sso_client.New(sso_client.Config{
	//	CookieSecret:   "secret",
	//	CookieLifetime: 10 * time.Minute,
	//	OauthAddr:      "http://localhost:8081",
	//	Oauth2Config: oauth2.Config{
	//		ClientID:     "7b51fdc9-23b9-40d0-bce8-f887f1bb8dcb",
	//		ClientSecret: "mysecret",
	//		Endpoint: oauth2.Endpoint{
	//			AuthURL:  "http://localhost:9010/oauth2/auth",
	//			TokenURL: "http://localhost:9010/oauth2/token",
	//		},
	//		RedirectURL: "http://localhost:8888/callback",
	//		Scopes:      []string{"offline"},
	//	},
	//})

	go startCron()

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

	userService := userSrv.New(userRepository)
	clientService := clientSrv.New(clientRepository)
	accountService := accountSrv.New(accountRepository)
	operationService := operation_service.New(operationRepo)
	transactionService := transaction_service.New(transactionRepo)
	bankingService := banking.New(operationService, transactionService, accountRepository)
	fundraiseService := fundraise_service.New(fundraiseRepo)

	// get usecases
	userUC := user.New(userService, accountService, transactionalRepo, hash, sessionRepo, im, tokenMaker, cfg.AccessTokenLifetime, cfg.RefreshTokenLifetime)
	clientUC := client.New(clientService, fundraiseService, accountService, transactionalRepo)
	bankingUC := banking_uc.New(transactionalRepo, bankingService, accountService, fundraiseService)
	fundraiseUC := fundraise.New(fundraiseService, accountService, clientService, transactionalRepo)

	// start http server
	r := chi.NewRouter()
	handler := v1.NewHandler(&v1.HandlerParams{
		Logger:            l,
		Cfg:               cfg,
		User:              userUC,
		Banking:           bankingUC,
		Client:            clientUC,
		Fundraise:         fundraiseUC,
		TokenMaker:        tokenMaker,
		CookieSecret:      []byte(cfg.Cookie.Secret),
		OauthServerClient: oauthServerClient,
		//Sso:               sso,
	})
	handler.Register(r, cfg.HTTP.Timeout)
	httpServer := httpserver.New(
		r,
		httpserver.Port(cfg.HTTP.Port),
	)
	l.Info("http server started", slog.String("env", cfg.Log.Level), slog.String("port", cfg.HTTP.Port))

	// start grpc server
	server := grpc.Register(l, userUC)
	grpcServer := grpcserver.New(server, cfg.GRPC.Port)
	l.Info("grpc server started", slog.String("env", cfg.Log.Level), slog.String("port", cfg.GRPC.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal:", slog.String("signal", s.String()))
	case err = <-httpServer.Notify():
		l.Error("app - Run - http_server.Notify:", logger.Err(err))
	case err = <-grpcServer.Notify():
		l.Error("app - Run - grpc_server.Notify:", logger.Err(err))
	}

	// shutdown
	grpcServer.Shutdown()

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
