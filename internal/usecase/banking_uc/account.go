package banking_uc

import (
	"context"
	"github.com/google/uuid"
	"komek/internal/domain/account/entity"
)

func (uc *UseCase) CreateAccount(ctx context.Context, in entity.CreateIn) (out entity.Account, err error) {
	return uc.account.Create(ctx, in)
}

func (uc *UseCase) GetAccount(ctx context.Context, accID uuid.UUID) (out entity.Account, err error) {
	return uc.account.GetByID(ctx, accID)
}

func (uc *UseCase) GetAccountByUserID(ctx context.Context, userID uuid.UUID) (out entity.Account, err error) {
	return uc.account.GetByOwnerID(ctx, userID)
}
