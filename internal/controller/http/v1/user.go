package v1

import (
	"context"
	"github.com/go-chi/chi/v5"
	customMiddleware "komek/internal/controller/http/middleware"
	"komek/internal/dto"
	"komek/pkg/logger"
	"net/http"
)

type (
	UserUseCase interface {
		Register(ctx context.Context, req dto.UserRegisterRequest) error
		Delete(ctx context.Context, req dto.UserDeleteRequest) error
		ChangePassword(ctx context.Context, req dto.UserChangePasswordRequest) error
		Update(ctx context.Context, req dto.UserUpdateRequest) error
		// todo: make it later
		Login(ctx context.Context) error
		Logout(ctx context.Context) error
	}
)

func (h *Handler) userRoutes(r *chi.Mux) {
	r.Route("/user", func(r chi.Router) {
		// protected
		r.Route("/", func(r chi.Router) {
			// todo: use middleware with jwt auth
			r.Use(customMiddleware.AuthOauth2(h.cookieSecret))

			r.Delete("/delete", h.userDelete)
			r.Put("/change-password", h.userChangePassword)
			r.Put("/update", h.userUpdate)
			r.Post("/logout", h.userLogout)
			r.Get("/{:id}", h.userGet)
			r.Get("/find", h.usersFind)
		})
		// public
		r.Post("/register", h.userRegister)
		r.Post("/login", h.userLogin)
	})
}

func (h *Handler) userRegister(w http.ResponseWriter, r *http.Request) {
	req := dto.UserRegisterRequest{}
	if err := req.ParseAndValidate(r); err != nil {
		h.l.Error("userRegister - ParseAndValidate", logger.Err(err))
		h.Err(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.Resp(w, "success", http.StatusOK)
}

func (h *Handler) userDelete(w http.ResponseWriter, r *http.Request) {
	req := dto.UserDeleteRequest{}
	if err := req.ParseAndValidate(r); err != nil {
		h.l.Error("userDelete - ParseAndValidate", logger.Err(err))
		h.Err(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.Resp(w, "success", http.StatusOK)
}

func (h *Handler) userChangePassword(w http.ResponseWriter, r *http.Request) {
	req := dto.UserChangePasswordRequest{}
	if err := req.ParseAndValidate(r); err != nil {
		h.l.Error("userChangePassword - ParseAndValidate", logger.Err(err))
		h.Err(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.Resp(w, "success", http.StatusOK)
}

func (h *Handler) userUpdate(w http.ResponseWriter, r *http.Request) {
	req := dto.UserUpdateRequest{}
	if err := req.ParseAndValidate(r); err != nil {
		h.l.Error("userUpdate - ParseAndValidate", logger.Err(err))
		h.Err(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.Resp(w, "success", http.StatusOK)
}

func (h *Handler) userLogin(w http.ResponseWriter, r *http.Request) {
	req := dto.UserLoginRequest{}
	if err := req.ParseAndValidate(r); err != nil {
		h.l.Error("userLogin - ParseAndValidate", logger.Err(err))
		h.Err(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.Resp(w, "success", http.StatusOK)
}

func (h *Handler) userLogout(w http.ResponseWriter, r *http.Request) {
	req := dto.UserRegisterRequest{}
	if err := req.ParseAndValidate(r); err != nil {
		h.l.Error("userRegister - ParseAndValidate", logger.Err(err))
		h.Err(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.Resp(w, "success", http.StatusOK)
}

func (h *Handler) userGet(w http.ResponseWriter, r *http.Request) {
	req := dto.UserRegisterRequest{}
	if err := req.ParseAndValidate(r); err != nil {
		h.l.Error("userRegister - ParseAndValidate", logger.Err(err))
		h.Err(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.Resp(w, "success", http.StatusOK)
}

func (h *Handler) usersFind(w http.ResponseWriter, r *http.Request) {
	req := dto.UserRegisterRequest{}
	if err := req.ParseAndValidate(r); err != nil {
		h.l.Error("userRegister - ParseAndValidate", logger.Err(err))
		h.Err(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.Resp(w, "success", http.StatusOK)
}
