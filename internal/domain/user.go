package domain

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID            uuid.UUID
	Name          string
	Phone         Phone
	Login         string
	Email         Email
	EmailVerified bool
	PasswordHash  string
	Roles         []Role
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
