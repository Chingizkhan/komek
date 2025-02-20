package entity

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

	RegisterIn struct {
		Login    string            `json:"login"`
		Phone    phone.Phone       `json:"phone"`
		Password password.Password `json:"password"`
		Roles    role.Roles        `json:"roles"`

		PasswordHash string
	}

	GetRequest struct {
		ID        uuid.UUID
		Name      string
		Login     string
		Phone     phone.Phone
		Email     email.Email
		AccountID int64
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
		Phone: req.Phone,
		Login: req.Login,
		Roles: req.Roles,
	}
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
