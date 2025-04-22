package banking_uc

import (
	"context"
	"github.com/google/uuid"
	"komek/internal/domain/transaction/entity"
	"komek/internal/errs"
)

type FindDonationsIn struct {
	FromAccountID uuid.UUID
	ToAccountID   uuid.UUID
}

func (uc *UseCase) FindDonations(ctx context.Context, in FindDonationsIn) (out []entity.Donation, err error) {
	switch {
	case in.FromAccountID != uuid.Nil && in.ToAccountID != uuid.Nil:
		return uc.transaction.FindDonationsByAccounts(ctx, in.FromAccountID, in.ToAccountID)
	case in.FromAccountID != uuid.Nil:
		return uc.transaction.GetDonationsByAccountID(ctx, in.FromAccountID)
	case in.ToAccountID != uuid.Nil:
		return uc.transaction.GetDonationsByAccountID(ctx, in.ToAccountID)
	default:
		return nil, errs.FindParameterNotSpecified
	}
}
