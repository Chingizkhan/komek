package service

import (
	"context"
	"github.com/google/uuid"
	"komek/internal/domain/account/entity"
	operation "komek/internal/domain/operation/entity"
	transaction "komek/internal/domain/transaction/entity"
)

type (
	TransactionService interface {
		GetByAccountID(ctx context.Context, accountID uuid.UUID) ([]transaction.Transaction, error)
		Create(ctx context.Context, transaction transaction.Transaction) (transaction.Transaction, error)
	}

	OperationService interface {
		Create(ctx context.Context, in operation.CreateIn) (operation.Operation, error)
		GetByTransactionID(ctx context.Context, transactionID uuid.UUID) ([]operation.Operation, error)
	}

	AccountService interface {
		AddBalance(ctx context.Context, in entity.AddBalanceIn) (acc entity.Account, err error)
		GetByID(ctx context.Context, id uuid.UUID) (entity.Account, error)
	}
)
