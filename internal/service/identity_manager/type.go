package identity_manager

import (
	"github.com/Nerzal/gocloak/v13"
)

type (
	SignInRequest struct {
		Username string
		Password string
	}

	UserWithJWTResponse struct {
		User
		JWT
	}

	User struct {
		ID            string
		UserName      string
		Enabled       bool
		Email         string
		EmailVerified bool
		FirstName     string
		LastName      string
		Phone         string
		CreatedAt     int64
	}

	JWT struct {
		AccessToken      string
		RefreshToken     string
		ExpiresIn        int
		RefreshExpiresIn int
		SessionState     string
	}
)

func convertUser(u *gocloak.User) User {
	attr := *u.Attributes
	phone := attr["mobile"]

	return User{
		ID:            *u.ID,
		UserName:      *u.Username,
		Enabled:       *u.Enabled,
		Email:         *u.Email,
		EmailVerified: *u.EmailVerified,
		FirstName:     *u.FirstName,
		LastName:      *u.LastName,
		Phone:         phone[0],
		CreatedAt:     *u.CreatedTimestamp,
	}
}

func convertJWT(jwt *gocloak.JWT) JWT {
	return JWT{
		AccessToken:      jwt.AccessToken,
		RefreshToken:     jwt.RefreshToken,
		ExpiresIn:        jwt.ExpiresIn,
		RefreshExpiresIn: jwt.RefreshExpiresIn,
		SessionState:     jwt.SessionState,
	}
}
