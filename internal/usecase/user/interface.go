package user

import (
	"context"
	"github.com/Nerzal/gocloak/v13"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"komek/internal/domain"
	"komek/internal/domain/email"
	"komek/internal/domain/phone"
	"komek/internal/domain/user/entity"
)

type (
	Service interface {
		Register(ctx context.Context, u entity.RegisterIn) (entity.User, error)
		Update(ctx context.Context, req entity.UpdateIn) (entity.User, error)
		Delete(ctx context.Context, id uuid.UUID) error
		Get(ctx context.Context, req entity.GetIn) (entity.User, error)
	}

	UserRepository interface {
		Save(ctx context.Context, u entity.User) (entity.User, error)
		GetByID(ctx context.Context, userID uuid.UUID) (entity.User, error)
		GetByPhone(ctx context.Context, phone phone.Phone) (entity.User, error)
		GetByEmail(ctx context.Context, email email.Email) (entity.User, error)
		GetByLogin(ctx context.Context, login string) (entity.User, error)
		GetByAccount(ctx context.Context, accountID int64) (entity.User, error)
		Update(ctx context.Context, req entity.UpdateIn) (entity.User, error)
		Delete(ctx context.Context, id uuid.UUID) error
		Find(ctx context.Context, req entity.FindRequest) ([]entity.User, error)
	}

	IdentityManager interface {
		CreateUser(ctx context.Context, user gocloak.User, password, role string) (*gocloak.User, error)
	}

	Hasher interface {
		Hash(value string) (string, error)
		CheckHash(password, hash string) bool
	}

	SessionRepository interface {
		Get(ctx context.Context, tx pgx.Tx, id uuid.UUID) (domain.Session, error)
		Save(ctx context.Context, tx pgx.Tx, s domain.Session) (domain.Session, error)
	}
)
