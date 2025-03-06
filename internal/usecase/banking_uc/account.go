package banking_uc

import (
	"context"
	"github.com/google/uuid"
	"komek/internal/domain"
	"komek/internal/dto"
)

func (uc *UseCase) CreateAccount(ctx context.Context, in dto.CreateAccountIn) (out domain.Account, err error) {
	panic("implement me")
	//out, err = uc.account.Create(ctx, in)
	//if err != nil {
	//	return out, fmt.Errorf("banking service -> create account: %w", err)
	//}
	//return
}

func (uc *UseCase) GetAccount(ctx context.Context, accID uuid.UUID) (out domain.Account, err error) {
	panic("implement me")
	//out, err = uc.banking.InfoAccount(ctx, accID)
	//if err != nil {
	//	return out, fmt.Errorf("banking service -> info account: %w", err)
	//}
	//return
}

func (uc *UseCase) ListAccounts(ctx context.Context, userID uuid.UUID) (out []domain.Account, err error) {
	panic("implement me")
	////todo: get accountIDs via userID
	//out, err = uc.banking.ListAccounts(ctx, []string{})
	//if err != nil {
	//	return out, fmt.Errorf("banking service -> list accounts: %w", err)
	//}
	//return
}
