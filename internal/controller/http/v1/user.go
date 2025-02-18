package v1

import (
	"github.com/go-chi/chi/v5"
	customMiddleware "komek/internal/controller/http/middleware"
	"komek/internal/dto"
	"komek/internal/mapper"
	"komek/pkg/logger"
	"log"
	"net/http"
)

// todo: get accounts that belong to user with pagination
// todo: make transfer money
// todo: checking on transfer if currency the same on 2 accounts

func (h *Handler) userRoutes(r *chi.Mux) {
	r.Route("/user", func(r chi.Router) {
		// protected
		r.Route("/", func(r chi.Router) {
			r.Use(customMiddleware.Auth(h.tokenMaker))

			r.Delete("/delete", h.userDelete)
			r.Put("/change-password", h.userChangePassword)
			r.Put("/update", h.userUpdate)
			r.Post("/logout", h.userLogout)
			r.Get("/", h.userGet)
			//r.Get("/find", h.usersFind)
		})
		// public
		r.Post("/register", h.userRegister)
		r.Post("/login", h.userLogin)
		r.Post("/refresh-token", h.userRefreshToken)
	})
}

func (h *Handler) userRefreshToken(w http.ResponseWriter, r *http.Request) {
	req := dto.UserRefreshTokensIn{}
	if err := req.ParseAndValidate(r); err != nil {
		h.l.Error("userRefreshToken - ParseAndValidate", logger.Err(err))
		h.Error(w, err, http.StatusBadRequest)
		return
	}

	refreshTokenResponse, err := h.user.RefreshTokens(r.Context(), req)
	if err != nil {
		h.l.Error("userRefreshToken - h.user.RefreshTokens", logger.Err(err))
		h.Error(w, err, http.StatusInternalServerError)
		return
	}

	h.Resp(w, refreshTokenResponse, http.StatusOK)
}

func (h *Handler) userRegister(w http.ResponseWriter, r *http.Request) {
	req := dto.UserRegisterRequest{}
	if err := req.ParseAndValidate(r); err != nil {
		h.l.Error("userRegister - ParseAndValidate", logger.Err(err))
		h.Error(w, err, http.StatusBadRequest)
		return
	}

	user, err := h.user.Register(r.Context(), req)
	if err != nil {
		log.Println("err:", err)
		h.l.Error("userRegister - h.user.Register", logger.Err(err))
		h.Error(w, err, http.StatusInternalServerError)
		return
	}

	h.Resp(w, mapper.ConvUserResponse(user), http.StatusOK)
}

func (h *Handler) userDelete(w http.ResponseWriter, r *http.Request) {
	req := dto.UserDeleteRequest{}

	if err := req.ParseAndValidate(r); err != nil {
		h.l.Error("userDelete - ParseAndValidate", logger.Err(err))
		h.Error(w, err, http.StatusBadRequest)
		return
	}

	payload := h.payload(r)
	err := h.user.Delete(r.Context(), dto.UserDeleteRequest{ID: payload.UserID})
	if err != nil {
		h.l.Error("userDelete - ParseAndValidate", logger.Err(err))
		h.Error(w, err, http.StatusBadRequest)
		return
	}

	h.Resp(w, map[string]any{
		"user_id": payload.UserID,
	}, http.StatusOK)
}

func (h *Handler) userChangePassword(w http.ResponseWriter, r *http.Request) {
	req := dto.UserChangePasswordRequest{}
	if err := req.ParseAndValidate(r); err != nil {
		h.l.Error("userChangePassword - ParseAndValidate", logger.Err(err))
		h.Error(w, err, http.StatusBadRequest)
		return
	}

	payload := h.payload(r)
	req.ID = payload.UserID

	err := h.user.ChangePassword(r.Context(), req)
	if err != nil {
		h.l.Error("userChangePassword - user.ChangePassword", logger.Err(err))
		h.Error(w, err, http.StatusConflict)
		return
	}

	h.Resp(w, map[string]any{
		"status": "success",
	}, http.StatusOK)
}

func (h *Handler) userUpdate(w http.ResponseWriter, r *http.Request) {
	req := dto.UserUpdateRequest{}
	if err := req.ParseAndValidate(r); err != nil {
		h.l.Error("userUpdate - ParseAndValidate", logger.Err(err))
		h.Error(w, err, http.StatusBadRequest)
		return
	}

	payload := h.payload(r)
	req.ID = payload.UserID
	user, err := h.user.Update(r.Context(), req)
	if err != nil {
		h.l.Error("userUpdate - h.user.Update", logger.Err(err))
		h.Error(w, err, http.StatusInternalServerError)
		return
	}

	h.Resp(w, mapper.ConvUserResponse(user), http.StatusOK)
}

func (h *Handler) userLogin(w http.ResponseWriter, r *http.Request) {
	req := dto.UserLoginRequest{}
	if err := req.ParseAndValidate(r); err != nil {
		h.l.Error("userLogin - ParseAndValidate", logger.Err(err))
		h.Error(w, err, http.StatusBadRequest)
		return
	}

	response, err := h.user.Login(r.Context(), req)
	if err != nil {
		h.l.Error("h.user.Login", logger.Err(err))
		h.Error(w, err, http.StatusInternalServerError)
		return
	}

	h.Resp(w, response, http.StatusOK)
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
	req := dto.UserGetRequest{}
	if err := req.ParseAndValidate(r); err != nil {
		h.l.Error("userGet - ParseAndValidate", logger.Err(err))
		h.Error(w, err, http.StatusBadRequest)
		return
	}

	user, err := h.user.Get(r.Context(), req)
	if err != nil {
		h.l.Error("userGet - h.user.Get", logger.Err(err))
		h.Error(w, err, http.StatusInternalServerError)
		return
	}

	h.Resp(w, mapper.ConvUserResponse(user), http.StatusOK)
}

func (h *Handler) usersFind(w http.ResponseWriter, r *http.Request) {
	req := dto.UserRegisterRequest{}
	if err := req.ParseAndValidate(r); err != nil {
		h.l.Error("usersFind - ParseAndValidate", logger.Err(err))
		h.Error(w, err, http.StatusBadRequest)
		return
	}

	h.Resp(w, "success", http.StatusOK)
}
