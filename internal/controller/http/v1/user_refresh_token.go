package v1

import (
	"fmt"
	golangJwt "github.com/golang-jwt/jwt/v5"
	"golang.org/x/net/context"
	"komek/internal/controller/http/api_util"
	"komek/internal/domain/jwt"
	"komek/internal/service/identity_manager"
	"komek/pkg/header"
	"komek/pkg/logger"
	"komek/pkg/rsa"
	"net/http"
)

type RefreshTokenResponse struct {
	JWT  `json:"jwt"`
	User `json:"user"`
}

func (h *Handler) refreshTokens(w http.ResponseWriter, r *http.Request) {
	const fnName = "user_http - refreshTokens"

	refreshTokenRaw := header.Get(r, "Refresh-Token", "Bearer")

	token := jwt.New(refreshTokenRaw)
	pk, err := rsa.ParseKeycloakRSAPublicKey(h.cfg.RealmRS256PublicKey)
	if err != nil {
		h.l.Error(fnName, logger.Err(err), 1)
		api_util.RenderErrorResponse(w, "refresh tokens failed", http.StatusInternalServerError)
		return
	}
	golangJwt.Parse(refreshTokenRaw, func(token *golangJwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*golangJwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return pk, nil
	})
	err = token.Parse(pk)
	if err != nil {
		h.l.Error(fnName, logger.Err(err), 2)
		api_util.RenderErrorResponse(w, "refresh tokens failed", http.StatusInternalServerError)
		return
	}

	token.GetUserIdClaim()

	r = r.WithContext(context.WithValue(r.Context(), jwt.CtxKeyClaims, token.Claims))

	resp, err := h.tokenUC.RefreshTokens(r.Context(), refreshTokenRaw)
	if err != nil {
		h.l.Error(fnName, logger.Err(err), 3)
		api_util.RenderErrorResponse(w, "refresh tokens failed", http.StatusInternalServerError)
		return
	}

	api_util.RenderResponse(w,
		convertRefreshTokensResponse(resp),
		http.StatusOK,
	)
}

func convertRefreshTokensResponse(r *identity_manager.UserWithJWTResponse) *SignInResponse {
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
