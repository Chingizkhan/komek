package banking

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"komek/internal/domain"
	"komek/internal/service/banking/pb"
)

func (s *Banking) InfoAccount(ctx context.Context, accountID uuid.UUID) (out domain.Account, err error) {
	response, err := s.client.InfoAccount(ctx, &pb.InfoAccountIn{AccountId: accountID.String()})
	if err != nil {
		return out, fmt.Errorf("info account: %w", err)
	}

	acc := response.Account

	id, err := uuid.Parse(acc.Id)
	if err != nil {
		return out, fmt.Errorf("conv acc_id: %w", err)
	}
	ownerID, err := uuid.Parse(acc.Owner)
	if err != nil {
		return out, fmt.Errorf("conv oswner_id: %w", err)
	}

	return domain.Account{
		ID:          id,
		Owner:       ownerID,
		Balance:     acc.Balance,
		HoldBalance: acc.HoldBalance,
		Status:      domain.AccountStatus(acc.Status),
		Currency:    domain.Currency(acc.Currency),
		Country:     domain.Country(acc.Country),
		CreatedAt:   acc.CreatedAt.AsTime(),
		UpdatedAt:   acc.UpdatedAt.AsTime(),
	}, nil
}
