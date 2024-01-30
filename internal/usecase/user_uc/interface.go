package user_uc

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"komek/internal/domain"
	"komek/internal/dto"
)

type (
	UserRepository interface {
		Save(ctx context.Context, tx pgx.Tx, u domain.User) error
		Get(ctx context.Context, userID uuid.UUID) (domain.User, error)
		GetWithTX(ctx context.Context, tx pgx.Tx, userID uuid.UUID) (domain.User, error)
		UpdateName(ctx context.Context, tx pgx.Tx, id uuid.UUID, value string) (uuid.UUID, error)
		UpdateLogin(ctx context.Context, tx pgx.Tx, id uuid.UUID, value string) (uuid.UUID, error)
		UpdateEmail(ctx context.Context, tx pgx.Tx, id uuid.UUID, value domain.Email) (uuid.UUID, error)
		UpdateEmailVerified(ctx context.Context, tx pgx.Tx, id uuid.UUID, value bool) (uuid.UUID, error)
		UpdatePhone(ctx context.Context, tx pgx.Tx, id uuid.UUID, value domain.Phone) (uuid.UUID, error)
		UpdatePasswordHash(ctx context.Context, tx pgx.Tx, id uuid.UUID, value string) (uuid.UUID, error)
		Delete(ctx context.Context, tx pgx.Tx, id uuid.UUID) error
		Find(ctx context.Context, req dto.UserFindRequest) ([]domain.User, error)
		FindWithTX(ctx context.Context, tx pgx.Tx, req dto.UserFindRequest) ([]domain.User, error)
	}

	Transactional interface {
		Start(ctx context.Context) (pgx.Tx, error)
	}

	Hasher interface {
		Hash(value string) (string, error)
		CheckHash(password, hash string) bool
	}
)
