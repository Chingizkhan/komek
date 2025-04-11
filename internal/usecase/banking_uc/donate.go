package banking_uc

import (
	"context"
	"fmt"
	"komek/internal/errs"
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
	fund, err := uc.funds.GetByID(ctx, in.FundraiseID)
	if err != nil {
		return fmt.Errorf("get fund: %w", err)
	}

	if !fund.IsActive {
		return errs.ErrFundIsNotActive
	}

	var withCache bool

	transactions, err := uc.FindTransactionsByAccounts(ctx, in.FromAccountID, in.ToAccountID)
	if err != nil {
		return fmt.Errorf("uc.FindTransactionsByAccounts: %w", err)
	}
	if len(transactions) == 0 {
		withCache = true
	}

	_, err = uc.banking.Transfer(ctx, entity.TransferIn{
		ToAccountID:   in.ToAccountID,
		FromAccountID: in.FromAccountID,
		Amount:        in.Amount,
	})
	if err != nil {
		return fmt.Errorf("banking.Transfer: %w", err)
	}

	if err = uc.funds.Donate(ctx, in.FundraiseID, in.Amount, withCache); err != nil {
		return fmt.Errorf("funds.Donate: %w", err)
	}

	achieved, err := uc.funds.IsGoalAchieved(ctx, in.FundraiseID)
	if err != nil {
		return fmt.Errorf("funds.IsGoalAchieved: %w", err)
	}

	if !achieved {
		return nil
	}

	if err = uc.funds.Close(ctx, in.FundraiseID); err != nil {
		return fmt.Errorf("funds.Close: %w", err)
	}

	return nil
}
