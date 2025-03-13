package client

import (
	"context"
	"github.com/google/uuid"
	client "komek/internal/domain/client/entity"
	fundraise "komek/internal/domain/fundraise/entity"
)

type (
	ClientService interface {
		GetByID(ctx context.Context, id uuid.UUID) (client.Client, error)
		List(ctx context.Context) (client.Clients, error)
		Create(ctx context.Context, in client.CreateIn) (client client.Client, err error)
	}

	FundraiseService interface {
		GetByID(ctx context.Context, id uuid.UUID) (fundraise.Fundraise, error)
		GetByAccountID(ctx context.Context, id uuid.UUID) ([]fundraise.Fundraise, error)
		Create(ctx context.Context, in fundraise.CreateIn) (fundraise.Fundraise, error)
		CreateType(ctx context.Context, name string) error
		ListActive(ctx context.Context) ([]fundraise.Fundraise, error)
	}
)
