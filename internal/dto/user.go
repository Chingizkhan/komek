package dto

import (
	"github.com/google/uuid"
	"komek/internal/domain"
)

type UpdateUserRequest struct {
	ID            uuid.UUID
	Name          string
	Login         string
	Email         domain.Email
	EmailVerified bool
	Phone         domain.Phone
	PasswordHash  string
}
