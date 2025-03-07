package client

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"komek/internal/domain/client/entity"
	fundraise "komek/internal/domain/fundraise/entity"
	"komek/internal/usecase"
)

type UseCase struct {
	client    ClientService
	fundraise FundraiseService
	tr        usecase.Transactional
}

func New(client ClientService, fundraise FundraiseService, tr usecase.Transactional) *UseCase {
	return &UseCase{
		client:    client,
		fundraise: fundraise,
		tr:        tr,
	}
}

func (uc *UseCase) ListClients(ctx context.Context) (entity.Clients, error) {
	clients, err := uc.client.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("list clients via service: %w", err)
	}

	return clients, nil
}

func (uc *UseCase) GetClientByID(ctx context.Context, clientID uuid.UUID) (entity.Client, error) {
	client, err := uc.client.GetByID(ctx, clientID)
	if err != nil {
		return entity.Client{}, fmt.Errorf("get client by id via service: %w", err)
	}

	return client, nil
}

func (uc *UseCase) CreateClient(ctx context.Context, in entity.CreateIn) (entity.Client, error) {
	client, err := uc.client.Create(ctx, in)
	if err != nil {
		return entity.Client{}, fmt.Errorf("create client via service: %w", err)
	}

	return client, nil
}

func (uc *UseCase) CreateFundraise(ctx context.Context, in fundraise.CreateIn) (fundraise.Fundraise, error) {
	fund, err := uc.fundraise.Create(ctx, in)
	if err != nil {
		return fundraise.Fundraise{}, fmt.Errorf("create fundraise via service: %w", err)
	}

	return fund, nil
}

func (uc *UseCase) GetFundraiseByID(ctx context.Context, id uuid.UUID) (fundraise.Fundraise, error) {
	fund, err := uc.fundraise.GetByID(ctx, id)
	if err != nil {
		return fundraise.Fundraise{}, fmt.Errorf("get fundraise_by_id via service: %w", err)
	}

	return fund, nil
}

func (uc *UseCase) ListActiveFundraises(ctx context.Context) ([]fundraise.Fundraise, error) {
	funds, err := uc.fundraise.ListActive(ctx)
	if err != nil {
		return nil, fmt.Errorf("list active fundraises via service: %w", err)
	}

	return funds, nil
}
