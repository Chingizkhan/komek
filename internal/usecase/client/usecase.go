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

func (uc *UseCase) CreateClient(ctx context.Context, in entity.CreateIn) (client entity.Client, err error) {
	if err = uc.tr.ExecContext(ctx, func(txCtx context.Context) error {
		if client, err = uc.client.Create(txCtx, in); err != nil {
			return fmt.Errorf("create client via service: %w", err)
		}
		return nil
	}); err != nil {
		return client, fmt.Errorf("exec via transaction: %w", err)
	}

	return client, nil
}

func (uc *UseCase) CreateFundraise(ctx context.Context, in fundraise.CreateIn) (fund fundraise.Fundraise, err error) {
	if err = uc.tr.ExecContext(ctx, func(txCtx context.Context) error {
		if fund, err = uc.fundraise.Create(txCtx, in); err != nil {
			return fmt.Errorf("create fundraise via service: %w", err)
		}
		return nil
	}); err != nil {
		return fund, fmt.Errorf("exec via transaction: %w", err)
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
