package user

import (
	"context"
	"github.com/Nerzal/gocloak/v13"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"komek/internal/domain"
	"komek/internal/domain/user/entity"
)

type (
	Service interface {
		Register(ctx context.Context, u entity.RegisterIn) (entity.User, error)
		Update(ctx context.Context, req entity.UpdateIn) (entity.User, error)
		Delete(ctx context.Context, id uuid.UUID) error
		Get(ctx context.Context, req entity.GetIn) (entity.User, error)
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
