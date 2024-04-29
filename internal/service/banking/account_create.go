package banking

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"komek/internal/domain"
	"komek/internal/dto"
	"komek/internal/service/banking/pb"
)

func (s *Banking) CreateAccount(ctx context.Context, in dto.CreateAccountIn) (out domain.Account, err error) {
	response, err := s.client.CreateAccount(ctx, &pb.CreateAccountIn{
		Currency: in.Currency,
		Country:  in.Country,
	})
	if err != nil {
		return
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
		ID:        id,
		Owner:     ownerID,
		Balance:   acc.Balance,
		Currency:  domain.Currency(acc.Currency),
		Country:   domain.Country(acc.Country),
		Status:    domain.AccountStatus(acc.Status),
		CreatedAt: acc.CreatedAt.AsTime(),
		UpdatedAt: acc.UpdatedAt.AsTime(),
	}, nil
}
