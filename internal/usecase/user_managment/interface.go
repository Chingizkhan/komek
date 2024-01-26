package user_managment

import (
	"context"
	"komek/internal/service/identity_manager"
)

type (
	IdentityManager interface {
		CreateUser(ctx context.Context, user gocloak.User, password, role string) (*gocloak.User, error)
		SignIn(ctx context.Context, request identity_manager.SignInRequest) (*identity_manager.UserWithJWTResponse, error)
		Logout(ctx context.Context, userID string) error
	}
)
