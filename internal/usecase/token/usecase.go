package token

import (
	"context"
	"komek/internal/domain/jwt"
	"komek/internal/service/identity_manager"
)

type (
	IdentityManager interface {
		RetrospectToken(ctx context.Context, accessToken string) (*gocloak.IntroSpectTokenResult, error)
		RefreshTokens(ctx context.Context, refreshToken, userID string) (*identity_manager.UserWithJWTResponse, error)
	}

	UseCase struct {
		identityManager IdentityManager
	}
)

func New(im IdentityManager) *UseCase {
	return &UseCase{im}
}

func (uc *UseCase) Retrospect(ctx context.Context, accessToken string) (*gocloak.IntroSpectTokenResult, error) {
	token, err := uc.identityManager.RetrospectToken(ctx, accessToken)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (uc *UseCase) RefreshTokens(ctx context.Context, refreshToken string) (*identity_manager.UserWithJWTResponse, error) {
	claims := jwt.GetClaims(ctx)
	tokens, err := uc.identityManager.RefreshTokens(ctx, refreshToken, claims.Id)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}
