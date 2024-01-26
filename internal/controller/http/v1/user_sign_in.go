package v1

import (
	"encoding/json"
	"komek/internal/controller/http/api_util"
	"komek/internal/usecase/user_managment"
	"komek/pkg/logger"
	"net/http"
)

type (
	SignInRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	SignInResponse struct {
		User `json:"user"`
		JWT  `json:"jwt"`
	}

	User struct {
		ID            string `json:"id"`
		UserName      string `json:"user_name"`
		Enabled       bool   `json:"enabled"`
		Email         string `json:"email"`
		EmailVerified bool   `json:"email_verified"`
		FirstName     string `json:"first_name"`
		LastName      string `json:"last_name"`
		Phone         string `json:"phone"`
		CreatedAt     int64  `json:"created_at"`
	}

	JWT struct {
		AccessToken      string `json:"access_token"`
		RefreshToken     string `json:"refresh_token"`
		ExpiresIn        int    `json:"expires_in"`
		RefreshExpiresIn int    `json:"refresh_expires_in"`
		SessionState     string `json:"session_state"`
	}
)

func (h *Handler) sigIn(w http.ResponseWriter, r *http.Request) {
	const fnName = "user_http - sign_in"

	var req SignInRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.l.Error(fnName, logger.Err(err))
		api_util.RenderErrorResponse(w, "invalid request", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	response, err := h.userUC.SignIn(r.Context(), convertSignInRequestToDomain(req))
	if err != nil {
		h.l.Error(fnName, logger.Err(err))
		api_util.RenderErrorResponse(w, "sign in failed", http.StatusInternalServerError)
		return
	}

	api_util.RenderResponse(w,
		convertSignInResponse(response),
		http.StatusOK,
	)
}

func convertSignInRequestToDomain(r SignInRequest) user_managment.SignInRequest {
	return user_managment.SignInRequest{
		Username: r.Username,
		Password: r.Password,
	}
}

func convertSignInResponse(r *user_managment.SignInResponse) *SignInResponse {
	return &SignInResponse{
		User: User{
			ID:            r.User.ID,
			UserName:      r.User.UserName,
			Enabled:       r.User.Enabled,
			Email:         r.User.Email,
			EmailVerified: r.User.EmailVerified,
			FirstName:     r.User.FirstName,
			LastName:      r.User.LastName,
			Phone:         r.User.Phone,
			CreatedAt:     r.User.CreatedAt,
		},
		JWT: JWT{
			AccessToken:      r.JWT.AccessToken,
			RefreshToken:     r.JWT.RefreshToken,
			ExpiresIn:        r.JWT.ExpiresIn,
			RefreshExpiresIn: r.JWT.RefreshExpiresIn,
			SessionState:     r.JWT.SessionState,
		},
	}
}
