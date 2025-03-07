package banking_uc

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	banking "komek/internal/service/banking/entity"
)

type DonateIn struct {
	FromAccountID uuid.UUID
	ToAccountID   uuid.UUID
	Amount        int64
}

func (uc *UseCase) Donate(ctx context.Context, in DonateIn) error {
	tr, err := uc.banking.Transfer(ctx, banking.TransferIn{
		ToAccountID:   in.ToAccountID,
		FromAccountID: in.FromAccountID,
		Amount:        in.Amount,
	})
	if err != nil {
		return fmt.Errorf("banking.Transfer: %w", err)
	}

	// accumulations
	//

	return nil
}
