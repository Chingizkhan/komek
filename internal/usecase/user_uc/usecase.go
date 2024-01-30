package user_uc

import (
	"context"
	"fmt"
	"komek/internal/domain"
	"komek/internal/dto"
)

type UseCase struct {
	r      UserRepository
	tr     Transactional
	hasher Hasher
}

func New(r UserRepository, tr Transactional, hasher Hasher) *UseCase {
	return &UseCase{r, tr, hasher}
}

func (u *UseCase) Register(ctx context.Context, req dto.UserRegisterRequest) error {
	tx, err := u.tr.Start(ctx)
	if err != nil {
		return fmt.Errorf("u.tr.Start - %w", err)
	}
	defer tx.Rollback(ctx)

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

	err = u.r.Save(ctx, tx, user)
	if err != nil {
		return fmt.Errorf("u.r.Save - %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("tx.Commit - %w", err)
	}
	return nil
}

func (u *UseCase) Login(ctx context.Context) error {
	tx, err := u.tr.Start(ctx)
	if err != nil {
		return fmt.Errorf("u.tr.Start - %w", err)
	}
	defer tx.Rollback(ctx)

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("tx.Commit - %w", err)
	}
	return nil
}

func (u *UseCase) Logout(ctx context.Context) error {
	tx, err := u.tr.Start(ctx)
	if err != nil {
		return fmt.Errorf("u.tr.Start - %w", err)
	}
	defer tx.Rollback(ctx)

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("tx.Commit - %w", err)
	}
	return nil
}

func (u *UseCase) Delete(ctx context.Context, req dto.UserDeleteRequest) error {
	tx, err := u.tr.Start(ctx)
	if err != nil {
		return fmt.Errorf("u.tr.Start - %w", err)
	}
	defer tx.Rollback(ctx)

	err = u.r.Delete(ctx, tx, req.ID)
	if err != nil {
		return fmt.Errorf("u.r.Delete - %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("tx.Commit - %w", err)
	}
	return nil
}

func (u *UseCase) ChangePassword(ctx context.Context, req dto.UserChangePasswordRequest) error {
	tx, err := u.tr.Start(ctx)
	if err != nil {
		return fmt.Errorf("u.tr.Start - %w", err)
	}
	defer tx.Rollback(ctx)

	passwordHash, err := u.hasher.Hash(req.Password)
	if err != nil {
		return fmt.Errorf("u.hasher.Hash - %w", err)
	}

	_, err = u.r.UpdatePasswordHash(ctx, tx, req.ID, passwordHash)
	if err != nil {
		return fmt.Errorf("u.r.UpdatePasswordHash - %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("tx.Commit - %w", err)
	}
	return nil
}

func (u *UseCase) Update(ctx context.Context, req dto.UserUpdateRequest) error {
	tx, err := u.tr.Start(ctx)
	if err != nil {
		return fmt.Errorf("u.tr.Start - %w", err)
	}
	defer tx.Rollback(ctx)

	if req.Name != "" {
		_, err = u.r.UpdateName(ctx, tx, req.ID, req.Name)
		if err != nil {
			return fmt.Errorf("u.r.UpdateName - %w", err)
		}
	}
	if req.Login != "" {
		_, err = u.r.UpdateLogin(ctx, tx, req.ID, req.Login)
		if err != nil {
			return fmt.Errorf("u.r.UpdateLogin - %w", err)
		}
	}
	if req.Email != "" {
		_, err = u.r.UpdateEmail(ctx, tx, req.ID, req.Email)
		if err != nil {
			return fmt.Errorf("u.r.UpdateEmail - %w", err)
		}
	}
	if req.EmailVerified {
		_, err = u.r.UpdateEmailVerified(ctx, tx, req.ID, req.EmailVerified)
		if err != nil {
			return fmt.Errorf("u.r.UpdateEmailVerified - %w", err)
		}
	}
	if req.Phone != "" {
		_, err = u.r.UpdatePhone(ctx, tx, req.ID, req.Phone)
		if err != nil {
			return fmt.Errorf("u.r.UpdatePhone - %w", err)
		}
	}
	if req.PasswordHash != "" {
		_, err = u.r.UpdatePasswordHash(ctx, tx, req.ID, req.PasswordHash)
		if err != nil {
			return fmt.Errorf("u.r.UpdatePasswordHash - %w", err)
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("tx.Commit - %w", err)
	}
	return nil
}
