package user_uc

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"komek/internal/domain"
)

type UserRepository interface {
	Save(ctx context.Context, tx pgx.Tx, u domain.User) error
	Get(ctx context.Context, userID uuid.UUID) (domain.User, error)
	UpdateName(ctx context.Context, tx pgx.Tx, id uuid.UUID, value string) (uuid.UUID, error)
	UpdateLogin(ctx context.Context, tx pgx.Tx, id uuid.UUID, value string) (uuid.UUID, error)
	UpdateEmail(ctx context.Context, tx pgx.Tx, id uuid.UUID, value domain.Email) (uuid.UUID, error)
	UpdateEmailVerified(ctx context.Context, tx pgx.Tx, id uuid.UUID, value bool) (uuid.UUID, error)
	UpdatePhone(ctx context.Context, tx pgx.Tx, id uuid.UUID, value domain.Phone) (uuid.UUID, error)
	UpdatePasswordHash(ctx context.Context, tx pgx.Tx, id uuid.UUID, value string) (uuid.UUID, error)
}
