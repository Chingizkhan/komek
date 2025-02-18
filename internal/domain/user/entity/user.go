package entity

import (
	"github.com/google/uuid"
	"komek/internal/domain/email"
	"komek/internal/domain/phone"
	"komek/internal/domain/role"
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

	GetRequest struct {
		ID        uuid.UUID
		Name      string
		Login     string
		Phone     phone.Phone
		Email     email.Email
		AccountID int64
	}
)
