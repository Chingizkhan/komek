package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"komek/internal/domain/client/entity"
)

type Service struct {
	clientRepo ClientRepository
}

func New(r ClientRepository) *Service {
	return &Service{
		clientRepo: r,
	}
}

func (s *Service) GetByID(ctx context.Context, id uuid.UUID) (entity.Client, error) {
	return s.clientRepo.GetByID(ctx, id)
}

func (s *Service) List(ctx context.Context) (entity.Clients, error) {
	return s.clientRepo.List(ctx)
}

func (s *Service) Create(ctx context.Context, in entity.CreateIn) (client entity.Client, err error) {
	client, err = s.clientRepo.Save(ctx, entity.Client{
		Name:          in.Name,
		Phone:         in.Phone,
		Email:         in.Email,
		Age:           in.Age,
		City:          in.City,
		Address:       in.Address,
		Description:   in.Description,
		Circumstances: in.Circumstances,
		ImageURL:      in.ImageURL,
	})
	if err != nil {
		return entity.Client{}, fmt.Errorf("create client via repo: %w", err)
	}

	if in.CategoryIDs == nil {
		return client, nil
	}

	if err = s.clientRepo.BindCategories(ctx, entity.BindCategories{
		ClientID:    client.ID,
		CategoryIDs: in.CategoryIDs,
	}); err != nil {
		return entity.Client{}, fmt.Errorf("bind categories via repo: %w", err)
	}

	client, err = s.clientRepo.GetByID(ctx, client.ID)
	if err != nil {
		return entity.Client{}, fmt.Errorf("get client via repo: %w", err)
	}

	return client, nil
}
