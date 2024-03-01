package user_uc

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"komek/internal/domain"
	"komek/internal/dto"
	"komek/internal/service/token"
	"log"
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

func (u *UseCase) Register(ctx context.Context, req dto.UserRegisterRequest) (domain.User, error) {
	passHash, err := u.hasher.Hash(string(req.Password))
	if err != nil {
		return domain.User{}, fmt.Errorf("u.hasher.Hash - %w", err)
	}

	user := domain.User{
		Phone:        req.Phone,
		Login:        req.Login,
		Roles:        req.Roles,
		PasswordHash: passHash,
	}

	err = u.tr.Exec(ctx, func(tx pgx.Tx) error {

		if user, err = u.r.Save(ctx, tx, user); err != nil {
			return fmt.Errorf("u.r.Save - %w", err)
		}
		return nil
	})
	if err != nil {
		return domain.User{}, fmt.Errorf("tr.Exec: %w", err)
	}

	return user, nil
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
	accessToken, err := u.tokenMaker.CreateToken(user.ID, time.Minute*15)
	if err != nil {
		return nil, fmt.Errorf("tokenMaker.CreateToken: %w", err)
	}

	return &dto.UserLoginResponse{
		AccessToken: accessToken,
		User: dto.UserResponse{
			ID:            user.ID,
			Name:          user.Name,
			Login:         user.Login,
			Email:         user.Email,
			EmailVerified: user.EmailVerified,
			Roles:         user.Roles,
			CreatedAt:     user.CreatedAt,
			UpdatedAt:     user.UpdatedAt,
		},
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
	user, err := u.r.Get(ctx, nil, req.ID)
	if err != nil {
		return fmt.Errorf("get user: %w", err)
	}

	match := u.hasher.CheckHash(string(req.OldPassword), user.PasswordHash)
	log.Println("match:", match)
	if !match {
		return errors.New("wrong old password")
	}

	passwordHash, err := u.hasher.Hash(string(req.NewPassword))
	if err != nil {
		return fmt.Errorf("u.hasher.Hash - %w", err)
	}

	if err = u.tr.Exec(ctx, func(tx pgx.Tx) error {
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

func (u *UseCase) Update(ctx context.Context, req dto.UserUpdateRequest) (domain.User, error) {
	var (
		user domain.User
		err  error
	)
	if err = u.tr.Exec(ctx, func(tx pgx.Tx) error {
		user, err = u.r.Update(ctx, tx, req)
		if err != nil {
			return fmt.Errorf("u.r.Update - %w", err)
		}
		return nil
	}); err != nil {
		return domain.User{}, fmt.Errorf("tr.Exec: %w", err)
	}
	return user, nil
}
