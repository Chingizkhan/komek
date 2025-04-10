package fundraise

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"komek/internal/domain/fundraise/entity"
	"komek/internal/usecase"
)

type UseCase struct {
	funds   usecase.FundraiseService
	account usecase.AccountService
	client  usecase.ClientService
	tr      usecase.Transactional
}

func New(
	funds usecase.FundraiseService,
	account usecase.AccountService,
	client usecase.ClientService,
	tr usecase.Transactional,
) *UseCase {
	return &UseCase{
		funds:   funds,
		account: account,
		client:  client,
		tr:      tr,
	}
}

func (uc *UseCase) List(ctx context.Context) ([]entity.ListOut, error) {
	funds, err := uc.funds.ListActive(ctx)
	if err != nil {
		return nil, fmt.Errorf("list active funds: %w", err)
	}

	res := make([]entity.ListOut, 0, len(funds))

	for _, fund := range funds {
		acc, err := uc.account.GetByID(ctx, fund.AccountID)
		if err != nil {
			return nil, fmt.Errorf("get account: %w", err)
		}

		client, err := uc.client.GetByID(ctx, acc.Owner)
		if err != nil {
			return nil, fmt.Errorf("get client: %w", err)
		}

		res = append(res, entity.ListOut{
			ID:         fund.ID,
			Name:       client.Name,
			ImageUrl:   client.ImageURL,
			City:       client.City,
			Categories: client.Categories.Names(),
			Goal:       fund.Goal,
			Collected:  fund.Collected,
		})
	}

	return res, nil
}

func (uc *UseCase) GetByID(ctx context.Context, id uuid.UUID) (entity.GetOut, error) {
	fundraise, err := uc.funds.GetByID(ctx, id)
	if err != nil {
		return entity.GetOut{}, fmt.Errorf("get fundraise: %w", err)
	}

	acc, err := uc.account.GetByID(ctx, fundraise.AccountID)
	if err != nil {
		return entity.GetOut{}, fmt.Errorf("get account: %w", err)
	}

	client, err := uc.client.GetByID(ctx, acc.Owner)
	if err != nil {
		return entity.GetOut{}, fmt.Errorf("get client: %w", err)
	}

	return entity.GetOut{
		ID:          fundraise.ID,
		Name:        client.Name,
		ImageUrl:    client.ImageURL,
		City:        client.City,
		Categories:  client.Categories.Names(),
		Goal:        fundraise.Goal,
		Collected:   fundraise.Collected,
		Description: client.Description,
	}, nil
}
