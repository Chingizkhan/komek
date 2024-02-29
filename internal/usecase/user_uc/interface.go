package user_uc

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"komek/internal/domain"
	"komek/internal/dto"
)

type (
	UserRepository interface {
		Save(ctx context.Context, tx pgx.Tx, u domain.User) (domain.User, error)
		Get(ctx context.Context, tx pgx.Tx, userID uuid.UUID) (domain.User, error)
		GetUserByLogin(ctx context.Context, tx pgx.Tx, login string) (domain.User, error)
		Update(ctx context.Context, tx pgx.Tx, req dto.UserUpdateRequest) (domain.User, error)
		Delete(ctx context.Context, tx pgx.Tx, id uuid.UUID) error
		Find(ctx context.Context, tx pgx.Tx, req dto.UserFindRequest) ([]domain.User, error)
	}

	Transactional interface {
		Exec(ctx context.Context, fn func(tx pgx.Tx) error) error
	}

	Hasher interface {
		Hash(value string) (string, error)
		CheckHash(password, hash string) bool
	}
)
