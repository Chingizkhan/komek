package httpclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type (
	Response struct {
		R    *http.Response
		Body []byte
	}

	ErrorResponse struct {
		Error string `json:"error"`
	}
)

func (r *Response) Error() (err error) {
	if err != nil {
		return err
	}
	isErr := strings.Contains(strings.ToLower(string(r.Body)), "error")
	if isErr {
		return r.extract()
	}
	return nil
}

func (r *Response) extract() error {
	var res ErrorResponse
	err := json.Unmarshal(r.Body, &res)
	if err != nil {
		fmt.Printf("Error unmarshal %s", err)
		return err
	}
	return errors.New(res.Error)
}

func (r *Response) Close() {
	_ = r.R.Body.Close()
}
