package service

import (
	"context"
	"github.com/google/uuid"
	"komek/internal/domain/email"
	"komek/internal/domain/phone"
	"komek/internal/domain/user/entity"
)

type UserRepository interface {
	Save(ctx context.Context, u entity.User) (entity.User, error)
	GetByID(ctx context.Context, userID uuid.UUID) (entity.User, error)
	GetByPhone(ctx context.Context, phone phone.Phone) (entity.User, error)
	GetByEmail(ctx context.Context, email email.Email) (entity.User, error)
	GetByLogin(ctx context.Context, login string) (entity.User, error)
	GetByAccount(ctx context.Context, accountID uuid.UUID) (entity.User, error)
	Update(ctx context.Context, req entity.UpdateIn) (entity.User, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Find(ctx context.Context, req entity.FindRequest) ([]entity.User, error)
}
