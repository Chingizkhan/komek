package banking_uc

import (
	"context"
	"fmt"
	"komek/db/sqlc"
	"komek/internal/dto"
	"komek/internal/mapper"
)

func (s *UseCase) Transfer(ctx context.Context, in dto.TransferIn) (dto.TransferOut, error) {
	var result dto.TransferOut

	err := s.tx.Exec(ctx, func(q *sqlc.Queries) error {
		transfer, err := q.CreateTransfer(ctx, sqlc.CreateTransferParams{
			FromAccountID: in.FromAccountID,
			ToAccountID:   in.ToAccountID,
			Amount:        in.Amount,
		})
		if err != nil {
			return fmt.Errorf("create transfer: %w", err)
		}

		fromEntry, err := q.CreateEntry(ctx, sqlc.CreateEntryParams{
			AccountID: in.FromAccountID,
			Amount:    -in.Amount,
		})
		if err != nil {
			return fmt.Errorf("create from entry: %w", err)
		}

		toEntry, err := q.CreateEntry(ctx, sqlc.CreateEntryParams{
			AccountID: in.ToAccountID,
			Amount:    in.Amount,
		})
		if err != nil {
			return fmt.Errorf("create to entry: %w", err)
		}

		var fromAcc sqlc.Account
		var toAcc sqlc.Account

		if in.FromAccountID < in.ToAccountID {
			fromAcc, toAcc, err = s.addMoney(ctx, q, in.FromAccountID, -in.Amount, in.ToAccountID, in.Amount)
			if err != nil {
				return fmt.Errorf("addMoney 1: %w", err)
			}
		} else {
			toAcc, fromAcc, err = s.addMoney(ctx, q, in.ToAccountID, in.Amount, in.FromAccountID, -in.Amount)
			if err != nil {
				return fmt.Errorf("addMoney 2: %w", err)
			}
		}

		result.Transfer = mapper.ConvTransferToDomain(transfer)
		result.FromEntry = mapper.ConvEntryToDomain(fromEntry)
		result.ToEntry = mapper.ConvEntryToDomain(toEntry)
		result.FromAccount = mapper.ConvAccountToDomain(fromAcc)
		result.ToAccount = mapper.ConvAccountToDomain(toAcc)

		return nil
	})
	if err != nil {
		return dto.TransferOut{}, fmt.Errorf("tx.Exec: %w", err)
	}

	return result, nil
}

func (s *UseCase) addMoney(
	ctx context.Context,
	q *sqlc.Queries,
	accID1 int64,
	amount1 int64,
	accID2 int64,
	amount2 int64,
) (acc1 sqlc.Account, acc2 sqlc.Account, err error) {
	acc1, err = q.AddAccountBalance(ctx, sqlc.AddAccountBalanceParams{
		Amount: amount1,
		ID:     accID1,
	})
	if err != nil {
		return
	}

	acc2, err = q.AddAccountBalance(ctx, sqlc.AddAccountBalanceParams{
		Amount: amount2,
		ID:     accID2,
	})
	if err != nil {
		return
	}

	return
}
