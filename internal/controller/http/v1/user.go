package v1

import (
	"github.com/go-chi/chi/v5"
	customMiddleware "komek/internal/controller/http/middleware"
	"komek/internal/dto"
	"komek/internal/service/token"
	"komek/pkg/logger"
	"net/http"
)

func (h *Handler) userRoutes(r *chi.Mux) {
	r.Route("/user", func(r chi.Router) {
		// protected
		r.Route("/", func(r chi.Router) {
			r.Use(customMiddleware.Auth(h.tokenMaker))
			// todo: use middleware with jwt auth
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

func (h *Handler) payload(r *http.Request) *token.Payload {
	return r.Context().Value(customMiddleware.AuthorizationPayloadKey).(*token.Payload)
}

func (h *Handler) userDelete(w http.ResponseWriter, r *http.Request) {
	req := dto.UserDeleteRequest{}

	if err := req.ParseAndValidate(r); err != nil {
		h.l.Error("userDelete - ParseAndValidate", logger.Err(err))
		h.Err(w, err.Error(), http.StatusBadRequest)
		return
	}

	payload := h.payload(r)
	err := h.user.Delete(r.Context(), dto.UserDeleteRequest{ID: payload.UserID})
	if err != nil {
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
