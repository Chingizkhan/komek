package identity

import (
	"context"
	"fmt"
	"github.com/Nerzal/gocloak/v13"
	"strings"
)

type identityManager struct {
	baseURL      string
	realm        string
	clientID     string
	clientSecret string
}

func NewIdentityManager(baseURL, realm, clientID, clientSecret string) *identityManager {
	return &identityManager{
		baseURL:      baseURL,
		realm:        realm,
		clientID:     clientID,
		clientSecret: clientID,
	}
}

func (im *identityManager) loginRestApiClient(ctx context.Context) (*gocloak.JWT, error) {
	client := gocloak.NewClient(im.baseURL)

	token, err := client.LoginClient(ctx, im.clientID, im.clientSecret, im.realm)
	if err != nil {
		return nil, fmt.Errorf("unable to login the rest api client: %w", err)
	}
	return token, nil
}

func (im *identityManager) CreateUser(ctx context.Context, user gocloak.User, password, role string) (*gocloak.User, error) {
	token, err := im.loginRestApiClient(ctx)
	if err != nil {
		return nil, err
	}

	client := gocloak.NewClient(im.baseURL)

	userID, err := client.CreateUser(ctx, token.AccessToken, im.realm, user)
	if err != nil {
		return nil, fmt.Errorf("unable to create user: %w", err)
	}

	err = client.SetPassword(ctx, token.AccessToken, userID, im.realm, password, false)
	if err != nil {
		return nil, fmt.Errorf("unable to set password: %w", err)
	}

	var roleNameLowerCase = strings.ToLower(role)
	roleKeycloak, err := client.GetRealmRole(ctx, token.AccessToken, im.realm, roleNameLowerCase)
	if err != nil {
		return nil, fmt.Errorf("unable to get realm role '%s': %w", roleNameLowerCase, err)
	}
	err = client.AddRealmRoleToUser(ctx, token.AccessToken, im.realm, userID, []gocloak.Role{
		*roleKeycloak,
	})

	userKeycloak, err := client.GetUserByID(ctx, token.AccessToken, im.realm, userID)
	if err != nil {
		return nil, fmt.Errorf("unable to get user by id: %w", err)
	}
	
	return userKeycloak, nil
}
