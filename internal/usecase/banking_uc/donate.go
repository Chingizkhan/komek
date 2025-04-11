package banking_uc

import (
	"context"
	"fmt"
	"komek/internal/service/banking/entity"
)

func (uc *UseCase) Donate(ctx context.Context, in entity.DonateIn) error {
	if err := uc.tr.ExecContext(ctx, func(ctx context.Context) error {
		return uc.donate(ctx, in)
	}); err != nil {
		return fmt.Errorf("exec transaction: %w", err)
	}

	return nil
}

func (uc *UseCase) donate(ctx context.Context, in entity.DonateIn) error {
	tr, err := uc.banking.Transfer(ctx, entity.TransferIn{
		ToAccountID:   in.ToAccountID,
		FromAccountID: in.FromAccountID,
		Amount:        in.Amount,
	})
	if err != nil {
		return fmt.Errorf("banking.Transfer: %w", err)
	}

	if err = uc.funds.Donate(ctx, in.FundraiseID, in.Amount); err != nil {
		return fmt.Errorf("funds.Donate: %w", err)
	}

	_ = tr
	return nil
}
