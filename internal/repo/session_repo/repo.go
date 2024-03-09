package session_repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"komek/db/sqlc"
	"komek/internal/domain"
	"komek/internal/mapper"
	"komek/internal/repo"
	"komek/pkg/postgres"
)

type Repository struct {
	pool *pgxpool.Pool
	q    *sqlc.Queries
}

func New(pg *postgres.Postgres) *Repository {
	return &Repository{pg.Pool, sqlc.New(pg.Pool)}
}

func (r *Repository) Get(ctx context.Context, tx pgx.Tx, id uuid.UUID) (domain.Session, error) {
	qtx := r.queries(tx)

	s, err := qtx.GetSession(ctx, repo.ConvertToUUID(id))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Session{}, ErrSessionNotFound
		}
		return domain.Session{}, fmt.Errorf("r.q.GetSession :%w", err)
	}
	return mapper.ConvSessionToDomain(s), nil
}

func (r *Repository) Save(ctx context.Context, tx pgx.Tx, s domain.Session) (domain.Session, error) {
	qtx := r.queries(tx)

	session, err := qtx.CreateSession(ctx, sqlc.CreateSessionParams{
		ID:           repo.ConvertToUUID(s.ID),
		UserID:       repo.ConvertToUUID(s.UserID),
		RefreshToken: s.RefreshToken,
		UserAgent:    s.UserAgent,
		ClientIp:     s.ClientIp,
		IsBlocked:    s.IsBlocked,
		ExpiresAt:    repo.ConvertToTimestamptz(s.ExpiresAt),
	})
	if err != nil {
		if err = checkConstraints(err); err != nil {
			return domain.Session{}, err
		}
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Session{}, ErrSessionNotFound
		}
		return domain.Session{}, fmt.Errorf("qtx.SaveSession: %w", err)
	}

	if err != nil {
		return domain.Session{}, fmt.Errorf("transactional.Commit: %w", err)
	}
	return mapper.ConvSessionToDomain(session), nil
}

func checkConstraints(err error) error {
	var e *pgconn.PgError
	if errors.As(err, &e) && e.Code == pgerrcode.UniqueViolation {
		switch e.ConstraintName {
		case ConstraintUserIDFKey:
			return ErrSessionUserIDNotFound
		}
	}
	return nil
}

func (r *Repository) queries(tx pgx.Tx) *sqlc.Queries {
	qtx := r.q
	if tx != nil {
		qtx = r.q.WithTx(tx)
	}
	return qtx
}
