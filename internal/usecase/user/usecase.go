package user

import (
	"context"
	"errors"
	"fmt"
	"komek/internal/domain"
	account "komek/internal/domain/account/entity"
	country "komek/internal/domain/country/entity"
	currency "komek/internal/domain/currency/entity"
	"komek/internal/domain/user/entity"
	"komek/internal/errs"
	"komek/internal/service/token"
	"komek/internal/usecase"
	"komek/pkg/money"
	"time"
)

type UseCase struct {
	s                    Service
	account              usecase.AccountService
	tr                   usecase.Transactional
	hasher               Hasher
	session              SessionRepository
	im                   IdentityManager
	tokenMaker           token.Maker
	accessTokenLifetime  time.Duration
	refreshTokenLifetime time.Duration
}

func New(
	s Service,
	account usecase.AccountService,
	tr usecase.Transactional,
	hasher Hasher,
	session SessionRepository,
	im IdentityManager,
	tokenMaker token.Maker,
	accessTokenLifetime, refreshTokenLifetime time.Duration,
) *UseCase {
	return &UseCase{
		s,
		account,
		tr,
		hasher,
		session,
		im,
		tokenMaker,
		accessTokenLifetime,
		refreshTokenLifetime,
	}
}

func (u *UseCase) userResponse(user entity.User, acc account.Account) entity.UserResponse {
	return entity.UserResponse{
		ID:                user.ID,
		Name:              user.Name,
		Phone:             user.Phone,
		Login:             user.Login,
		Email:             user.Email,
		EmailVerified:     user.EmailVerified,
		Roles:             user.Roles,
		CreatedAt:         user.CreatedAt.Unix(),
		UpdatedAt:         user.UpdatedAt.Unix(),
		PasswordChangedAt: user.PasswordChangedAt.Unix(),
		Account: entity.AccountResponse{
			ID:       user.ID,
			Balance:  money.ToFloat(acc.Balance),
			Currency: acc.Currency,
			Country:  acc.Country,
			Status:   acc.Status,
		},
	}
}

func (u *UseCase) Register(ctx context.Context, req entity.RegisterIn) (resp entity.UserResponse, err error) {
	var (
		user entity.User
		acc  account.Account
	)

	passHash, err := u.hasher.Hash(string(req.Password))
	if err != nil {
		return resp, fmt.Errorf("u.hasher.Hash - %w", err)
	}
	req.PasswordHash = passHash

	fn := func(txCtx context.Context) error {
		// create user
		if user, err = u.s.Register(txCtx, req); err != nil {
			return fmt.Errorf("u.s.Save - %w", err)
		}
		// create account
		if acc, err = u.account.Create(txCtx, account.CreateIn{
			Owner:    user.ID,
			Balance:  0,
			Country:  country.KAZ,
			Currency: currency.KZT,
		}); err != nil {
			return fmt.Errorf("u.account.Create - %w", err)
		}
		return nil
	}

	if err = u.tr.ExecContext(ctx, fn); err != nil {
		return resp, fmt.Errorf("tr.Exec: %w", err)
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

	return u.userResponse(user, acc), nil
}

func (u *UseCase) Get(ctx context.Context, in entity.GetIn) (entity.GetOut, error) {
	user, err := u.s.Get(ctx, in)
	if err != nil {
		return entity.GetOut{}, fmt.Errorf("get user via service: %w", err)
	}

	acc, err := u.account.GetByOwnerID(ctx, user.ID)
	if err != nil {
		return entity.GetOut{}, fmt.Errorf("get account via service: %w", err)
	}

	return entity.GetOut{
		User:    user,
		Account: acc,
	}, nil
}

func (u *UseCase) Login(ctx context.Context, in entity.LoginIn) (*entity.LoginOut, error) {
	user, err := u.s.Get(ctx, entity.GetIn{
		Login: in.Login,
		Phone: in.Phone,
		Email: in.Email,
	})
	if err != nil {
		return nil, fmt.Errorf("get user via service: %w", err)
	}

	// check password
	if !u.hasher.CheckHash(in.Password, user.PasswordHash) {
		return nil, errs.IncorrectPassword
	}

	acc, err := u.account.GetByOwnerID(ctx, user.ID)
	if err != nil {
		return nil, fmt.Errorf("get account via service: %w", err)
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

	return &entity.LoginOut{
		SessionID:             createdSession.ID,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
		User:                  u.userResponse(user, acc),
	}, nil
}

func (u *UseCase) Logout(ctx context.Context) error {
	return nil
}

func (u *UseCase) Delete(ctx context.Context, in entity.DeleteIn) error {
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

func (u *UseCase) ChangePassword(ctx context.Context, in entity.ChangePasswordIn) error {
	user, err := u.s.Get(ctx, entity.GetIn{
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
		_, err = u.s.Update(txCtx, entity.UpdateIn{
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

func (u *UseCase) Update(ctx context.Context, req entity.UpdateIn) (entity.User, error) {
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

func (u *UseCase) RefreshTokens(ctx context.Context, in entity.RefreshTokensIn) (*entity.RefreshTokensOut, error) {
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

	refreshToken, refreshPayload, err := u.tokenMaker.CreateToken(
		payload.UserID,
		u.refreshTokenLifetime,
	)

	return &entity.RefreshTokensOut{
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
	}, nil
}
