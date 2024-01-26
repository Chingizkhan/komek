package jwt

import (
	"context"
	"crypto/rsa"
	"fmt"
	golangJwt "github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"komek/pkg/header"
	"net/http"
)

const (
	authorizationKey = "Authorization"
	tokenType        = "Bearer"
	CtxKeyClaims     = "ctx_key_claims"
)

type (
	JWT struct {
		Raw    string
		Token  *golangJwt.Token
		Claims Claims
	}

	// todo: maybe add isParsed flag to Claims and Methods to extract Claims like Id() which returns id if Claims were parsed, otherwise error

	Claims struct {
		Id       string
		Name     string
		UserName string
		Email    Email
		Roles    []string
		Exp      float64
	}

	Email struct {
		Verified bool
		Value    string
	}
)

func FromRequestHeader(r *http.Request) JWT {
	return JWT{
		Raw: header.Get(r, authorizationKey, tokenType),
	}
}

func New(raw string) JWT {
	return JWT{
		Raw:    raw,
		Token:  nil,
		Claims: Claims{},
	}
}

func (jwt *JWT) Parse(pk *rsa.PublicKey) error {
	tok, err := golangJwt.Parse(jwt.Raw, parse(pk))
	if err != nil {
		return errors.Wrap(err, "unable to parse token")
	}
	jwt.Token = tok
	return nil
}

func (jwt *JWT) GetUserIdClaim() {
	claims, ok := jwt.Token.Claims.(golangJwt.MapClaims)
	if ok && jwt.Token.Valid {
		jwt.Claims.Id = claims["sub"].(string)
	}
}

func (jwt *JWT) GetClaims() {
	// todo: check errors
	claims, ok := jwt.Token.Claims.(golangJwt.MapClaims)
	if ok && jwt.Token.Valid {
		jwt.Claims.Id = claims["sub"].(string)
		jwt.Claims.Name = claims["name"].(string)
		jwt.Claims.UserName = claims["preferred_username"].(string)
		jwt.Claims.Email.Value = claims["email"].(string)
		jwt.Claims.Email.Verified = claims["email_verified"].(bool)
		jwt.Claims.Exp = claims["exp"].(float64)

		realmAccess := claims["realm_access"].(map[string]interface{})
		roles := realmAccess["roles"].([]interface{})

		jwt.Claims.Roles = make([]string, 0, 5)
		for _, r := range roles {
			jwt.Claims.Roles = append(jwt.Claims.Roles, r.(string))
		}
	}
}

func (jwt *JWT) SetClaims(ctx context.Context, r *http.Request) (req *http.Request) {
	jwt.GetClaims()

	req = r.WithContext(context.WithValue(ctx, CtxKeyClaims, jwt.Claims))
	return req
}

func GetClaims(ctx context.Context) Claims {
	claims := ctx.Value(CtxKeyClaims)
	cl := claims.(Claims)
	return cl
}

func parse(pk *rsa.PublicKey) golangJwt.Keyfunc {
	return func(token *golangJwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*golangJwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return pk, nil
	}
}
