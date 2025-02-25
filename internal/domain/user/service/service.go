package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"komek/internal/domain/email"
	"komek/internal/domain/phone"
	"komek/internal/domain/user/entity"
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

	// create empty account and save it connected to user

	return s.r.Save(ctx, in.ToEntity())
}

func (s *Service) Update(ctx context.Context, req entity.UpdateIn) (entity.User, error) {
	return s.r.Update(ctx, req)
}

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	return s.r.Delete(ctx, id)
}

func (s *Service) Get(ctx context.Context, in entity.GetIn) (user entity.User, err error) {
	if in.ID != uuid.Nil {
		user, err = s.r.GetByID(ctx, in.ID)
		if err != nil {
			return entity.User{}, fmt.Errorf("get user by id: %w", err)
		}
	}
	if in.Phone != "" {
		user, err = s.r.GetByPhone(ctx, phone.Phone(in.Phone))
		if err != nil {
			return entity.User{}, fmt.Errorf("get user by phone: %w", err)
		}
	}
	if in.Login != "" {
		user, err = s.r.GetByLogin(ctx, in.Login)
		if err != nil {
			return entity.User{}, fmt.Errorf("get user by login: %w", err)
		}
	}
	if in.Email != "" {
		user, err = s.r.GetByEmail(ctx, email.Email(in.Email))
		if err != nil {
			return entity.User{}, fmt.Errorf("get user by email: %w", err)
		}
	}
	if in.AccountID != 0 {
		user, err = s.r.GetByAccount(ctx, in.AccountID)
		if err != nil {
			return entity.User{}, fmt.Errorf("get user by account: %w", err)
		}
	}
	return
}
