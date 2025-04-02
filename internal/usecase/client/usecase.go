package client

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	account "komek/internal/domain/account/entity"
	"komek/internal/domain/client/entity"
	country "komek/internal/domain/country/entity"
	currency "komek/internal/domain/currency/entity"
	fundraise "komek/internal/domain/fundraise/entity"
	"komek/internal/usecase"
)

type UseCase struct {
	client    ClientService
	fundraise FundraiseService
	account   usecase.AccountService
	tr        usecase.Transactional
}

func New(
	client ClientService,
	fundraise FundraiseService,
	account usecase.AccountService,
	tr usecase.Transactional,
) *UseCase {
	return &UseCase{
		client:    client,
		fundraise: fundraise,
		account:   account,
		tr:        tr,
	}
}

func (uc *UseCase) ListClients(ctx context.Context) (entity.Clients, error) {
	clients, err := uc.client.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("list clients via service: %w", err)
	}

	return clients, nil
}

func (uc *UseCase) GetClientByID(ctx context.Context, clientID uuid.UUID) (cl Client, err error) {
	client, err := uc.client.GetByID(ctx, clientID)
	if err != nil {
		return cl, fmt.Errorf("get client by id via service: %w", err)
	}

	acc, err := uc.account.GetByOwnerID(ctx, clientID)
	if err != nil {
		return cl, fmt.Errorf("get account by user_id via service: %w", err)
	}

	fundraises, err := uc.fundraise.GetByAccountID(ctx, acc.ID)
	if err != nil {
		return cl, fmt.Errorf("get fundraises by account_id via service: %w", err)
	}

	return *cl.Fill(client).
		WithFundraises(fundraises).
		WithAccount(acc), nil
}

func (uc *UseCase) CreateClient(ctx context.Context, in entity.CreateIn) (client Client, err error) {
	var clientEntity entity.Client
	if err = uc.tr.ExecContext(ctx, func(txCtx context.Context) error {
		if clientEntity, err = uc.client.Create(txCtx, in); err != nil {
			return fmt.Errorf("create client via service: %w", err)
		}

		client.Fill(clientEntity)

		// create account for client
		if client.Account, err = uc.account.Create(txCtx, account.CreateIn{
			Owner:    client.ID,
			Balance:  0,
			Country:  country.KAZ,
			Currency: currency.KZT,
		}); err != nil {
			return fmt.Errorf("create account via service: %w", err)
		}
		return nil
	}); err != nil {
		return client, fmt.Errorf("exec via transaction: %w", err)
	}

	return client, nil
}

func (uc *UseCase) CreateFundraise(ctx context.Context, in fundraise.CreateIn) (fund fundraise.Fundraise, err error) {
	if err = uc.tr.ExecContext(ctx, func(txCtx context.Context) error {
		if fund, err = uc.fundraise.Create(txCtx, in); err != nil {
			return fmt.Errorf("create fundraise via service: %w", err)
		}
		return nil
	}); err != nil {
		return fund, fmt.Errorf("exec via transaction: %w", err)
	}

	return fund, nil
}

func (uc *UseCase) CreateFundraiseType(ctx context.Context, name string) error {
	if err := uc.tr.ExecContext(ctx, func(txCtx context.Context) error {
		if err := uc.fundraise.CreateType(txCtx, name); err != nil {
			return fmt.Errorf("create fundraise type via service: %w", err)
		}
		return nil
	}); err != nil {
		return fmt.Errorf("exec via transaction: %w", err)
	}

	return nil
}

func (uc *UseCase) GetFundraiseByID(ctx context.Context, id uuid.UUID) (fundraise.Fundraise, error) {
	fund, err := uc.fundraise.GetByID(ctx, id)
	if err != nil {
		return fundraise.Fundraise{}, fmt.Errorf("get fundraise_by_id via service: %w", err)
	}

	return fund, nil
}

func (uc *UseCase) ListActiveFundraises(ctx context.Context) ([]fundraise.Fundraise, error) {
	funds, err := uc.fundraise.ListActive(ctx)
	if err != nil {
		return nil, fmt.Errorf("list active fundraises via service: %w", err)
	}

	return funds, nil
}
