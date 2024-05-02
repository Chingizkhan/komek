package banking

import (
	"context"
	"fmt"
	"komek/internal/domain"
	"komek/internal/dto"
	"komek/internal/mapper"
	"komek/internal/service/banking/pb"
)

func (s *Banking) CreateAccount(ctx context.Context, in dto.CreateAccountIn) (out domain.Account, err error) {
	response, err := s.client.CreateAccount(ctx, &pb.CreateAccountIn{
		Owner:       in.Owner.String(),
		Currency:    in.Currency,
		Country:     in.Country,
		Balance:     in.Balance,
		HoldBalance: in.HoldBalance,
	})
	if err != nil {
		return out, fmt.Errorf("create account: %w", err)
	}

	account, err := mapper.ConvAccountProtoToDomain(response.Account)
	if err != nil {
		return out, fmt.Errorf("conv account: %w", err)
	}

	return account, nil
}
