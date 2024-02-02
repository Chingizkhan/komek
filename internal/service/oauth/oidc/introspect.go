package oidc

import (
	"context"
	"fmt"
	"komek/pkg/http/httpclient"
	"net/http"
)

const (
	// todo: maybe should move
	tokenIntrospectURL = "http://localhost:8081/oauth/token/introspect"
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
		return fmt.Errorf("resp.Error: %w", err)
	}

	return nil
}
