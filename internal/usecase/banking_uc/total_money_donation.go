package banking_uc

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

func (uc *UseCase) GetTotalMoneyDonation(ctx context.Context, accountID uuid.UUID) (out int64, err error) {
	amount, err := uc.transaction.GetTotalDonationsAmount(ctx, accountID)
	if err != nil {
		return 0, fmt.Errorf("get total donations amount via service: %w", err)
	}
	return amount, nil
}
