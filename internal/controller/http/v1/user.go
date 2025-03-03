package v1

import (
	"github.com/go-chi/chi/v5"
	customMiddleware "komek/internal/controller/http/middleware"
	"komek/internal/domain/user/entity"
	"komek/internal/mapper"
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
	req := entity.RefreshTokensIn{}
	if err := req.ParseHttpBody(r); err != nil {
		h.Error(w, err, http.StatusBadRequest, "userRefreshToken - ParseAndValidate")
		return
	}

	refreshTokenResponse, err := h.user.RefreshTokens(r.Context(), req)
	if err != nil {
		h.Error(w, err, http.StatusInternalServerError, "userRefreshToken - h.user.RefreshTokens")
		return
	}

	h.Resp(w, refreshTokenResponse, http.StatusOK)
}

func (h *Handler) userRegister(w http.ResponseWriter, r *http.Request) {
	req := entity.RegisterIn{}
	if err := req.ParseHttpBody(r); err != nil {
		h.Error(w, err, http.StatusBadRequest, "userRegister - Parse")
		return
	}

	user, err := h.user.Register(r.Context(), req)
	if err != nil {
		h.Error(w, err, http.StatusInternalServerError, "userRegister - h.user.Register")
		return
	}

	h.Resp(w, user.ToResponse(), http.StatusOK)
}

func (h *Handler) userDelete(w http.ResponseWriter, r *http.Request) {
	req := entity.DeleteIn{}

	if err := req.ParseHttpBody(r); err != nil {
		h.Error(w, err, http.StatusBadRequest, "userDelete - ParseAndValidate")
		return
	}

	payload := h.payload(r)
	req.ID = payload.UserID

	err := h.user.Delete(r.Context(), req)
	if err != nil {
		h.Error(w, err, http.StatusBadRequest, "userDelete - ParseAndValidate")
		return
	}

	h.Resp(w, map[string]any{
		"user_id": payload.UserID,
	}, http.StatusOK)
}

func (h *Handler) userChangePassword(w http.ResponseWriter, r *http.Request) {
	req := entity.ChangePasswordIn{}
	if err := req.ParseHttpBody(r); err != nil {
		h.Error(w, err, http.StatusBadRequest, "userChangePassword - ParseAndValidate")
		return
	}

	payload := h.payload(r)
	req.ID = payload.UserID

	err := h.user.ChangePassword(r.Context(), req)
	if err != nil {
		h.Error(w, err, http.StatusConflict, "userChangePassword - user.ChangePassword")
		return
	}

	h.Resp(w, map[string]any{
		"status": "success",
	}, http.StatusOK)
}

func (h *Handler) userUpdate(w http.ResponseWriter, r *http.Request) {
	req := entity.UpdateIn{}
	if err := req.ParseHttpBody(r); err != nil {
		h.Error(w, err, http.StatusBadRequest, "userUpdate - ParseAndValidate")
		return
	}

	payload := h.payload(r)
	req.ID = payload.UserID
	response, err := h.user.Update(r.Context(), req)
	if err != nil {
		h.Error(w, err, http.StatusInternalServerError, "userUpdate - h.user.Update")
		return
	}

	_ = response

	//h.Resp(w, mapper.ConvUserResponse(response), http.StatusOK)
}

func (h *Handler) userLogin(w http.ResponseWriter, r *http.Request) {
	req := entity.LoginIn{}
	if err := req.ParseHttpBody(r); err != nil {
		h.Error(w, err, http.StatusBadRequest, "userLogin - Parse")
		return
	}

	response, err := h.user.Login(r.Context(), req)
	if err != nil {
		h.Error(w, err, http.StatusInternalServerError, "h.user.Login")
		return
	}

	h.Resp(w, response, http.StatusOK)
}

func (h *Handler) userLogout(w http.ResponseWriter, r *http.Request) {
	req := entity.LogoutIn{}
	if err := req.ParseHttpBody(r); err != nil {
		h.Error(w, err, http.StatusBadRequest, "userRegister - Parse")
		return
	}

	h.Resp(w, "success", http.StatusOK)
}

func (h *Handler) userGet(w http.ResponseWriter, r *http.Request) {
	req := entity.GetIn{}

	payload := h.payload(r)
	req.ID = payload.UserID

	user, err := h.user.Get(r.Context(), req)
	if err != nil {
		h.Error(w, err, http.StatusInternalServerError, "userGet - h.user.Get")
		return
	}

	h.Resp(w, mapper.ConvUserResponse(user), http.StatusOK)
}

func (h *Handler) usersFind(w http.ResponseWriter, r *http.Request) {
	req := entity.FindRequest{}
	if err := req.ParseHttpBody(r); err != nil {
		h.Error(w, err, http.StatusBadRequest, "usersFind - ParseAndValidate")
		return
	}

	h.Resp(w, "success", http.StatusOK)
}
