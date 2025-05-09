// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: transaction.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createTransaction = `-- name: CreateTransaction :one

INSERT INTO transaction(
    from_account_id, to_account_id, amount
) VALUES (
    $1, $2, $3
) RETURNING id, from_account_id, to_account_id, amount, created_at
`

type CreateTransactionParams struct {
	FromAccountID pgtype.UUID `json:"from_account_id"`
	ToAccountID   pgtype.UUID `json:"to_account_id"`
	Amount        int64       `json:"amount"`
}

// -- name: CreateTransfer :one
// INSERT INTO transfers(
//
//	from_account_id,
//	to_account_id,
//	amount
//
// ) VALUES (
//
//	$1, $2, $3
//
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
func (q *Queries) CreateTransaction(ctx context.Context, arg CreateTransactionParams) (Transaction, error) {
	row := q.db.QueryRow(ctx, createTransaction, arg.FromAccountID, arg.ToAccountID, arg.Amount)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const getDonationsTotalAmountByAccountID = `-- name: GetDonationsTotalAmountByAccountID :one
SELECT sum(amount)
FROM transaction
where from_account_id = $1
`

func (q *Queries) GetDonationsTotalAmountByAccountID(ctx context.Context, fromAccountID pgtype.UUID) (int64, error) {
	row := q.db.QueryRow(ctx, getDonationsTotalAmountByAccountID, fromAccountID)
	var sum int64
	err := row.Scan(&sum)
	return sum, err
}

const getTransactionByAccountID = `-- name: GetTransactionByAccountID :many
SELECT id, from_account_id, to_account_id, amount, created_at
FROM transaction
WHERE from_account_id = $1 OR
      to_account_id = $1
ORDER BY created_at DESC
`

func (q *Queries) GetTransactionByAccountID(ctx context.Context, accountID pgtype.UUID) ([]Transaction, error) {
	rows, err := q.db.Query(ctx, getTransactionByAccountID, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Transaction{}
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.ID,
			&i.FromAccountID,
			&i.ToAccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTransactionByID = `-- name: GetTransactionByID :one
SELECT id, from_account_id, to_account_id, amount, created_at
FROM transaction
WHERE id = $1
`

func (q *Queries) GetTransactionByID(ctx context.Context, id pgtype.UUID) (Transaction, error) {
	row := q.db.QueryRow(ctx, getTransactionByID, id)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const getTransactionsByAccounts = `-- name: GetTransactionsByAccounts :many
SELECT id, from_account_id, to_account_id, amount, created_at
FROM transaction
WHERE from_account_id = $1 AND
      to_account_id = $2
`

type GetTransactionsByAccountsParams struct {
	FromAccountID pgtype.UUID `json:"from_account_id"`
	ToAccountID   pgtype.UUID `json:"to_account_id"`
}

func (q *Queries) GetTransactionsByAccounts(ctx context.Context, arg GetTransactionsByAccountsParams) ([]Transaction, error) {
	rows, err := q.db.Query(ctx, getTransactionsByAccounts, arg.FromAccountID, arg.ToAccountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Transaction{}
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.ID,
			&i.FromAccountID,
			&i.ToAccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
