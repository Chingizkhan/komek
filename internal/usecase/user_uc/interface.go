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
		GetByID(ctx context.Context, tx pgx.Tx, userID uuid.UUID) (domain.User, error)
		GetByPhone(ctx context.Context, tx pgx.Tx, phone domain.Phone) (domain.User, error)
		GetByEmail(ctx context.Context, tx pgx.Tx, email domain.Email) (domain.User, error)
		GetByLogin(ctx context.Context, tx pgx.Tx, login string) (domain.User, error)
		GetByAccount(ctx context.Context, tx pgx.Tx, accountID int64) (domain.User, error)
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

	SessionRepository interface {
		Get(ctx context.Context, tx pgx.Tx, id uuid.UUID) (domain.Session, error)
		Save(ctx context.Context, tx pgx.Tx, s domain.Session) (domain.Session, error)
	}
)
