package user_managment

import (
	"context"
	u "komek/internal/domain/user"
	"komek/internal/service/identity_manager"
	"strings"
)

type (
	UseCase struct {
		identityManager IdentityManager
	}

	RegisterResponse struct {
		User *gocloak.User
	}

	RegisterRequest struct {
		Username     string
		Password     string
		FirstName    string
		LastName     string
		Email        string
		MobileNumber string
	}

	SignInRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	SignInResponse struct {
		User `json:"user"`
		JWT  `json:"jwt"`
	}

	User struct {
		ID            string `json:"id"`
		UserName      string `json:"user_name"`
		Enabled       bool   `json:"enabled"`
		Email         string `json:"email"`
		EmailVerified bool   `json:"email_verified"`
		FirstName     string `json:"first_name"`
		LastName      string `json:"last_name"`
		Phone         string `json:"phone"`
		CreatedAt     int64  `json:"created_at"`
	}

	JWT struct {
		AccessToken      string `json:"access_token"`
		RefreshToken     string `json:"refresh_token"`
		ExpiresIn        int    `json:"expires_in"`
		RefreshExpiresIn int    `json:"refresh_expires_in"`
		SessionState     string `json:"session_state"`
	}

	LogoutRequest struct {
		UserID string `json:"user_id"`
	}
)

func New(im IdentityManager) *UseCase {
	return &UseCase{im}
}

func (uc *UseCase) Register(ctx context.Context, request RegisterRequest) (*RegisterResponse, error) {
	user := gocloak.User{
		Username:      gocloak.StringP(request.Username),
		FirstName:     gocloak.StringP(request.FirstName),
		LastName:      gocloak.StringP(request.LastName),
		Email:         gocloak.StringP(request.Email),
		EmailVerified: gocloak.BoolP(true),
		Enabled:       gocloak.BoolP(true),
		Attributes:    &map[string][]string{},
	}
	if strings.TrimSpace(request.MobileNumber) != "" {
		(*user.Attributes)["mobile"] = []string{request.MobileNumber}
	}

	userResponse, err := uc.identityManager.CreateUser(ctx, user, request.Password, string(u.RoleViewer))
	if err != nil {
		return nil, err
	}
	return &RegisterResponse{userResponse}, nil
}

func (uc *UseCase) SignIn(ctx context.Context, request SignInRequest) (*SignInResponse, error) {
	info, err := uc.identityManager.SignIn(ctx, identity_manager.SignInRequest{
		Username: request.Username,
		Password: request.Password,
	})
	if err != nil {
		return nil, err
	}

	return &SignInResponse{
		User: User{
			ID:            info.User.ID,
			UserName:      info.User.UserName,
			Enabled:       info.User.Enabled,
			Email:         info.User.Email,
			EmailVerified: info.User.EmailVerified,
			FirstName:     info.User.FirstName,
			LastName:      info.User.LastName,
			Phone:         info.User.Phone,
			CreatedAt:     info.User.CreatedAt,
		},
		JWT: JWT{
			AccessToken:      info.JWT.AccessToken,
			RefreshToken:     info.JWT.RefreshToken,
			ExpiresIn:        info.JWT.ExpiresIn,
			RefreshExpiresIn: info.JWT.RefreshExpiresIn,
			SessionState:     info.JWT.SessionState,
		},
	}, nil
}

func (uc *UseCase) Logout(ctx context.Context, request LogoutRequest) error {
	err := uc.identityManager.Logout(ctx, request.UserID)
	if err != nil {
		return err
	}
	return nil
}
