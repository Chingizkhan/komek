package repository

import (
	"encoding/json"
	"fmt"
	"komek/db/sqlc"
	"komek/internal/domain/client/entity"
	"komek/internal/domain/email"
	"komek/internal/domain/phone"
)

func (r *Repository) listClientsToDomain(model []sqlc.ListClientsRow) (entity.Clients, error) {
	res := make(entity.Clients, 0, len(model))

	for _, m := range model {
		var categories []entity.Category
		if err := json.Unmarshal(m.Categories, &categories); err != nil {
			return nil, fmt.Errorf("unmarshal: %w", err)
		}
		res = append(res, entity.Client{
			ID:            m.ID.Bytes,
			Name:          m.Name.String,
			Phone:         phone.Phone(m.Phone.String),
			Email:         email.Email(m.Email.String),
			Age:           int(m.Age.Int.Int64()),
			City:          m.City.String,
			Address:       m.Address.String,
			Description:   m.Description.String,
			Circumstances: m.Circumstances.String,
			ImageURL:      m.ImageUrl.String,
			Categories:    categories,
			CreatedAt:     m.CreatedAt.Time,
			UpdatedAt:     m.UpdatedAt.Time,
		})
	}

	return res, nil
}

func (r *Repository) getClientToDomain(m sqlc.GetClientByIDRow) (entity.Client, error) {
	var categories []entity.Category
	if err := json.Unmarshal(m.Categories, &categories); err != nil {
		return entity.Client{}, fmt.Errorf("unmarshal: %w", err)
	}
	return entity.Client{
		ID:            m.ID.Bytes,
		Name:          m.Name.String,
		Phone:         phone.Phone(m.Phone.String),
		Email:         email.Email(m.Email.String),
		Age:           int(m.Age.Int.Int64()),
		City:          m.City.String,
		Address:       m.Address.String,
		Description:   m.Description.String,
		Circumstances: m.Circumstances.String,
		ImageURL:      m.ImageUrl.String,
		Categories:    categories,
		CreatedAt:     m.CreatedAt.Time,
		UpdatedAt:     m.UpdatedAt.Time,
	}, nil
}

func (r *Repository) clientToDomain(m sqlc.Client) entity.Client {
	return entity.Client{
		ID:            m.ID.Bytes,
		Name:          m.Name.String,
		Phone:         phone.Phone(m.Phone.String),
		Email:         email.Email(m.Email.String),
		Age:           int(m.Age.Int.Int64()),
		City:          m.City.String,
		Address:       m.Address.String,
		Description:   m.Description.String,
		Circumstances: m.Circumstances.String,
		ImageURL:      m.ImageUrl.String,
		Categories:    nil,
		CreatedAt:     m.CreatedAt.Time,
		UpdatedAt:     m.UpdatedAt.Time,
	}
}

func (r *Repository) categoriesToDomain(m sqlc.Category) entity.Category {
	return entity.Category{
		ID:   m.ID.Bytes,
		Name: m.Name.String,
	}
}
