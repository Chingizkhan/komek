package repository

import (
	"errors"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"komek/internal/errs"
)

const (
	constraintAccountID     = "fk_account_id"
	constraintTransactionID = "fk_transaction_id"
)

func (r *Repository) checkConstraints(err error) error {
	var e *pgconn.PgError
	if errors.As(err, &e) {
		switch e.Code {
		case pgerrcode.ForeignKeyViolation:
			switch e.ConstraintName {
			case constraintAccountID:
				return errs.AccountNotFound
			case constraintTransactionID:
				return errs.TransactionNotFound
			}
			//todo: finish unique, return operation already exists
		case pgerrcode.UniqueViolation:
			switch e.ConstraintName {
			case constraintTransactionID:
			}
		}
	}
	return nil
}
