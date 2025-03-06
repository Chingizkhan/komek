package service

import (
	"context"
	"fmt"
	account "komek/internal/domain/account/entity"
	operation "komek/internal/domain/operation/entity"
	transaction "komek/internal/domain/transaction/entity"
	"komek/internal/errs"
	entity "komek/internal/service/banking/entity"
)

type Service struct {
	transaction TransactionService
	operation   OperationService
	account     AccountService
}

func New(
	operation OperationService,
	transaction TransactionService,
	account AccountService,
) *Service {
	return &Service{
		operation:   operation,
		transaction: transaction,
		account:     account,
	}
}

func (s *Service) Transfer(ctx context.Context, in entity.TransferIn) (tr entity.Transaction, err error) {
	transaction, err := s.transaction.Create(ctx, transaction.Transaction{
		FromAccountID: in.FromAccountID,
		ToAccountID:   in.ToAccountID,
		Amount:        in.Amount,
	})
	if err != nil {
		return tr, fmt.Errorf("create transaction via service: %w", err)
	}

	withdraw, err := s.Withdraw(ctx, entity.WithdrawIn{
		TransactionID: transaction.ID,
		AccountID:     in.FromAccountID,
		Amount:        in.Amount,
	})
	if err != nil {
		return tr, fmt.Errorf("withdraw via service: %w", err)
	}

	refill, err := s.Refill(ctx, entity.RefillIn{
		TransactionID: transaction.ID,
		AccountID:     in.ToAccountID,
		Amount:        in.Amount,
	})
	if err != nil {
		return tr, fmt.Errorf("refill via service: %w", err)
	}

	return entity.NewTransaction(transaction, withdraw, refill), nil
}

func (s *Service) Withdraw(ctx context.Context, in entity.WithdrawIn) (op operation.Operation, err error) {
	acc, err := s.account.GetByID(ctx, in.AccountID)
	if err != nil {
		return op, fmt.Errorf("get account by id: %w", err)
	}

	if in.Amount < 0 {
		return op, errs.InsufficientAmount
	}

	if acc.Balance < in.Amount {
		return op, errs.NotEnoughBalance
	}

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

func (s *Service) Refill(ctx context.Context, in entity.RefillIn) (op operation.Operation, err error) {
	acc, err := s.account.GetByID(ctx, in.AccountID)
	if err != nil {
		return op, fmt.Errorf("get account by id: %w", err)
	}

	if in.Amount < 0 {
		return op, errs.InsufficientAmount
	}

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

	if _, err = s.account.AddBalance(ctx, account.AddBalanceIn{
		Amount:    in.Amount,
		AccountID: acc.ID,
	}); err != nil {
		return op, fmt.Errorf("add balance via account service: %w", err)
	}

	return op, nil
}
