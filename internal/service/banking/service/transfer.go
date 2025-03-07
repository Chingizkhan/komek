package service

import (
	"context"
	"fmt"
	transaction "komek/internal/domain/transaction/entity"
	"komek/internal/errs"
	"komek/internal/service/banking/entity"
)

func (s *Service) Transfer(ctx context.Context, in entity.TransferIn) (tr entity.Transaction, err error) {
	trans, err := s.transaction.Create(ctx, transaction.Transaction{
		FromAccountID: in.FromAccountID,
		ToAccountID:   in.ToAccountID,
		Amount:        in.Amount,
	})
	if err != nil {
		return tr, fmt.Errorf("create transaction via service: %w", err)
	}

	accTo, err := s.account.GetByID(ctx, in.ToAccountID)
	if err != nil {
		return tr, fmt.Errorf("get to_account by id via service: %w", err)
	}
	if !accTo.Status.IsActive() {
		return tr, errs.InactiveAccountStatus
	}

	accFrom, err := s.account.GetByID(ctx, in.FromAccountID)
	if err != nil {
		return tr, fmt.Errorf("get from_account by id via service: %w", err)
	}
	if !accFrom.Status.IsActive() {
		return tr, errs.InactiveAccountStatus
	}

	if in.Amount <= 0 {
		return tr, errs.InsufficientAmount
	}

	if accTo.Currency != accFrom.Currency {
		return tr, errs.CurrencyMismatch
	}

	if accFrom.Balance < in.Amount {
		return tr, errs.NotEnoughBalance
	}

	withdraw, err := s.withdraw(
		ctx,
		entity.WithdrawIn{
			TransactionID: trans.ID,
			AccountID:     in.FromAccountID,
			Amount:        in.Amount,
		},
		accFrom,
	)
	if err != nil {
		return tr, fmt.Errorf("withdraw via service: %w", err)
	}

	refill, err := s.refill(
		ctx,
		entity.RefillIn{
			TransactionID: trans.ID,
			AccountID:     in.ToAccountID,
			Amount:        in.Amount,
		},
		accTo,
	)
	if err != nil {
		return tr, fmt.Errorf("refill via service: %w", err)
	}

	return entity.NewTransaction(trans, withdraw, refill), nil
}
