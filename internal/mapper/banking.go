package mapper

import (
	"fmt"
	"github.com/google/uuid"
	"komek/internal/domain"
	"komek/internal/service/banking_old/pb"
)

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
		ID:          id,
		Owner:       ownerID,
		Balance:     acc.Balance,
		HoldBalance: acc.HoldBalance,
		Currency:    domain.Currency(acc.Currency),
		Country:     domain.Country(acc.Country),
		Status:      domain.AccountStatus(acc.Status),
		CreatedAt:   acc.CreatedAt.AsTime(),
		UpdatedAt:   acc.UpdatedAt.AsTime(),
	}, nil
}

func ConvAccountsProtoToDomain(accounts []*pb.Account) ([]domain.Account, error) {
	res := make([]domain.Account, 0, len(accounts))
	for _, acc := range accounts {
		accDomain, err := ConvAccountProtoToDomain(acc)
		if err != nil {
			return nil, fmt.Errorf("parse account from proto to domain: %w", err)
		}
		res = append(res, accDomain)
	}
	return res, nil
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
