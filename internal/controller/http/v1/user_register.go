package v1

import (
	"encoding/json"
	"komek/internal/controller/http/api_util"
	"komek/pkg/logger"
	"net/http"
)

type RegisterRequest struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	MobileNumber string `json:"mobile_number"`
}

type RegisterResponse struct {
	Status string `json:"status"`
}

func (h *Handler) register(w http.ResponseWriter, r *http.Request) {
	const fnName = "user_http - register"

	var req RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.l.Error(fnName, logger.Err(err))
		api_util.RenderErrorResponse(w, "invalid request", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	_, err = h.userUC.Register(r.Context(), convertRegisterRequestToDomain(req))
	if err != nil {
		h.l.Error(fnName, logger.Err(err))
		api_util.RenderErrorResponse(w, "register failed", http.StatusInternalServerError)
		return
	}

	api_util.RenderResponse(w,
		&RegisterResponse{
			Status: "success",
		},
		http.StatusCreated,
	)
}
