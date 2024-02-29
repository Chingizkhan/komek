package dto

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"komek/internal/domain"
	"net/http"
	"time"
)

type (
	UserUpdateRequest struct {
		ID            uuid.UUID    `json:"id"`
		Name          string       `json:"name"`
		Login         string       `json:"login"`
		Email         domain.Email `json:"email"`
		EmailVerified *bool        `json:"email_verified"`
		Phone         domain.Phone `json:"phone"`
		PasswordHash  string       `json:"password_hash"`
		Roles         domain.Roles `json:"roles"`
	}

	UserChangePasswordRequest struct {
		ID          uuid.UUID       `json:"id"`
		OldPassword domain.Password `json:"old_password"`
		NewPassword domain.Password `json:"new_password"`
	}

	UserDeleteRequest struct {
		ID uuid.UUID `json:"id"`
	}

	UserRegisterRequest struct {
		Login    string          `json:"login"`
		Phone    domain.Phone    `json:"phone"`
		Password domain.Password `json:"password"`
		Roles    domain.Roles    `json:"roles"`
	}

	UserResponse struct {
		ID            uuid.UUID    `json:"id"`
		Name          string       `json:"name"`
		Login         string       `json:"login"`
		Email         domain.Email `json:"email"`
		EmailVerified bool         `json:"email_verified"`
		Roles         domain.Roles `json:"roles"`
		CreatedAt     time.Time    `json:"created_at"`
		UpdatedAt     time.Time    `json:"updated_at"`
	}

	UserLoginRequest struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	UserLoginResponse struct {
		AccessToken string       `json:"access_token"`
		User        UserResponse `json:"user"`
	}

	UserLogoutRequest struct {
	}

	UserGetRequest struct {
		ID uuid.UUID
	}

	UserFindRequest struct {
		Name  string       `json:"name"`
		Login string       `json:"login"`
		Email domain.Email `json:"email"`
	}
)

func (req *UserUpdateRequest) ParseAndValidate(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return fmt.Errorf("can not decode body: %w", err)
	}
	if !req.Roles.Allowed() {
		return errors.New("role not allowed")
	}
	return nil
}

func (req *UserChangePasswordRequest) ParseAndValidate(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return fmt.Errorf("can not decode body: %w", err)
	}
	if req.OldPassword == req.NewPassword {
		return errors.New("passwords are same")
	}
	if err = req.NewPassword.Validate(); err != nil {
		return err
	}
	return nil
}

func (req *UserDeleteRequest) ParseAndValidate(r *http.Request) error {
	return nil
}

func (req *UserRegisterRequest) ParseAndValidate(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return fmt.Errorf("can not decode body: %w", err)
	}
	if len(req.Phone) != 11 {
		return errors.New("invalid phone length")
	}
	// todo: add validation on phone mask
	if !req.Roles.Allowed() {
		return errors.New("role not allowed")
	}
	if len(req.Login) < 6 {
		return errors.New("login too short: must be >= 6")
	}
	// todo: add validation for password (min:6, chars, digits)
	if err = req.Password.Validate(); err != nil {
		return err
	}
	return nil
}

func (req *UserLoginRequest) ParseAndValidate(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return fmt.Errorf("can not decode body: %w", err)
	}
	return nil
}

func (req *UserLogoutRequest) ParseAndValidate(r *http.Request) error {
	return nil
}

func (req *UserGetRequest) ParseAndValidate(r *http.Request) error {
	userID := chi.URLParam(r, "id")
	id, err := uuid.Parse(userID)
	if err != nil {
		return errors.New("invalid_user_id")
	}
	req.ID = id
	return nil
}

func (req *UserFindRequest) ParseAndValidate(r *http.Request) error {
	return nil
}
