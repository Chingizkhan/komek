package banking_uc

import (
	"context"
	"github.com/google/uuid"
	"komek/internal/domain/transaction/entity"
	"komek/internal/errs"
)

type FindTransactionsIn struct {
	FromAccountID uuid.UUID
	ToAccountID   uuid.UUID
}

func (uc *UseCase) FindTransactions(ctx context.Context, in FindTransactionsIn) (out []entity.Transaction, err error) {
	switch {
	case in.FromAccountID != uuid.Nil && in.ToAccountID != uuid.Nil:
		return uc.transaction.FindTransactionsByAccounts(ctx, in.FromAccountID, in.ToAccountID)
	case in.FromAccountID != uuid.Nil:
		return uc.transaction.GetByAccountID(ctx, in.FromAccountID)
	case in.ToAccountID != uuid.Nil:
		return uc.transaction.GetByAccountID(ctx, in.ToAccountID)
	default:
		return nil, errs.FindParameterNotSpecified
	}
}
