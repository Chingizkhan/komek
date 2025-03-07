package repository

import (
	"errors"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"komek/internal/errs"
)

const (
	constraintFromAccountID = "fk_from_account_id"
	constraintToAccountID   = "fk_to_account_id"
)

func checkConstraints(err error) error {
	var e *pgconn.PgError
	if errors.As(err, &e) && e.Code == pgerrcode.ForeignKeyViolation {
		switch e.ConstraintName {
		case constraintFromAccountID:
			return errs.FromAccountNotFound
		case constraintToAccountID:
			return errs.ToAccountNotFound
		}
	}
	return nil
}
