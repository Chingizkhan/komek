package dto

import (
	"github.com/google/uuid"
	"komek/internal/domain"
	"net/http"
)

type (
	UserUpdateRequest struct {
		ID            uuid.UUID
		Name          string
		Login         string
		Email         domain.Email
		EmailVerified bool
		Phone         domain.Phone
		PasswordHash  string
	}

	UserChangePasswordRequest struct {
		ID       uuid.UUID
		Password string
	}

	UserDeleteRequest struct {
		ID uuid.UUID
	}

	UserRegisterRequest struct {
		Login    string
		Phone    domain.Phone
		Password string
		Roles    domain.Roles
	}

	UserLoginRequest struct {
	}

	UserLogoutRequest struct {
	}

	UserGetRequest struct {
	}

	UserFindRequest struct {
	}
)

func (req *UserUpdateRequest) ParseAndValidate(r *http.Request) error {
	return nil
}

func (req *UserChangePasswordRequest) ParseAndValidate(r *http.Request) error {
	return nil
}

func (req *UserDeleteRequest) ParseAndValidate(r *http.Request) error {
	return nil
}

func (req *UserRegisterRequest) ParseAndValidate(r *http.Request) error {
	return nil
}

func (req *UserLoginRequest) ParseAndValidate(r *http.Request) error {
	return nil
}

func (req *UserLogoutRequest) ParseAndValidate(r *http.Request) error {
	return nil
}

func (req *UserGetRequest) ParseAndValidate(r *http.Request) error {
	return nil
}

func (req *UserFindRequest) ParseAndValidate(r *http.Request) error {
	return nil
}
