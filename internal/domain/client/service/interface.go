package service

import (
	"context"
	"github.com/google/uuid"
	"komek/internal/domain/client/entity"
)

type (
	ClientRepository interface {
		List(ctx context.Context) (entity.Clients, error)
		GetByID(ctx context.Context, id uuid.UUID) (entity.Client, error)
		Save(ctx context.Context, client entity.Client) (entity.Client, error)
		SaveCategories(ctx context.Context, categories entity.Categories) (entity.Categories, error)
		BindCategories(ctx context.Context, in entity.BindCategories) error
	}
)
