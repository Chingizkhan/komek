package service

import (
	"context"
	"fmt"
	account "komek/internal/domain/account/entity"
	operation "komek/internal/domain/operation/entity"
	"komek/internal/errs"
	"komek/internal/service/banking/entity"
)

func (s *Service) Withdraw(ctx context.Context, in entity.WithdrawIn) (op operation.Operation, err error) {
	acc, err := s.account.GetByID(ctx, in.AccountID)
	if err != nil {
		return op, fmt.Errorf("get account by id: %w", err)
	}

	if !acc.Status.IsActive() {
		return op, errs.InactiveAccountStatus
	}

	if in.Amount <= 0 {
		return op, errs.InsufficientAmount
	}

	if acc.Balance < in.Amount {
		return op, errs.NotEnoughBalance
	}

	if op, err = s.withdraw(ctx, in, acc); err != nil {
		return op, err
	}

	return op, nil
}

func (s *Service) withdraw(ctx context.Context, in entity.WithdrawIn, acc account.Account) (op operation.Operation, err error) {
	if op, err = s.operation.Create(ctx, operation.CreateIn{
		TransactionID: in.TransactionID,
		AccountID:     acc.ID,
		Type:          operation.TypeWithdraw,
		Amount:        in.Amount,
		BalanceBefore: acc.Balance,
		BalanceAfter:  acc.Balance - in.Amount,
	}); err != nil {
		return op, fmt.Errorf("create operation via operation service: %w", err)
	}

	if acc, err = s.account.AddBalance(ctx, account.AddBalanceIn{
		Amount:    -in.Amount,
		AccountID: acc.ID,
	}); err != nil {
		return op, fmt.Errorf("add balance via account service: %w", err)
	}

	return op, nil
}
