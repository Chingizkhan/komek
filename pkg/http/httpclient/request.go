package httpclient

import (
	"context"
	"fmt"
	"io"
	"komek/pkg/token"
	"net/http"
)

type Request struct {
	*http.Request
}

func NewRequest(ctx context.Context, method, url string, body io.Reader) (*Request, error) {
	request, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Accept", "application/json")

	return &Request{request}, nil
}

func (r *Request) WithAuthorization(tok string) {
	r.Header.Add("Authorization", fmt.Sprintf("%s %s", token.TypeBearer, tok))
}
