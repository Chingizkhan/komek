package mapper

import (
	"fmt"
	"github.com/google/uuid"
	"komek/db/sqlc"
	"komek/internal/domain"
	"komek/internal/service/banking/pb"
)

func ConvTransferToDomain(transfer sqlc.Transfer) domain.Transfer {
	return domain.Transfer{
		ID:            transfer.ID,
		FromAccountID: transfer.FromAccountID,
		ToAccountID:   transfer.ToAccountID,
		Amount:        transfer.Amount,
		CreatedAt:     transfer.CreatedAt.Time,
	}
}

func ConvEntryToDomain(entry sqlc.Entry) domain.Entry {
	return domain.Entry{
		ID:        entry.ID,
		AccountID: entry.AccountID,
		Amount:    entry.Amount,
		CreatedAt: entry.CreatedAt.Time,
	}
}

func ConvAccountToDomain(acc sqlc.Account) domain.Account {
	return domain.Account{
		//ID:        acc.ID,
		Owner:     acc.Owner.Bytes,
		Balance:   acc.Balance,
		Currency:  acc.Currency,
		CreatedAt: acc.CreatedAt.Time,
	}
}

func ConvAccountProtoToDomain(acc *pb.Account) (domain.Account, error) {
	id, err := uuid.Parse(acc.Id)
	if err != nil {
		return domain.Account{}, fmt.Errorf("parse transaction_id: %w", err)
	}
	ownerID, err := uuid.Parse(acc.Owner)
	if err != nil {
		return domain.Account{}, fmt.Errorf("parse transaction_id: %w", err)
	}
	return domain.Account{
		ID:        id,
		Owner:     ownerID,
		Balance:   acc.Balance,
		Currency:  acc.Currency,
		CreatedAt: acc.CreatedAt.AsTime(),
	}, nil
}

func ConvTransactionToDomain(tr *pb.Transaction) (out domain.Transaction, err error) {
	id, err := uuid.Parse(tr.Id)
	if err != nil {
		return out, fmt.Errorf("parse transaction_id: %w", err)
	}

	accID, err := uuid.Parse(tr.AccountId)
	if err != nil {
		return out, fmt.Errorf("parse accout_id: %w", err)
	}

	var refundedBy uuid.UUID
	if tr.RefundedBy != "" {
		refundedBy, err = uuid.Parse(tr.RefundedBy)
		if err != nil {
			return out, fmt.Errorf("parse refunded_by_id: %w", err)
		}
	}

	operations, err := ConvOperationsToDomain(tr.Operations)
	if err != nil {
		return out, fmt.Errorf("convert operations: %w", err)
	}

	return domain.Transaction{
		ID:         id,
		AccountID:  accID,
		Type:       domain.TransactionType(tr.Type),
		Amount:     tr.Amount,
		Operations: operations,
		RefundedBy: refundedBy,
		CreatedAt:  tr.CreatedAt.AsTime(),
	}, nil
}

func ConvOperationsToDomain(operations []*pb.Operation) ([]domain.Operation, error) {
	res := make([]domain.Operation, 0, len(operations))
	for _, op := range operations {
		operation, err := ConvOperationToDomain(op)
		if err != nil {
			return nil, fmt.Errorf("convert operation: %w", err)
		}
		res = append(res, operation)
	}
	return res, nil
}

func ConvOperationToDomain(op *pb.Operation) (domain.Operation, error) {
	id, err := uuid.Parse(op.Id)
	if err != nil {
		return domain.Operation{}, fmt.Errorf("parse operation_id: %w", err)
	}

	transactionID, err := uuid.Parse(op.TransactionId)
	if err != nil {
		return domain.Operation{}, fmt.Errorf("parse transaction_id: %w", err)
	}

	accountID, err := uuid.Parse(op.AccountId)
	if err != nil {
		return domain.Operation{}, fmt.Errorf("parse account_id: %w", err)
	}

	return domain.Operation{
		ID:                id,
		TransactionID:     transactionID,
		AccountID:         accountID,
		Type:              domain.OperationType(op.Type),
		Amount:            op.Amount,
		BalanceBefore:     op.BalanceBefore,
		BalanceAfter:      op.BalanceAfter,
		HoldBalanceBefore: op.HoldBalanceBefore,
		HoldBalanceAfter:  op.HoldBalanceAfter,
		CreatedAt:         op.CreatedAt.AsTime(),
	}, nil
}
