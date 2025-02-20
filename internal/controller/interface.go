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
		Get(ctx context.Context, req dto.UserGetRequest) (entity.User, error)
		Delete(ctx context.Context, req dto.UserDeleteRequest) error
		ChangePassword(ctx context.Context, req dto.UserChangePasswordRequest) error
		Update(ctx context.Context, req dto.UserUpdateRequest) (entity.User, error)
		Login(ctx context.Context, in dto.UserLoginRequest) (*dto.UserLoginResponse, error)
		RefreshTokens(ctx context.Context, in dto.UserRefreshTokensIn) (*dto.UserRefreshTokensOut, error)
		Logout(ctx context.Context) error
	}

	Banking interface {
		CreateAccount(ctx context.Context, in dto.CreateAccountIn) (domain.Account, error)
		GetAccount(ctx context.Context, accountID uuid.UUID) (out domain.Account, err error)
		ListAccounts(ctx context.Context, userID uuid.UUID) (out []domain.Account, err error)
		Transfer(ctx context.Context, in dto.TransferIn) (dto.TransferOut, error)
	}
)
