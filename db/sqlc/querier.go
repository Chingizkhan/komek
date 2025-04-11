// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	AddAccountBalance(ctx context.Context, arg AddAccountBalanceParams) (Account, error)
	BindClientCategories(ctx context.Context, arg BindClientCategoriesParams) error
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	CreateFundraise(ctx context.Context, arg CreateFundraiseParams) (Fundraise, error)
	CreateFundraiseType(ctx context.Context, name string) (FundraiseType, error)
	CreateOperation(ctx context.Context, arg CreateOperationParams) (Operation, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	// -- name: CreateTransfer :one
	// INSERT INTO transfers(
	//     from_account_id,
	//     to_account_id,
	//     amount
	// ) VALUES (
	//     $1, $2, $3
	// )
	// RETURNING *;
	//
	// -- name: GetTransfer :one
	// SELECT *
	// FROM transfers
	// WHERE id = $1
	// LIMIT 1;
	//
	// -- name: ListTransfers :many
	// SELECT *
	// FROM transfers
	// ORDER BY id
	// LIMIT $1
	// OFFSET $2;
	CreateTransaction(ctx context.Context, arg CreateTransactionParams) (Transaction, error)
	DeleteAccount(ctx context.Context, id pgtype.UUID) error
	DeleteOrganisation(ctx context.Context, id pgtype.UUID) error
	DonateFundraise(ctx context.Context, arg DonateFundraiseParams) error
	// AND email_verified = $4
	// AND phone = $5
	FindUsers(ctx context.Context, arg FindUsersParams) ([]User, error)
	GetAccount(ctx context.Context, id pgtype.UUID) (Account, error)
	GetAccountByOwnerID(ctx context.Context, ownerID pgtype.UUID) (Account, error)
	GetAccountForUpdate(ctx context.Context, id pgtype.UUID) (Account, error)
	GetClientByID(ctx context.Context, id pgtype.UUID) (GetClientByIDRow, error)
	GetFundraiseByAccountID(ctx context.Context, accountID pgtype.UUID) (Fundraise, error)
	GetFundraiseByID(ctx context.Context, id pgtype.UUID) (Fundraise, error)
	GetFundraisesByAccountID(ctx context.Context, accountID pgtype.UUID) ([]Fundraise, error)
	// -- name: CreateEntry :one
	// INSERT INTO entries(
	//     account_id,
	//     amount
	// ) VALUES (
	//     $1, $2
	// )
	// RETURNING *;
	//
	// -- name: GetEntry :one
	// SELECT *
	// FROM entries
	// WHERE id = $1
	// LIMIT 1;
	//
	// -- name: ListEntries :many
	// SELECT *
	// FROM entries
	// ORDER BY id
	// LIMIT $1
	// OFFSET $2;
	GetOperationsByTransactionID(ctx context.Context, transactionID pgtype.UUID) ([]Operation, error)
	GetOrganisation(ctx context.Context, id pgtype.UUID) (Organisation, error)
	GetSession(ctx context.Context, id pgtype.UUID) (Session, error)
	GetTransactionByAccountID(ctx context.Context, accountID pgtype.UUID) ([]Transaction, error)
	GetTransactionByID(ctx context.Context, id pgtype.UUID) (Transaction, error)
	GetTransactionsByAccounts(ctx context.Context, arg GetTransactionsByAccountsParams) ([]Transaction, error)
	GetUserByAccount(ctx context.Context, id pgtype.UUID) (User, error)
	GetUserByEmail(ctx context.Context, email pgtype.Text) (User, error)
	GetUserByID(ctx context.Context, id pgtype.UUID) (User, error)
	GetUserByLogin(ctx context.Context, login pgtype.Text) (User, error)
	GetUserByPhone(ctx context.Context, phone pgtype.Text) (User, error)
	ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error)
	ListActiveFundraises(ctx context.Context) ([]Fundraise, error)
	ListClients(ctx context.Context) ([]ListClientsRow, error)
	ListFinishedFundraises(ctx context.Context) ([]Fundraise, error)
	ListOrganisation(ctx context.Context) ([]Organisation, error)
	ListUsers(ctx context.Context) ([]User, error)
	RemoveUser(ctx context.Context, id pgtype.UUID) (pgtype.UUID, error)
	SaveClient(ctx context.Context, arg SaveClientParams) (Client, error)
	SaveClientCategories(ctx context.Context, names []string) ([]ClientCategory, error)
	SaveOrganisation(ctx context.Context, arg SaveOrganisationParams) error
	SaveUser(ctx context.Context, arg SaveUserParams) (User, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error)
	UpdateOrganisationBin(ctx context.Context, arg UpdateOrganisationBinParams) error
	UpdateOrganisationName(ctx context.Context, arg UpdateOrganisationNameParams) error
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
