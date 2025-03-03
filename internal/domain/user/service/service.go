package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"komek/internal/domain/user/entity"
	"log"
)

type Service struct {
	r UserRepository
}

func New(r UserRepository) *Service {
	return &Service{
		r: r,
	}
}

func (s *Service) Register(ctx context.Context, in entity.RegisterIn) (entity.User, error) {
	if err := in.Validate(); err != nil {
		return entity.User{}, fmt.Errorf("validate: %w", err)
	}

	return s.r.Save(ctx, in.ToEntity())
}

func (s *Service) Update(ctx context.Context, req entity.UpdateIn) (entity.User, error) {
	return s.r.Update(ctx, req)
}

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	return s.r.Delete(ctx, id)
}

func (s *Service) Get(ctx context.Context, in entity.GetIn) (user entity.User, err error) {
	if err = in.Validate(); err != nil {
		return user, fmt.Errorf("validate: %w", err)
	}

	switch {
	case in.ID != uuid.Nil:
		user, err = s.r.GetByID(ctx, in.ID)
		if err != nil {
			return user, fmt.Errorf("get user by id via repo: %w", err)
		}
	case in.Phone != "":
		log.Println("in.Phone:", in.Phone)
		user, err = s.r.GetByPhone(ctx, in.Phone)
		if err != nil {
			return user, fmt.Errorf("get user by phone via repo: %w", err)
		}
	case in.Login != "":
		user, err = s.r.GetByLogin(ctx, in.Login)
		if err != nil {
			return user, fmt.Errorf("get user by login via repo: %w", err)
		}
	case in.Email != "":
		user, err = s.r.GetByEmail(ctx, in.Email)
		if err != nil {
			return user, fmt.Errorf("get user by email via repo: %w", err)
		}
	case in.AccountID != uuid.Nil:
		user, err = s.r.GetByAccount(ctx, in.AccountID)
		if err != nil {
			return user, fmt.Errorf("get user by account via repo: %w", err)
		}
	}
	return
}
