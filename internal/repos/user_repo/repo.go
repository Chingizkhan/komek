package user_repo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	user_db "komek/db/sqlc"
	"komek/internal/domain"
	"komek/pkg/postgres"
)

const (
	_defaultEntityCap = 50
)

type Repository struct {
	pool *pgxpool.Pool
	q    *user_db.Queries
}

func New(pg *postgres.Postgres) *Repository {
	return &Repository{pg.Pool, user_db.New(pg.Pool)}
}

func (r *Repository) Get(ctx context.Context, userID uuid.UUID) (domain.User, error) {
	u, err := r.q.GetUser(ctx, userID)
	if err != nil {
		return domain.User{}, fmt.Errorf("r.q.GetUser :%w", err)
	}
	return domain.User{
		ID:            u.ID,
		Name:          u.Name.String,
		Phone:         domain.Phone(u.Phone.String),
		Login:         u.Login,
		EmailVerified: u.EmailVerified.Bool,
		PasswordHash:  u.PasswordHash,
		Email:         domain.Email(u.Email.String),
		CreatedAt:     u.CreatedAt.Time,
		UpdatedAt:     u.UpdatedAt.Time,
	}, nil
}

func (r *Repository) Save(ctx context.Context, tx pgx.Tx, u domain.User) error {
	qtx := r.q.WithTx(tx)

	err := qtx.SaveUser(ctx, user_db.SaveUserParams{
		Login:        u.Login,
		PasswordHash: u.PasswordHash,
	})
	if err != nil {
		return fmt.Errorf("qtx.SaveUser: %w", err)
	}

	if err != nil {
		return fmt.Errorf("transactional.Commit: %w", err)
	}
	return nil
}

func (r *Repository) UpdateName(ctx context.Context, tx pgx.Tx, id uuid.UUID, value string) (uuid.UUID, error) {
	qtx := r.q.WithTx(tx)
	name := sql.NullString{
		String: value,
		Valid:  true,
	}
	id, err := qtx.UpdateUserName(ctx, user_db.UpdateUserNameParams{
		ID:   id,
		Name: name,
	})
	if err != nil {
		return uuid.Nil, fmt.Errorf("r.q.UpdateName: %w", err)
	}
	return id, nil
}

func (r *Repository) UpdateLogin(ctx context.Context, tx pgx.Tx, id uuid.UUID, value string) (uuid.UUID, error) {
	qtx := r.q.WithTx(tx)
	id, err := qtx.UpdateUserLogin(ctx, user_db.UpdateUserLoginParams{
		ID:    id,
		Login: value,
	})
	if err != nil {
		return uuid.Nil, fmt.Errorf("r.q.UpdateLogin: %w", err)
	}
	return id, nil
}

func (r *Repository) UpdateEmail(ctx context.Context, tx pgx.Tx, id uuid.UUID, value domain.Email) (uuid.UUID, error) {
	qtx := r.q.WithTx(tx)
	email := sql.NullString{
		String: string(value),
		Valid:  true,
	}
	id, err := qtx.UpdateUserEmail(ctx, user_db.UpdateUserEmailParams{
		ID:    id,
		Email: email,
	})
	if err != nil {
		return uuid.Nil, fmt.Errorf("r.q.UpdateEmail: %w", err)
	}
	return id, nil
}

func (r *Repository) UpdateEmailVerified(ctx context.Context, tx pgx.Tx, id uuid.UUID, value bool) (uuid.UUID, error) {
	qtx := r.q.WithTx(tx)
	verified := sql.NullBool{
		Bool:  value,
		Valid: true,
	}
	id, err := qtx.UpdateUserEmailVerified(ctx, user_db.UpdateUserEmailVerifiedParams{
		ID:            id,
		EmailVerified: verified,
	})
	if err != nil {
		return uuid.Nil, fmt.Errorf("r.q.UpdateEmailVerified: %w", err)
	}
	return id, nil
}

func (r *Repository) UpdatePhone(ctx context.Context, tx pgx.Tx, id uuid.UUID, value domain.Phone) (uuid.UUID, error) {
	qtx := r.q.WithTx(tx)
	phone := sql.NullString{
		String: string(value),
		Valid:  true,
	}
	id, err := qtx.UpdateUserPhone(ctx, user_db.UpdateUserPhoneParams{
		ID:    id,
		Phone: phone,
	})
	if err != nil {
		return uuid.Nil, fmt.Errorf("r.q.UpdatePhone: %w", err)
	}
	return id, nil
}

func (r *Repository) UpdatePasswordHash(ctx context.Context, tx pgx.Tx, id uuid.UUID, value string) (uuid.UUID, error) {
	qtx := r.q.WithTx(tx)

	id, err := qtx.UpdateUserPasswordHash(ctx, user_db.UpdateUserPasswordHashParams{
		ID:           id,
		PasswordHash: value,
	})
	if err != nil {
		return uuid.Nil, fmt.Errorf("r.q.UpdatePasswordHash: %w", err)
	}
	return id, nil
}

func (r *Repository) Delete(ctx context.Context, tx pgx.Tx, id uuid.UUID) error {
	qtx := r.q.WithTx(tx)

	_, err := qtx.RemoveUser(ctx, id)
	if err != nil {
		return fmt.Errorf("qtx.RemoveUser: %w", err)
	}
	return nil
}
