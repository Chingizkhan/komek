package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"komek/db/sqlc"
	"komek/internal/domain/operation/entity"
	"komek/internal/service/transactional"
	"komek/pkg/null_value"
	"komek/pkg/postgres"
)

type Repository struct {
	pool *pgxpool.Pool
	q    *sqlc.Queries
}

func New(pg *postgres.Postgres) *Repository {
	return &Repository{pg.Pool, sqlc.New(pg.Pool)}
}

func (r *Repository) Create(ctx context.Context, in entity.CreateIn) (entity.Operation, error) {
	qtx := r.queries(ctx)

	operation, err := qtx.CreateOperation(ctx, sqlc.CreateOperationParams{
		TransactionID: null_value.UUID(in.TransactionID),
		AccountID:     null_value.UUID(in.AccountID),
		Type:          sqlc.OperationType(in.Type),
		Amount:        in.Amount,
		BalanceBefore: in.BalanceBefore,
		BalanceAfter:  in.BalanceAfter,
	})
	if err != nil {
		if err = r.checkConstraints(err); err != nil {
			return entity.Operation{}, fmt.Errorf("constraint: %w", err)
		}
		return entity.Operation{}, fmt.Errorf("create operation: %w", err)
	}

	return r.operationToDomain(operation), nil
}

func (r *Repository) GetByTransactionID(ctx context.Context, transactionID uuid.UUID) ([]entity.Operation, error) {
	qtx := r.queries(ctx)

	operations, err := qtx.GetOperationsByTransactionID(ctx, null_value.UUID(transactionID))
	if err != nil {
		return nil, fmt.Errorf("get operations by qtx: %w", err)
	}

	return r.operationsToDomain(operations), nil
}

func (r *Repository) queries(ctx context.Context) *sqlc.Queries {
	tx, ok := ctx.Value(transactional.TxKey).(pgx.Tx)
	if ok {
		return r.q.WithTx(tx)
	}
	return r.q
}
