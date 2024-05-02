package banking

import (
	"context"
	"fmt"
	"komek/internal/domain"
	"komek/internal/mapper"
	"komek/internal/service/banking/pb"
)

func (s *Banking) ListAccounts(ctx context.Context, accountIDs []string) (out []domain.Account, err error) {
	response, err := s.client.ListAccounts(ctx, &pb.ListAccountsIn{
		AccountIds: accountIDs,
	})
	if err != nil {
		return out, fmt.Errorf("create account: %w", err)
	}
	accounts, err := mapper.ConvAccountsProtoToDomain(response.Accounts)
	if err != nil {
		return nil, fmt.Errorf("convert accounts from proto to domain: %w", err)
	}
	return accounts, nil
}
