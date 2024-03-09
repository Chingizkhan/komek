package account_repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"komek/db/sqlc"
	"komek/internal/domain"
	"komek/internal/dto"
	"komek/internal/mapper"
	"komek/pkg/postgres"
)

type Repository struct {
	pool *pgxpool.Pool
	q    *sqlc.Queries
}

func New(pg *postgres.Postgres) *Repository {
	return &Repository{pg.Pool, sqlc.New(pg.Pool)}
}

func (r *Repository) Get(ctx context.Context, tx pgx.Tx, id int64) (domain.Account, error) {
	qtx := r.queries(tx)
	acc, err := qtx.GetAccount(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Account{}, ErrAccountNotFound
		}
		return domain.Account{}, fmt.Errorf("r.q.GetAccount: %w", err)
	}
	return mapper.ConvAccountToDomain(acc), nil
}

func (r *Repository) Create(ctx context.Context, tx pgx.Tx, in dto.CreateAccountIn) (domain.Account, error) {
	qtx := r.queries(tx)

	account, err := qtx.CreateAccount(ctx, sqlc.CreateAccountParams{
		Owner: pgtype.UUID{
			Bytes: in.Owner,
			Valid: true,
		},
		Balance:  in.Balance,
		Currency: in.Currency,
	})
	if err != nil {
		if err = r.checkConstraints(err); err != nil {
			return domain.Account{}, err
		}
		return domain.Account{}, fmt.Errorf(": %w", err)
	}
	return mapper.ConvAccountToDomain(account), nil
}

func (r *Repository) queries(tx pgx.Tx) *sqlc.Queries {
	qtx := r.q
	if tx != nil {
		qtx = r.q.WithTx(tx)
	}
	return qtx
}

func (r *Repository) checkConstraints(err error) error {
	var e *pgconn.PgError
	if errors.As(err, &e) && e.Code == pgerrcode.UniqueViolation {
		switch e.ConstraintName {
		case ConstraintOwnerCurrencyKey:
			return ErrCurrencyAlreadyExists
		default:
			return ErrAccountAlreadyExists
		}
	}
	if e.Code == pgerrcode.ForeignKeyViolation && e.ConstraintName == ConstraintAccountsOwnerForeignKey {
		return ErrOwnerNotFound
	}
	return nil
}
