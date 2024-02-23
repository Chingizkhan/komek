// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package sqlc

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	AddAccountBalance(ctx context.Context, arg AddAccountBalanceParams) (Account, error)
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error)
	CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error)
	DeleteAccount(ctx context.Context, id int64) error
	DeleteOrganisation(ctx context.Context, id uuid.UUID) error
	// AND email_verified = $4
	// AND phone = $5
	FindUsers(ctx context.Context, arg FindUsersParams) ([]User, error)
	GetAccount(ctx context.Context, id int64) (Account, error)
	GetAccountForUpdate(ctx context.Context, id int64) (Account, error)
	GetEntry(ctx context.Context, id int64) (Entry, error)
	GetOrganisation(ctx context.Context, id uuid.UUID) (Organisation, error)
	GetTransfer(ctx context.Context, id int64) (Transfer, error)
	GetUser(ctx context.Context, id uuid.UUID) (User, error)
	GetUsers(ctx context.Context) ([]User, error)
	ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error)
	ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error)
	ListOrganisation(ctx context.Context) ([]Organisation, error)
	ListTransfers(ctx context.Context, arg ListTransfersParams) ([]Transfer, error)
	RemoveUser(ctx context.Context, id uuid.UUID) (uuid.UUID, error)
	SaveOrganisation(ctx context.Context, arg SaveOrganisationParams) error
	SaveUser(ctx context.Context, arg SaveUserParams) (User, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error)
	UpdateOrganisationBin(ctx context.Context, arg UpdateOrganisationBinParams) error
	UpdateOrganisationName(ctx context.Context, arg UpdateOrganisationNameParams) error
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
