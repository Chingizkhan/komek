package v1

import (
	"context"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"komek/config"
	"komek/internal/controller/http/middleware"
	"komek/internal/domain/word"
	"komek/internal/service/identity_manager"
	"komek/internal/usecase/user_managment"
	"komek/pkg/logger"
	"net/http"
)

type (
	IWordUseCase interface {
		Get(ctx context.Context, value string, userId uuid.UUID) (word.Word, error)
		Save(ctx context.Context, word word.Word) error
		Update(ctx context.Context, oldValue string, word word.Word) (word.Word, error)
		Delete(ctx context.Context, value string, userId uuid.UUID) error
	}

	IUserUseCase interface {
		Register(ctx context.Context, request user_managment.RegisterRequest) (*user_managment.RegisterResponse, error)
		SignIn(ctx context.Context, request user_managment.SignInRequest) (*user_managment.SignInResponse, error)
		Logout(ctx context.Context, request user_managment.LogoutRequest) error
	}

	ITokenUseCase interface {
		Retrospect(ctx context.Context, accessToken string) (*gocloak.IntroSpectTokenResult, error)
		RefreshTokens(ctx context.Context, refreshToken string) (*identity_manager.UserWithJWTResponse, error)
	}

	Handler struct {
		wordUC  IWordUseCase
		userUC  IUserUseCase
		tokenUC ITokenUseCase
		l       logger.ILogger
		cfg     *config.Config
	}
)

func NewHandler(wordUC IWordUseCase, userUC IUserUseCase, tokenUC ITokenUseCase, l logger.ILogger, cfg *config.Config) *Handler {
	return &Handler{wordUC, userUC, tokenUC, l, cfg}
}

func (h *Handler) Register(r *mux.Router) {
	r.Use(middleware.Logging(h.l))

	h.publicRoutes(r)
	h.protectedRoutes(r)
}

func (h *Handler) publicRoutes(r *mux.Router) {
	r.HandleFunc("/v1/auth/register", h.register).Methods(http.MethodPost)
	r.HandleFunc("/v1/auth/sign-in", h.sigIn).Methods(http.MethodPost)
	r.HandleFunc("/v1/auth/callback", h.callback).Methods(http.MethodPost)
	r.HandleFunc("/v1/auth/refresh", h.refreshTokens).Methods(http.MethodPost)
}

func (h *Handler) protectedRoutes(r *mux.Router) {
	api := r.PathPrefix("/v1").Subrouter()
	api.Use(middleware.JWT(h.l, h.cfg.RealmRS256PublicKey, h.tokenUC.Retrospect))

	api.HandleFunc("/word/{word}", h.get).Methods(http.MethodGet)
	api.HandleFunc("/word/add", h.add).Methods(http.MethodPost)
	api.HandleFunc("/word/remove", h.remove).Methods(http.MethodDelete)
	api.HandleFunc("/word/update", h.update).Methods(http.MethodPut)
	api.HandleFunc("/auth/logout", h.logout).Methods(http.MethodPost)
}
