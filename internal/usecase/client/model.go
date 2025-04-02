package client

import (
	"github.com/google/uuid"
	account "komek/internal/domain/account/entity"
	"komek/internal/domain/client/entity"
	"komek/internal/domain/email"
	fundraise "komek/internal/domain/fundraise/entity"
	"komek/internal/domain/phone"
	"komek/pkg/money"
	"time"
)

type (
	Client struct {
		ID            uuid.UUID         `json:"id"`
		Name          string            `json:"name"`
		Phone         phone.Phone       `json:"phone"`
		Email         email.Email       `json:"email"`
		Age           int               `json:"age"`
		City          string            `json:"city"`
		Address       string            `json:"address"`
		Description   string            `json:"description"`
		Circumstances string            `json:"circumstances"`
		ImageURL      string            `json:"image_url"`
		Account       account.Account   `json:"account"`
		Categories    entity.Categories `json:"categories"`
		Fundraises    []Fundraise       `json:"fundraises"`
		CreatedAt     time.Time         `json:"created_at"`
		UpdatedAt     time.Time         `json:"updated_at"`
	}

	Fundraise struct {
		ID        uuid.UUID      `json:"id"`
		Goal      float64        `json:"goal"`
		Collected float64        `json:"collected"`
		Type      fundraise.Type `json:"type"`
		AccountID uuid.UUID      `json:"account_id"`
		IsActive  bool           `json:"is_active"`
	}
)

func (f *Fundraise) FromDomain(fundraise fundraise.Fundraise) Fundraise {
	*f = Fundraise{
		ID:        fundraise.ID,
		Goal:      money.ToFloat(fundraise.Goal),
		Collected: money.ToFloat(fundraise.Collected),
		Type:      fundraise.Type,
		AccountID: fundraise.AccountID,
		IsActive:  fundraise.IsActive,
	}
	return *f
}

func (c *Client) Fill(client entity.Client) *Client {
	*c = Client{
		ID:            client.ID,
		Name:          client.Name,
		Phone:         client.Phone,
		Email:         client.Email,
		Age:           client.Age,
		City:          client.City,
		Address:       client.Address,
		Description:   client.Description,
		Circumstances: client.Circumstances,
		ImageURL:      client.ImageURL,
		Categories:    client.Categories,
		CreatedAt:     client.CreatedAt,
		UpdatedAt:     client.UpdatedAt,
	}
	return c
}

func (c *Client) WithFundraises(fundraises []fundraise.Fundraise) *Client {
	funds := make([]Fundraise, 0, len(fundraises))
	for _, fundraise := range fundraises {
		funds = append(funds, new(Fundraise).FromDomain(fundraise))
	}
	c.Fundraises = funds
	return c
}

func (c *Client) WithAccount(acc account.Account) *Client {
	c.Account = acc
	return c
}
