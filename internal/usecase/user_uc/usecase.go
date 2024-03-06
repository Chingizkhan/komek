package user_uc

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"komek/internal/domain"
	"komek/internal/dto"
	"komek/internal/service/token"
	"time"
)

type UseCase struct {
	r                    UserRepository
	tr                   Transactional
	hasher               Hasher
	session              SessionRepository
	tokenMaker           token.Maker
	accessTokenLifetime  time.Duration
	refreshTokenLifetime time.Duration
}

func New(
	r UserRepository,
	tr Transactional,
	hasher Hasher,
	session SessionRepository,
	tokenMaker token.Maker,
	accessTokenLifetime, refreshTokenLifetime time.Duration,
) *UseCase {
	return &UseCase{
		r,
		tr,
		hasher,
		session,
		tokenMaker,
		accessTokenLifetime,
		refreshTokenLifetime,
	}
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
	accessToken, accessPayload, err := u.tokenMaker.CreateToken(user.ID, u.accessTokenLifetime)
	if err != nil {
		return nil, fmt.Errorf("create access token: %w", err)
	}

	// get refresh token
	refreshToken, refreshPayload, err := u.tokenMaker.CreateToken(user.ID, u.refreshTokenLifetime)
	if err != nil {
		return nil, fmt.Errorf("create refresh token: %w", err)
	}

	session := domain.Session{
		ID:           refreshPayload.ID,
		UserID:       user.ID,
		RefreshToken: refreshToken,
		UserAgent:    "",
		ClientIp:     "",
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	}

	// save session_repo
	createdSession, err := u.session.Save(ctx, nil, session)
	if err != nil {
		return nil, fmt.Errorf("session.Save: %w", err)
	}

	return &dto.UserLoginResponse{
		SessionID:             createdSession.ID,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
		User: dto.UserResponse{
			ID:            user.ID,
			Name:          user.Name,
			Login:         user.Login,
			Email:         user.Email,
			EmailVerified: user.EmailVerified,
			Roles:         user.Roles,
			CreatedAt:     user.CreatedAt.Unix(),
			UpdatedAt:     user.UpdatedAt.Unix(),
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
