package banking_uc

import (
	"context"
	"github.com/google/uuid"
	transaction "komek/internal/domain/transaction/entity"
)

func (uc *UseCase) FindTransactionsByAccounts(ctx context.Context, fromAccountID, toAccountID uuid.UUID) (out []transaction.Transaction, err error) {
	return uc.transaction.FindTransactionsByAccounts(ctx, fromAccountID, toAccountID)
}
