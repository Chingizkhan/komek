package identity_manager

import (
	"context"
	"fmt"
	"github.com/Nerzal/gocloak/v13"
	"github.com/pkg/errors"
	"strings"
)

type IdentityManager struct {
	baseUrl             string
	realm               string
	restApiClientId     string
	restApiClientSecret string
}

func New(url, realm, clientId, clientSecret string) *IdentityManager {
	return &IdentityManager{
		baseUrl:             url,
		realm:               realm,
		restApiClientId:     clientId,
		restApiClientSecret: clientSecret,
	}
}

func (im *IdentityManager) loginRestApiClient(ctx context.Context) (*gocloak.JWT, error) {
	client := gocloak.NewClient(im.baseUrl)

	token, err := client.LoginClient(ctx, im.restApiClientId, im.restApiClientSecret, im.realm)
	if err != nil {
		return nil, errors.Wrap(err, "unable to login the goclack rest client")
	}

	return token, nil
}

func (im *IdentityManager) CreateUser(ctx context.Context, user gocloak.User, password, role string) (*gocloak.User, error) {
	token, err := im.loginRestApiClient(ctx)
	if err != nil {
		return nil, err
	}

	client := gocloak.NewClient(im.baseUrl)

	userId, err := client.CreateUser(ctx, token.AccessToken, im.realm, user)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create user")
	}

	err = client.SetPassword(ctx, token.AccessToken, userId, im.realm, password, false)
	if err != nil {
		return nil, errors.Wrap(err, "unable to set user's password")
	}

	var roleNameLowerCase = strings.ToLower(role)
	roleKeyCloak, err := client.GetRealmRole(ctx, token.AccessToken, im.realm, roleNameLowerCase)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("unable to get role by name: '%v'", role))
	}
	err = client.AddRealmRoleToUser(ctx, token.AccessToken, im.realm, userId, []gocloak.Role{
		*roleKeyCloak,
	})
	if err != nil {
		return nil, errors.Wrap(err, "unable to add realm role to user")
	}

	userKeycloak, err := client.GetUserByID(ctx, token.AccessToken, im.realm, userId)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get recently created user")
	}

	return userKeycloak, nil
}

func (im *IdentityManager) RetrospectToken(ctx context.Context, accessToken string) (*gocloak.IntroSpectTokenResult, error) {
	client := gocloak.NewClient(im.baseUrl)

	rtpResult, err := client.RetrospectToken(ctx, accessToken, im.restApiClientId, im.restApiClientSecret, im.realm)
	if err != nil {
		return nil, errors.Wrap(err, "unable to retrospect token")
	}
	return rtpResult, err
}

func (im *IdentityManager) RefreshTokens(ctx context.Context, refreshToken, userID string) (*UserWithJWTResponse, error) {
	token, err := im.loginRestApiClient(ctx)
	if err != nil {
		return nil, err
	}

	client := gocloak.NewClient(im.baseUrl)

	jwt, err := client.RefreshToken(ctx, refreshToken, im.restApiClientSecret, im.restApiClientSecret, im.realm)
	if err != nil {
		return nil, errors.Wrap(err, "unable to refresh token")
	}

	u, err := client.GetUserByID(ctx, token.AccessToken, im.realm, userID)
	if err != nil {
		return nil, err
	}

	return &UserWithJWTResponse{
		User: convertUser(u),
		JWT:  convertJWT(jwt),
	}, nil
}

func (im *IdentityManager) SignIn(ctx context.Context, request SignInRequest) (*UserWithJWTResponse, error) {
	token, err := im.loginRestApiClient(ctx)
	if err != nil {
		return nil, err
	}

	client := gocloak.NewClient(im.baseUrl)

	jwt, err := client.Login(ctx, im.restApiClientId, im.restApiClientSecret, im.realm, request.Username, request.Password)
	if err != nil {
		return nil, errors.Wrap(err, "unable to login client")
	}
	info, err := client.GetUsers(ctx, token.AccessToken, im.realm, gocloak.GetUsersParams{
		Username: &request.Username,
	})
	if err != nil {
		return nil, errors.Wrap(err, "unable to get user info")
	}

	if len(info) == 0 {
		return nil, ErrUserNotFound
	}

	u := info[0]

	return &UserWithJWTResponse{
		User: convertUser(u),
		JWT:  convertJWT(jwt),
	}, nil
}

func (im *IdentityManager) Logout(ctx context.Context, userID string) error {
	token, err := im.loginRestApiClient(ctx)
	if err != nil {
		return err
	}

	client := gocloak.NewClient(im.baseUrl)

	sessions, err := client.GetUserSessions(ctx, token.AccessToken, im.realm, "4b9db7ad-6aaf-4fbe-a499-bc8a88ec129f")
	if err != nil {
		return err
	}
	if err != nil {
		return errors.Wrap(err, "unable to get client user sessions")
	}

	if len(sessions) == 0 {
		return ErrNoSession
	}

	for _, s := range sessions {
		err = client.LogoutUserSession(ctx, token.AccessToken, im.realm, *s.ID)
		if err != nil {
			return errors.Wrap(err, "unable to logout client")
		}
	}

	return nil
}
