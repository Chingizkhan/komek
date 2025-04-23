package banking_uc

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	donation "komek/internal/domain/donation/entity"
	"komek/internal/domain/transaction/entity"
	"komek/internal/errs"
)

type FindDonationsIn struct {
	FromAccountID uuid.UUID
	ToAccountID   uuid.UUID
}

func (uc *UseCase) FindDonations(ctx context.Context, in FindDonationsIn) (out []donation.Donation, err error) {
	transactions := make([]entity.Transaction, 0, 32)

	switch {
	case in.FromAccountID != uuid.Nil && in.ToAccountID != uuid.Nil:
		transactions, err = uc.transaction.FindTransactionsByAccounts(ctx, in.FromAccountID, in.ToAccountID)
	case in.FromAccountID != uuid.Nil:
		transactions, err = uc.transaction.GetByAccountID(ctx, in.FromAccountID)
	case in.ToAccountID != uuid.Nil:
		transactions, err = uc.transaction.GetByAccountID(ctx, in.ToAccountID)
	default:
		return nil, errs.FindParameterNotSpecified
	}
	if err != nil {
		return nil, fmt.Errorf("find transactions: %w", err)
	}

	out = make([]donation.Donation, 0, len(transactions))
	for _, transaction := range transactions {
		donat, err := uc.donation.GetByTransactionID(ctx, transaction.ID)
		if err != nil {
			return nil, fmt.Errorf("get transaction donation: %w", err)
		}
		out = append(out, donat)
	}

	return out, nil
}
