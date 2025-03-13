package banking_uc

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	account "komek/internal/domain/account/entity"
	banking "komek/internal/service/banking/entity"
)

func (uc *UseCase) Transfer(ctx context.Context, in banking.TransferIn) (out banking.Transaction, err error) {
	var (
		toAccount   account.Account
		fromAccount account.Account
	)
	if in.ToAccountID == uuid.Nil {
		if toAccount, err = uc.account.GetByOwnerID(ctx, in.ToUserID); err != nil {
			return out, fmt.Errorf("get to_account by user_id via service: %w", err)
		}
		in.ToAccountID = toAccount.ID
	}

	if in.FromAccountID == uuid.Nil {
		if fromAccount, err = uc.account.GetByOwnerID(ctx, in.FromUserID); err != nil {
			return out, fmt.Errorf("get from_account by user_id via service: %w", err)
		}
		in.FromAccountID = fromAccount.ID
	}

	tr, err := uc.banking.Transfer(ctx, in)
	if err != nil {
		return out, fmt.Errorf("banking.Transfer: %w", err)
	}
	return tr, nil
}
