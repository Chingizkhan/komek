package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	customMiddleware "komek/internal/controller/http/middleware"
	"net/http"
)

func (h *Handler) clientRoutes(r *chi.Mux) {
	r.Route("/client", func(r chi.Router) {
		// protected
		r.Route("/", func(r chi.Router) {
			r.Use(customMiddleware.Auth(h.tokenMaker))

			//r.Delete("/delete", h.userDelete)
			//r.Put("/change-password", h.userChangePassword)
			//r.Put("/update", h.userUpdate)
			//r.Post("/logout", h.userLogout)
			//r.Get("/", h.userGet)
			//r.Get("/find", h.usersFind)
		})
		// public
		r.Get("/list", h.list)
		r.Get("/{clientID}", h.getByID)
		//r.Post("/register", h.userRegister)
		//r.Post("/login", h.userLogin)
		//r.Post("/refresh-token", h.userRefreshToken)
	})
}

func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
	clients, err := h.client.List(r.Context())
	if err != nil {
		h.Error(w, err, http.StatusInternalServerError)
		return
	}

	h.Resp(w, clients, http.StatusOK)
}

func (h *Handler) getByID(w http.ResponseWriter, r *http.Request) {
	clientIDParam := chi.URLParam(r, "clientID")

	clientID, err := uuid.Parse(clientIDParam)
	if err != nil {
		h.Error(w, err, http.StatusBadRequest)
		return
	}

	clients, err := h.client.GetByID(r.Context(), clientID)
	if err != nil {
		h.Error(w, err, http.StatusInternalServerError)
		return
	}

	h.Resp(w, clients, http.StatusOK)
}
