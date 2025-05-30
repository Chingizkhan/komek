package token

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const (
	TypeBearer       = "bearer"
	authorizationKey = "Authorization"

	AccessToken  = "Access-Token"
	RefreshToken = "Refresh-Token"
)

var (
	ErrInvalidAuthHeader   = errors.New("invalid authorization header")
	ErrUnsupportedAuthType = errors.New("unsupported authorization type: must be 'Bearer'")
)

func FromHttpRequest(r *http.Request) string {
	authKey := r.Header.Get(authorizationKey)
	return strings.TrimSpace(strings.Replace(authKey, TypeBearer, "", 1))
}

func Get(r *http.Request) (string, error) {
	authKey := r.Header.Get(authorizationKey)
	fields := strings.Fields(authKey)
	if len(fields) < 2 {
		return "", ErrInvalidAuthHeader
	}
	authType := strings.ToLower(fields[0])
	if authType != TypeBearer {
		return "", ErrUnsupportedAuthType
	}
	return fields[1], nil
}

func GetFromCookie(r *http.Request) (string, error) {
	token, err := r.Cookie(AccessToken)
	if err != nil {
		return "", fmt.Errorf("access_token not found in cookie: %w", err)
	}
	return token.Value, nil
}
