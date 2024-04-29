package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"komek/internal/domain"
	"komek/internal/dto"
)

type (
	Transactional interface {
		Exec(ctx context.Context, fn func(tx pgx.Tx) error) error
	}

	BankingService interface {
		CreateAccount(ctx context.Context, in dto.CreateAccountIn) (domain.Account, error)
		InfoAccount(ctx context.Context, accountID uuid.UUID) (out domain.Account, err error)
		Transfer(ctx context.Context, in dto.TransferIn) (dto.TransferOut, error)
	}
)
