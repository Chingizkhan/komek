package dto

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"komek/internal/domain/email"
	"komek/internal/domain/password"
	"komek/internal/domain/phone"
	"komek/internal/domain/role"
	"net/http"
	"strconv"
	"time"
)

type (
	UserUpdateRequest struct {
		ID            uuid.UUID   `json:"id"`
		Name          string      `json:"name"`
		Login         string      `json:"login"`
		Email         email.Email `json:"email"`
		EmailVerified *bool       `json:"email_verified"`
		Phone         phone.Phone `json:"phone"`
		PasswordHash  string      `json:"password_hash"`
		Roles         role.Roles  `json:"roles"`
	}

	UserChangePasswordRequest struct {
		ID          uuid.UUID         `json:"id"`
		OldPassword password.Password `json:"old_password"`
		NewPassword password.Password `json:"new_password"`
	}

	UserDeleteRequest struct {
		ID uuid.UUID `json:"id"`
	}

	UserRegisterRequest struct {
		Login    string            `json:"login"`
		Phone    phone.Phone       `json:"phone"`
		Password password.Password `json:"password"`
		Roles    role.Roles        `json:"roles"`
	}

	UserResponse struct {
		ID                uuid.UUID   `json:"id"`
		Name              string      `json:"name"`
		Phone             phone.Phone `json:"phone"`
		Login             string      `json:"login"`
		Email             email.Email `json:"email"`
		EmailVerified     bool        `json:"email_verified"`
		Roles             role.Roles  `json:"roles"`
		CreatedAt         int64       `json:"created_at"`
		UpdatedAt         int64       `json:"updated_at"`
		PasswordChangedAt int64       `json:"password_changed_at"`
	}

	UserLoginRequest struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	UserLoginResponse struct {
		SessionID             uuid.UUID    `json:"session_id"`
		AccessToken           string       `json:"access_token"`
		AccessTokenExpiresAt  time.Time    `json:"access_token_expires_at"`
		RefreshToken          string       `json:"refresh_token"`
		RefreshTokenExpiresAt time.Time    `json:"refresh_token_expires_at"`
		User                  UserResponse `json:"user"`
	}

	UserRefreshTokensIn struct {
		RefreshToken string `json:"refresh_token"`
	}

	UserRefreshTokensOut struct {
		AccessToken          string    `json:"access_token"`
		AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
	}

	UserLogoutRequest struct {
	}

	UserGetRequest struct {
		ID        uuid.UUID
		Name      string
		Phone     phone.Phone
		Email     email.Email
		Login     string
		AccountID int64
	}

	UserFindRequest struct {
		Name  string      `json:"name"`
		Login string      `json:"login"`
		Email email.Email `json:"email"`
	}
)

func (req *UserUpdateRequest) ParseAndValidate(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return fmt.Errorf("can not decode body: %w", err)
	}
	defer r.Body.Close()
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
	defer r.Body.Close()
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
	defer r.Body.Close()
	return req.Validate()
}

func (req *UserRegisterRequest) Validate() error {
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
	if err := req.Password.Validate(); err != nil {
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
	userID := r.URL.Query().Get("user_id")
	if userID != "" {
		id, err := uuid.Parse(userID)
		if err != nil {
			return fmt.Errorf("%s: %w", err, errors.New("invalid_user_id"))
		}
		req.ID = id
	}

	phoneArg := r.URL.Query().Get("phone")
	if phoneArg != "" {
		req.Phone = phone.Phone(phoneArg)
	}

	emailArg := r.URL.Query().Get("email")
	if emailArg != "" {
		req.Email = email.Email(emailArg)
	}

	req.Login = r.URL.Query().Get("login")

	accountID := r.URL.Query().Get("account_id")
	if accountID != "" {
		accID, err := strconv.Atoi(accountID)
		if err != nil {
			return fmt.Errorf("%s: %w", err, errors.New("invalid_account_id"))
		}
		req.AccountID = int64(accID)
	}

	if req.ID == uuid.Nil && req.Login == "" && req.Email == "" && req.Phone == "" && req.AccountID == 0 {
		return errors.New("params_not_specified")
	}

	return nil
}

func (req *UserFindRequest) ParseAndValidate(r *http.Request) error {
	return nil
}

func (req *UserRefreshTokensIn) ParseAndValidate(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return fmt.Errorf("can not decode body: %w", err)
	}
	return nil
}
