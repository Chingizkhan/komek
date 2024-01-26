package v1

import (
	"komek/internal/controller/http/api_util"
	"komek/internal/domain/jwt"
	"komek/internal/usecase/user_managment"
	"komek/pkg/logger"
	"net/http"
)

type LogoutRequest struct {
	UserID string `json:"user_id"`
}

type LogoutResponse struct {
	Status string `json:"status"`
}

func (h *Handler) logout(w http.ResponseWriter, r *http.Request) {
	const fnName = "user_http - register"

	claims := jwt.GetClaims(r.Context())

	err := h.userUC.Logout(r.Context(), user_managment.LogoutRequest{UserID: claims.Id})
	if err != nil {
		h.l.Error(fnName, logger.Err(err))
		api_util.RenderErrorResponse(w, "logout failed", http.StatusInternalServerError)
		return
	}

	api_util.RenderResponse(w,
		&LogoutResponse{
			Status: "success",
		},
		http.StatusOK,
	)
}
