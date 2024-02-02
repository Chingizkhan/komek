package oauthServer

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"komek/pkg/token"
	"net/http"
	"strings"
	"time"
)

const (
	introspectPath = "/oauth/token/introspect"
)

type OauthServer struct {
	client *http.Client
	path   string
}

func New(timeout time.Duration, path string) *OauthServer {
	return &OauthServer{
		client: &http.Client{Timeout: timeout},
		path:   path,
	}
}

func (s *OauthServer) Introspect(tok string) (IntrospectResponse, error) {
	req, err := http.NewRequest(http.MethodPost, s.path+introspectPath, nil)
	if err != nil {
		fmt.Printf("Error newRequest %s", err)
		return IntrospectResponse{}, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("%s %s", token.TypeBearer, tok))
	resp, err := s.client.Do(req)
	if err != nil {
		fmt.Printf("Error Do %s", err)
		return IntrospectResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if strings.Contains(string(body), "error") {
		var res ErrorResponse
		if err := getRes(body, &res); err != nil {
			return IntrospectResponse{}, errors.Wrap(err, "ErrorResponse")
		}
		return IntrospectResponse{}, errors.New(res.Error)
	}

	var response IntrospectResponse
	if err := getRes(body, &response); err != nil {
		return IntrospectResponse{}, errors.Wrap(err, "IntrospectResponse")
	}

	return IntrospectResponse{
		Active:   response.Active,
		ClientID: response.ClientID,
		Exp:      response.Exp,
		Sub:      response.Sub,
		UserName: response.UserName,
		TokenUse: response.TokenUse,
	}, nil
}

func getRes(body []byte, res any) error {
	err := json.Unmarshal(body, res)
	if err != nil {
		return errors.Wrap(err, "unmarshal JSON")
	}
	return nil
}
