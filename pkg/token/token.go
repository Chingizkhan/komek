package token

import (
	"net/http"
	"strings"
)

const (
	TypeBearer       = "Bearer"
	authorizationKey = "Authorization"

	AccessToken  = "Access-Token"
	RefreshToken = "Refresh-Token"
)

func FromHttpRequest(r *http.Request) string {
	return strings.TrimSpace(strings.Replace(r.Header.Get(authorizationKey), TypeBearer, "", 1))
}
