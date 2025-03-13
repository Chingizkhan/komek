package client

import (
	"github.com/google/uuid"
	account "komek/internal/domain/account/entity"
	"komek/internal/domain/client/entity"
	"komek/internal/domain/email"
	fundraise "komek/internal/domain/fundraise/entity"
	"komek/internal/domain/phone"
	"time"
)

type (
	Client struct {
		ID            uuid.UUID             `json:"id"`
		Name          string                `json:"name"`
		Phone         phone.Phone           `json:"phone"`
		Email         email.Email           `json:"email"`
		Age           int                   `json:"age"`
		City          string                `json:"city"`
		Address       string                `json:"address"`
		Description   string                `json:"description"`
		Circumstances string                `json:"circumstances"`
		ImageURL      string                `json:"image_url"`
		Account       account.Account       `json:"account"`
		Categories    entity.Categories     `json:"categories"`
		Fundraises    []fundraise.Fundraise `json:"fundraises"`
		CreatedAt     time.Time             `json:"created_at"`
		UpdatedAt     time.Time             `json:"updated_at"`
	}
)

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
	c.Fundraises = fundraises
	return c
}

func (c *Client) WithAccount(acc account.Account) *Client {
	c.Account = acc
	return c
}
