package client_credentials

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"komek/pkg/http/httpclient"
	"net/http"
)

const (
	// todo: maybe should move
	tokenIntrospectURL = "http://localhost:8081/client-credentials/token/introspect"
	tokenAuthURL       = "http://localhost:8081/client-credentials/auth"
)

func Introspect(ctx context.Context, tok string) error {
	req, err := httpclient.NewRequest(ctx, http.MethodPost, tokenIntrospectURL, nil)
	if err != nil {
		return fmt.Errorf("httpclient.NewRequest: %w", err)
	}
	req.WithAuthorization(tok)
	resp, err := httpclient.Client.Exec(req)
	if err != nil {
		return fmt.Errorf("httpclient.Client.Exec: %w", err)
	}
	defer resp.Close()
	err = resp.Error()
	if err != nil {
		return err
	}

	return nil
}

func Auth(ctx context.Context, clientID, clientSecret string) (string, error) {
	body := struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
	}{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}
	js, err := json.Marshal(body)
	if err != nil {
		return "", err
	}
	req, err := httpclient.NewRequest(ctx, http.MethodPost, tokenAuthURL, bytes.NewBuffer(js))
	if err != nil {
		return "", fmt.Errorf("httpclient.NewRequest: %w", err)
	}
	resp, err := httpclient.Client.Exec(req)
	if err != nil {
		return "", fmt.Errorf("httpclient.Client.Exec: %w", err)
	}
	defer resp.Close()
	err = resp.Error()
	if err != nil {
		return "", err
	}
	accessToken := resp.R.Header.Get("Access-Token")

	return accessToken, nil
}
