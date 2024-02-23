package national_bank

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type (
	NationalBank struct {
		host string
		c    *http.Client
	}
)

const (
	allRanksUrl = "/rss/rates_all.xml"
)

func New(host string) *NationalBank {
	return &NationalBank{
		host,
		http.DefaultClient,
	}
}

func (nb *NationalBank) GetRates(ctx context.Context) (*Rss, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, nb.host+allRanksUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext: %w", err)
	}

	response, err := nb.c.Do(request)
	if err != nil {
		return nil, fmt.Errorf("Client.Do: %w", err)
	}
	defer func() {
		err = response.Body.Close()
		if err != nil {
			//logrus.Error("national_bank -> response.Body.Close:", err)
		}
	}()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll: %w", err)
	}

	rss := new(Rss)
	err = xml.Unmarshal(body, rss)
	if err != nil {
		return nil, fmt.Errorf("xml.Unmarshal: %w", err)
	}

	return rss, nil
}
