package httpclient

import (
	"io"
	"net/http"
)

var (
	Client *HttpClient
)

type (
	HttpClient struct {
		c    *http.Client
		addr string
	}
)

func init() {
	Client = &HttpClient{
		c:    &http.Client{},
		addr: "http://localhost:8081",
	}
}

func (c *HttpClient) Exec(request *Request) (*Response, error) {
	response, err := c.c.Do(request.Request)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return &Response{response, body}, nil
}
