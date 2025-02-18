package user

import (
	"context"
	"errors"
	"fmt"
	"komek/internal/domain"
	"komek/internal/domain/user/entity"
	"komek/internal/dto"
	"komek/internal/errs"
	"komek/internal/service/token"
	"time"
)

type UseCase struct {
	s                    Service
	tr                   Transactional
	hasher               Hasher
	session              SessionRepository
	im                   IdentityManager
	tokenMaker           token.Maker
	accessTokenLifetime  time.Duration
	refreshTokenLifetime time.Duration
}

func New(
	s Service,
	tr Transactional,
	hasher Hasher,
	session SessionRepository,
	im IdentityManager,
	tokenMaker token.Maker,
	accessTokenLifetime, refreshTokenLifetime time.Duration,
) *UseCase {
	return &UseCase{
		s,
		tr,
		hasher,
		session,
		im,
		tokenMaker,
		accessTokenLifetime,
		refreshTokenLifetime,
	}
}

func (u *UseCase) Register(ctx context.Context, req dto.UserRegisterRequest) (entity.User, error) {
	passHash, err := u.hasher.Hash(string(req.Password))
	if err != nil {
		return entity.User{}, fmt.Errorf("u.hasher.Hash - %w", err)
	}

	user := entity.User{
		Phone:        req.Phone,
		Login:        req.Login,
		Roles:        req.Roles,
		PasswordHash: passHash,
	}

	fn := func(txCtx context.Context) error {
		if user, err = u.s.Save(txCtx, user); err != nil {
			return fmt.Errorf("u.s.Save - %w", err)
		}
		return nil
	}

	if err = u.tr.ExecContext(ctx, fn); err != nil {
		return entity.User{}, fmt.Errorf("tr.Exec: %w", err)
	}

	// todo: send email to verify mail
	// todo: remove
	//userID := user.ID.String()
	//
	// keycloak
	//userKeycloak := gocloak.User{
	//	ID:      &userID,
	//	Enabled: gocloak.BoolP(true),
	//}
	//
	//_, err = u.im.CreateUser(ctx, userKeycloak, string(req.Password), "user")
	//if err != nil {
	//	return entity.User{}, fmt.Errorf("unable to create keycloak user: %w", err)
	//}

	return user, nil
}

func (u *UseCase) Get(ctx context.Context, in dto.UserGetRequest) (entity.User, error) {
	user, err := u.s.Get(ctx, entity.GetRequest{
		ID:        in.ID,
		Name:      in.Name,
		Login:     in.Login,
		Phone:     in.Phone,
		Email:     in.Email,
		AccountID: in.AccountID,
	})
	if err != nil {
		return entity.User{}, fmt.Errorf("get user via service: %w", err)
	}

	return user, nil
}

func (u *UseCase) Login(ctx context.Context, in dto.UserLoginRequest) (*dto.UserLoginResponse, error) {
	user, err := u.s.Get(ctx, entity.GetRequest{
		Login: in.Login,
	})
	if err != nil {
		return nil, fmt.Errorf("get user by login via service: %w", err)
	}

	// check password
	if !u.hasher.CheckHash(in.Password, user.PasswordHash) {
		return nil, errs.IncorrectPassword
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

	// todo: add UserAgent and ClientIp to session
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

func (u *UseCase) Delete(ctx context.Context, in dto.UserDeleteRequest) error {
	err := u.tr.ExecContext(ctx, func(txCtx context.Context) error {
		err := u.s.Delete(txCtx, in.ID)
		if err != nil {
			return fmt.Errorf("u.s.Delete - %w", err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("tr.Exec: %w", err)
	}
	return nil
}

func (u *UseCase) ChangePassword(ctx context.Context, in dto.UserChangePasswordRequest) error {
	user, err := u.s.Get(ctx, entity.GetRequest{
		ID: in.ID,
	})
	if err != nil {
		return fmt.Errorf("get user: %w", err)
	}

	match := u.hasher.CheckHash(string(in.OldPassword), user.PasswordHash)
	if !match {
		return errors.New("wrong old password")
	}

	passwordHash, err := u.hasher.Hash(string(in.NewPassword))
	if err != nil {
		return fmt.Errorf("u.hasher.Hash - %w", err)
	}

	if err = u.tr.ExecContext(ctx, func(txCtx context.Context) error {
		_, err = u.s.Update(txCtx, dto.UserUpdateRequest{
			ID:           in.ID,
			PasswordHash: passwordHash,
		})
		if err != nil {
			return fmt.Errorf("u.s.UpdatePasswordHash - %w", err)
		}
		return nil
	}); err != nil {
		return fmt.Errorf("tr.Exec: %w", err)
	}
	return nil
}

func (u *UseCase) Update(ctx context.Context, req dto.UserUpdateRequest) (entity.User, error) {
	var (
		user entity.User
		err  error
	)
	if err = u.tr.ExecContext(ctx, func(txCtx context.Context) error {
		user, err = u.s.Update(txCtx, req)
		if err != nil {
			return fmt.Errorf("u.s.Update - %w", err)
		}
		return nil
	}); err != nil {
		return entity.User{}, fmt.Errorf("tr.Exec: %w", err)
	}
	return user, nil
}

func (u *UseCase) RefreshTokens(ctx context.Context, in dto.UserRefreshTokensIn) (*dto.UserRefreshTokensOut, error) {
	payload, err := u.tokenMaker.VerifyToken(in.RefreshToken)
	if err != nil {
		return nil, fmt.Errorf("verify token: %w", err)
	}

	session, err := u.session.Get(ctx, nil, payload.ID)
	if err != nil {
		return nil, fmt.Errorf("get session: %w", err)
	}

	if session.IsBlocked {
		return nil, ErrSessionBlocked
	}

	if session.UserID != payload.UserID {
		return nil, ErrSessionUser
	}

	if session.RefreshToken != in.RefreshToken {
		return nil, ErrMismatchSessionToken
	}

	if time.Now().After(payload.ExpiredAt) {
		return nil, ErrExpiredSession
	}

	accessToken, accessPayload, err := u.tokenMaker.CreateToken(
		payload.UserID,
		u.accessTokenLifetime,
	)

	return &dto.UserRefreshTokensOut{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessPayload.ExpiredAt,
	}, nil
}
