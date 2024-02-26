package user_uc

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"komek/internal/domain"
	"komek/internal/dto"
	"komek/internal/service/token"
	"time"
)

type UseCase struct {
	r          UserRepository
	tr         Transactional
	hasher     Hasher
	tokenMaker token.Maker
}

func New(r UserRepository, tr Transactional, hasher Hasher, tokenMaker token.Maker) *UseCase {
	return &UseCase{r, tr, hasher, tokenMaker}
}

func (u *UseCase) Register(ctx context.Context, req dto.UserRegisterRequest) error {
	passHash, err := u.hasher.Hash(req.Password)
	if err != nil {
		return fmt.Errorf("u.hasher.Hash - %w", err)
	}

	user := domain.User{
		Phone:        req.Phone,
		Login:        req.Login,
		Roles:        req.Roles,
		PasswordHash: passHash,
	}

	err = u.tr.Exec(ctx, func(tx pgx.Tx) error {

		if err = u.r.Save(ctx, tx, user); err != nil {
			return fmt.Errorf("u.r.Save - %w", err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("tr.Exec: %w", err)
	}

	return nil
}

func (u *UseCase) Get(ctx context.Context, req dto.UserGetRequest) (domain.User, error) {
	user, err := u.r.Get(ctx, nil, req.ID)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user: %w", err)
	}
	return user, nil
}

func (u *UseCase) Login(ctx context.Context, in dto.UserLoginRequest) (*dto.UserLoginResponse, error) {
	user, err := u.r.GetUserByLogin(ctx, nil, in.Login)
	if err != nil {
		return nil, fmt.Errorf("get user by login: %w", err)
	}

	// check password
	if !u.hasher.CheckHash(in.Password, user.PasswordHash) {
		return nil, ErrIncorrectPassword
	}

	// get access token
	accessToken, err := u.tokenMaker.CreateToken(user.ID, time.Minute*1)
	if err != nil {
		return nil, fmt.Errorf("tokenMaker.CreateToken: %w", err)
	}

	return &dto.UserLoginResponse{
		AccessToken: accessToken,
		User:        user,
	}, nil
}

func (u *UseCase) Logout(ctx context.Context) error {
	return nil
}

func (u *UseCase) Delete(ctx context.Context, req dto.UserDeleteRequest) error {
	err := u.tr.Exec(ctx, func(tx pgx.Tx) error {
		err := u.r.Delete(ctx, tx, req.ID)
		if err != nil {
			return fmt.Errorf("u.r.Delete - %w", err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("tr.Exec: %w", err)
	}
	return nil
}

func (u *UseCase) ChangePassword(ctx context.Context, req dto.UserChangePasswordRequest) error {
	passwordHash, err := u.hasher.Hash(req.Password)
	if err != nil {
		return fmt.Errorf("u.hasher.Hash - %w", err)
	}

	if err := u.tr.Exec(ctx, func(tx pgx.Tx) error {
		_, err = u.r.Update(ctx, tx, dto.UserUpdateRequest{
			ID:           req.ID,
			PasswordHash: passwordHash,
		})
		if err != nil {
			return fmt.Errorf("u.r.UpdatePasswordHash - %w", err)
		}
		return nil
	}); err != nil {
		return fmt.Errorf("tr.Exec: %w", err)
	}
	return nil
}

func (u *UseCase) Update(ctx context.Context, req dto.UserUpdateRequest) error {
	if err := u.tr.Exec(ctx, func(tx pgx.Tx) error {
		_, err := u.r.Update(ctx, tx, req)
		if err != nil {
			return fmt.Errorf("u.r.Update - %w", err)
		}
		return nil
	}); err != nil {
		return fmt.Errorf("tr.Exec: %w", err)
	}
	return nil
}
