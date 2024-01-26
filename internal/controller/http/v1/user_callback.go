package v1

import (
	"encoding/json"
	"fmt"
	"komek/internal/controller/http/api_util"
	"komek/pkg/logger"
	"net/http"
)

type CallbackRequest struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	MobileNumber string `json:"mobile_number"`
}

type CallbackResponse struct {
	Status string `json:"status"`
}

func (h *Handler) callback(w http.ResponseWriter, r *http.Request) {
	const fnName = "user_http - callback"

	var req CallbackRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.l.Error(fnName, logger.Err(err))
		api_util.RenderErrorResponse(w, "invalid request", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	h.l.Info("callback", "info", fmt.Sprintf("%#v", req))

	//_, err = h.userUC.Register(r.Context(), convertRegisterRequestToDomain(req))
	//if err != nil {
	//	h.l.Error(fnName, logger.Err(err))
	//	api_util.RenderErrorResponse(w, "register failed", http.StatusInternalServerError)
	//	return
	//}

	api_util.RenderResponse(w,
		&CallbackResponse{
			Status: "success",
		},
		http.StatusCreated,
	)
}
