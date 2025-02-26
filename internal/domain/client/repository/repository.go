package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"komek/db/sqlc"
	"komek/internal/domain/client/entity"
	"komek/internal/errs"
	"komek/internal/service/transactional"
	"komek/pkg/null_value"
	"komek/pkg/postgres"
)

type Repository struct {
	pool *pgxpool.Pool
	q    *sqlc.Queries
}

func New(pg *postgres.Postgres) *Repository {
	return &Repository{pg.Pool, sqlc.New(pg.Pool)}
}

func (r *Repository) List(ctx context.Context) (entity.Clients, error) {
	qtx := r.queries(ctx)

	clients, err := qtx.ListClients(ctx)
	if err != nil {
		return entity.Clients{}, fmt.Errorf("r.q.List: %w", err)
	}
	convertedClients, err := r.listClientsToDomain(clients)
	if err != nil {
		return entity.Clients{}, err
	}
	return convertedClients, nil
}

func (r *Repository) GetByID(ctx context.Context, id uuid.UUID) (entity.Client, error) {
	qtx := r.queries(ctx)

	client, err := qtx.GetClientByID(ctx, null_value.UUID(id))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Client{}, errs.ErrClientNotFound
		}
		return entity.Client{}, fmt.Errorf("r.q.GetByID: %w", err)
	}
	convertedClient, err := r.getClientToDomain(client)
	if err != nil {
		return entity.Client{}, fmt.Errorf("r.getClientToDomain: %w", err)
	}
	return convertedClient, nil
}

func (r *Repository) Save(ctx context.Context, client entity.Client) (entity.Client, error) {
	qtx := r.queries(ctx)

	model, err := qtx.SaveClient(ctx, sqlc.SaveClientParams{
		Name:          null_value.String(client.Name),
		Phone:         null_value.String(string(client.Phone)),
		Email:         null_value.String(string(client.Email)),
		Age:           null_value.Number(&client.Age),
		City:          null_value.String(client.City),
		Address:       null_value.String(client.Address),
		Description:   null_value.String(client.Description),
		Circumstances: null_value.String(client.Circumstances),
		ImageUrl:      null_value.String(client.ImageURL),
	})
	if err != nil {
		return entity.Client{}, fmt.Errorf("r.q.Save: %w", err)
	}

	return r.clientToDomain(model), nil
}

func (r *Repository) SaveCategories(ctx context.Context, categories entity.Categories) (entity.Categories, error) {
	qtx := r.queries(ctx)

	models, err := qtx.SaveClientCategories(ctx, categories.Names())
	if err != nil {
		return nil, fmt.Errorf("r.q.SaveCategories: %w", err)
	}

	res := make(entity.Categories, len(models))
	for i, model := range models {
		res[i] = r.categoriesToDomain(model)
	}

	return res, nil
}

func (r *Repository) BindCategories(ctx context.Context, in entity.BindCategories) error {
	qtx := r.queries(ctx)

	categoryIDs := make([]pgtype.UUID, len(in.CategoryIDs))
	for i, id := range in.CategoryIDs {
		categoryIDs[i] = null_value.UUID(id)
	}

	if err := qtx.BindClientCategories(ctx, sqlc.BindClientCategoriesParams{
		ClientID:    null_value.UUID(in.ClientID),
		CategoryIds: categoryIDs,
	}); err != nil {
		return fmt.Errorf("r.q.BindClientCategories: %w", err)
	}

	return nil
}

func (r *Repository) queries(ctx context.Context) *sqlc.Queries {
	tx, ok := ctx.Value(transactional.TxKey).(pgx.Tx)
	if ok {
		return r.q.WithTx(tx)
	}
	return r.q
}
