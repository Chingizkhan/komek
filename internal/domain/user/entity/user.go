package entity

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	account "komek/internal/domain/account/entity"
	country "komek/internal/domain/country/entity"
	currency "komek/internal/domain/currency/entity"
	"komek/internal/domain/email"
	"komek/internal/domain/password"
	"komek/internal/domain/phone"
	"komek/internal/domain/role"
	"komek/internal/errs"
	"net/http"
	"time"
)

type (
	User struct {
		ID                uuid.UUID
		Name              string
		Phone             phone.Phone
		Login             string
		Email             email.Email
		EmailVerified     bool
		PasswordHash      string
		Roles             role.Roles
		CreatedAt         time.Time
		UpdatedAt         time.Time
		PasswordChangedAt time.Time
	}

	AccountResponse struct {
		ID       uuid.UUID             `json:"id"`
		Balance  float64               `json:"balance"`
		Currency currency.Currency     `json:"currency"`
		Country  country.Country       `json:"country"`
		Status   account.AccountStatus `json:"status"`
	}

	UserResponse struct {
		ID                uuid.UUID       `json:"id"`
		Name              string          `json:"name"`
		Phone             phone.Phone     `json:"phone"`
		Login             string          `json:"login"`
		Email             email.Email     `json:"email"`
		EmailVerified     bool            `json:"email_verified"`
		Roles             role.Roles      `json:"roles"`
		CreatedAt         int64           `json:"created_at"`
		UpdatedAt         int64           `json:"updated_at"`
		PasswordChangedAt int64           `json:"password_changed_at"`
		Account           AccountResponse `json:"account"`
	}

	RegisterIn struct {
		Login    string            `json:"login"`
		Phone    phone.Phone       `json:"phone"`
		Password password.Password `json:"password"`
		Roles    role.Roles        `json:"roles"`

		PasswordHash string
	}

	LoginIn struct {
		Login    string      `json:"login"`
		Phone    phone.Phone `json:"phone"`
		Email    email.Email `json:"email"`
		Password string      `json:"password"`
	}

	LoginOut struct {
		SessionID             uuid.UUID    `json:"session_id"`
		AccessToken           string       `json:"access_token"`
		AccessTokenExpiresAt  time.Time    `json:"access_token_expires_at"`
		RefreshToken          string       `json:"refresh_token"`
		RefreshTokenExpiresAt time.Time    `json:"refresh_token_expires_at"`
		User                  UserResponse `json:"user"`
	}

	LogoutIn struct {
	}

	UpdateIn struct {
		ID            uuid.UUID   `json:"id"`
		Name          string      `json:"name"`
		Login         string      `json:"login"`
		Email         email.Email `json:"email"`
		EmailVerified *bool       `json:"email_verified"`
		Phone         phone.Phone `json:"phone"`
		PasswordHash  string      `json:"password_hash"`
		Roles         role.Roles  `json:"roles"`
	}

	GetIn struct {
		ID        uuid.UUID
		Name      string
		Login     string
		Phone     phone.Phone
		Email     email.Email
		AccountID uuid.UUID
	}

	GetOut struct {
		User    User
		Account account.Account
	}

	DeleteIn struct {
		ID uuid.UUID `json:"id"`
	}

	RefreshTokensIn struct {
		RefreshToken string `json:"refresh_token"`
	}

	RefreshTokensOut struct {
		AccessToken           string    `json:"access_token"`
		AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
		RefreshToken          string    `json:"refresh_token"`
		RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
	}

	ChangePasswordIn struct {
		ID          uuid.UUID         `json:"id"`
		OldPassword password.Password `json:"old_password"`
		NewPassword password.Password `json:"new_password"`
	}

	FindRequest struct {
		Name  string      `json:"name"`
		Login string      `json:"login"`
		Email email.Email `json:"email"`
	}
)

func (req *RegisterIn) ParseHttpBody(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return fmt.Errorf("can not decode body: %w", err)
	}
	defer r.Body.Close()
	return req.Validate()
}

func (req *RegisterIn) Validate() error {
	if len(req.Phone) != 11 {
		return errors.New("invalid phone length")
	}
	// todo: add validation on phone mask
	if len(req.Roles) != 0 {
		if !req.Roles.Allowed() {
			return errors.New("role not allowed")
		}
	} else {
		req.Roles = role.Roles{role.User}
	}

	//if len(req.Login) < 6 {
	//	return errors.New("login too short: must be >= 6")
	//}
	// todo: add validation for password (min:6, chars, digits)
	if err := req.Password.Validate(); err != nil {
		return err
	}
	return nil
}

func (req *RegisterIn) ToEntity() User {
	return User{
		Phone:        req.Phone,
		Login:        req.Login,
		Roles:        req.Roles,
		PasswordHash: req.PasswordHash,
	}
}

func (req *LoginIn) ParseHttpBody(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return fmt.Errorf("can not decode body: %w", err)
	}
	return nil
}

func (req *LogoutIn) ParseHttpBody(r *http.Request) error {
	return nil
}

func (req *GetIn) ParseHttpBody(r *http.Request) error {
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
		accID, err := uuid.Parse(accountID)
		if err != nil {
			return fmt.Errorf("%s: %w", err, errors.New("invalid_account_id"))
		}
		req.AccountID = accID
	}

	return nil
}

func (req *GetIn) Validate() error {
	if req.ID == uuid.Nil && req.Login == "" && req.Email == "" && req.Phone == "" && req.AccountID == uuid.Nil {
		return errs.UserGetParamsNotSpecified
	}
	return nil
}

func (req *UpdateIn) ParseHttpBody(r *http.Request) error {
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

func (user *User) ToResponse() UserResponse {
	return UserResponse{
		ID:                user.ID,
		Name:              user.Name,
		Login:             user.Login,
		Phone:             user.Phone,
		Email:             user.Email,
		EmailVerified:     user.EmailVerified,
		Roles:             user.Roles,
		CreatedAt:         user.CreatedAt.Unix(),
		UpdatedAt:         user.UpdatedAt.Unix(),
		PasswordChangedAt: user.PasswordChangedAt.Unix(),
	}
}

func (req *DeleteIn) ParseHttpBody(r *http.Request) error {
	return nil
}

func (req *RefreshTokensIn) ParseHttpBody(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return fmt.Errorf("can not decode body: %w", err)
	}
	return nil
}

func (req *ChangePasswordIn) ParseHttpBody(r *http.Request) error {
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

func (req *FindRequest) ParseHttpBody(r *http.Request) error {
	return nil
}
