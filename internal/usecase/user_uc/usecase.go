package user_uc

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"komek/internal/domain"
	"komek/internal/dto"
	"komek/internal/repos/tx"
)

type UseCase struct {
	r      UserRepository
	tr     Transactional
	hasher Hasher
	tx     tx.Tx
}

func New(r UserRepository, tr Transactional, hasher Hasher, tx tx.Tx) *UseCase {
	return &UseCase{r, tr, hasher, tx}
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

func (u *UseCase) Login(ctx context.Context) error {
	return nil
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
