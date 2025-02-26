package client

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"komek/internal/domain/client/entity"
	"komek/internal/usecase"
)

type UseCase struct {
	clientService ClientService
	tr            usecase.Transactional
}

func New(clientService ClientService, tr usecase.Transactional) *UseCase {
	return &UseCase{
		clientService: clientService,
		tr:            tr,
	}
}

func (uc *UseCase) List(ctx context.Context) (entity.Clients, error) {
	clients, err := uc.clientService.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("list clients via service: %w", err)
	}

	return clients, nil
}

func (uc *UseCase) GetByID(ctx context.Context, clientID uuid.UUID) (entity.Client, error) {
	client, err := uc.clientService.GetByID(ctx, clientID)
	if err != nil {
		return entity.Client{}, fmt.Errorf("get client by id via service: %w", err)
	}

	return client, nil
}
