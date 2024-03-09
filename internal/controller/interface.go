package controller

import (
	"context"
	"komek/internal/domain"
	"komek/internal/dto"
)

type (
	User interface {
		Register(ctx context.Context, req dto.UserRegisterRequest) (domain.User, error)
		Get(ctx context.Context, req dto.UserGetRequest) (domain.User, error)
		Delete(ctx context.Context, req dto.UserDeleteRequest) error
		ChangePassword(ctx context.Context, req dto.UserChangePasswordRequest) error
		Update(ctx context.Context, req dto.UserUpdateRequest) (domain.User, error)
		Login(ctx context.Context, in dto.UserLoginRequest) (*dto.UserLoginResponse, error)
		RefreshTokens(ctx context.Context, in dto.UserRefreshTokensIn) (*dto.UserRefreshTokensOut, error)
		Logout(ctx context.Context) error
	}

	Banking interface {
		CreateAccount(ctx context.Context, in dto.CreateAccountIn) (domain.Account, error)
		GetAccount(ctx context.Context, in dto.GetAccountIn) (domain.Account, error)
		Transfer(ctx context.Context, in dto.TransferIn) (dto.TransferOut, error)
	}
)
