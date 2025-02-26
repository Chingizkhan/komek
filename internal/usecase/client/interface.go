package client

import (
	"context"
	"github.com/google/uuid"
	"komek/internal/domain/client/entity"
)

type (
	ClientService interface {
		GetByID(ctx context.Context, id uuid.UUID) (entity.Client, error)
		List(ctx context.Context) (entity.Clients, error)
		Create(ctx context.Context, in entity.CreateIn) (client entity.Client, err error)
	}
)
