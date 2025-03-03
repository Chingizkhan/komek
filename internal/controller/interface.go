package controller

import (
	"context"
	"github.com/google/uuid"
	"komek/internal/domain"
	"komek/internal/domain/user/entity"
	"komek/internal/dto"
)

type (
	User interface {
		Register(ctx context.Context, req entity.RegisterIn) (entity.User, error)
		Get(ctx context.Context, req entity.GetIn) (entity.GetOut, error)
		Delete(ctx context.Context, req entity.DeleteIn) error
		ChangePassword(ctx context.Context, req entity.ChangePasswordIn) error
		Update(ctx context.Context, req entity.UpdateIn) (entity.User, error)
		Login(ctx context.Context, in entity.LoginIn) (*entity.LoginOut, error)
		RefreshTokens(ctx context.Context, in entity.RefreshTokensIn) (*entity.RefreshTokensOut, error)
		Logout(ctx context.Context) error
	}

	Banking interface {
		CreateAccount(ctx context.Context, in dto.CreateAccountIn) (domain.Account, error)
		GetAccount(ctx context.Context, accountID uuid.UUID) (out domain.Account, err error)
		ListAccounts(ctx context.Context, userID uuid.UUID) (out []domain.Account, err error)
		Transfer(ctx context.Context, in dto.TransferIn) (dto.TransferOut, error)
	}
)
