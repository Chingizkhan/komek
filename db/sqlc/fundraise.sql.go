// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: fundraise.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createFundraise = `-- name: CreateFundraise :one
insert into fundraises(
    goal, collected, type, account_id, is_active
) values (
    $1, $2, $3, $4, $5
) returning id, type, goal, collected, account_id, is_active
`

type CreateFundraiseParams struct {
	Goal      int64       `json:"goal"`
	Collected int64       `json:"collected"`
	Type      pgtype.UUID `json:"type"`
	AccountID pgtype.UUID `json:"account_id"`
	IsActive  pgtype.Bool `json:"is_active"`
}

func (q *Queries) CreateFundraise(ctx context.Context, arg CreateFundraiseParams) (Fundraise, error) {
	row := q.db.QueryRow(ctx, createFundraise,
		arg.Goal,
		arg.Collected,
		arg.Type,
		arg.AccountID,
		arg.IsActive,
	)
	var i Fundraise
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Goal,
		&i.Collected,
		&i.AccountID,
		&i.IsActive,
	)
	return i, err
}

const createFundraiseType = `-- name: CreateFundraiseType :one
insert into fundraise_types(
    name
) values (
    $1
) returning id, name
`

func (q *Queries) CreateFundraiseType(ctx context.Context, name string) (FundraiseType, error) {
	row := q.db.QueryRow(ctx, createFundraiseType, name)
	var i FundraiseType
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getFundraiseByAccountID = `-- name: GetFundraiseByAccountID :one
select id, type, goal, collected, account_id, is_active
from fundraises
where account_id = $1
limit 1
`

func (q *Queries) GetFundraiseByAccountID(ctx context.Context, accountID pgtype.UUID) (Fundraise, error) {
	row := q.db.QueryRow(ctx, getFundraiseByAccountID, accountID)
	var i Fundraise
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Goal,
		&i.Collected,
		&i.AccountID,
		&i.IsActive,
	)
	return i, err
}

const getFundraiseByID = `-- name: GetFundraiseByID :one
select id, type, goal, collected, account_id, is_active
from fundraises
where id = $1
limit 1
`

func (q *Queries) GetFundraiseByID(ctx context.Context, id pgtype.UUID) (Fundraise, error) {
	row := q.db.QueryRow(ctx, getFundraiseByID, id)
	var i Fundraise
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Goal,
		&i.Collected,
		&i.AccountID,
		&i.IsActive,
	)
	return i, err
}

const getFundraisesByAccountID = `-- name: GetFundraisesByAccountID :many
select id, type, goal, collected, account_id, is_active
from fundraises
where account_id = $1
and is_active
`

func (q *Queries) GetFundraisesByAccountID(ctx context.Context, accountID pgtype.UUID) ([]Fundraise, error) {
	rows, err := q.db.Query(ctx, getFundraisesByAccountID, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Fundraise{}
	for rows.Next() {
		var i Fundraise
		if err := rows.Scan(
			&i.ID,
			&i.Type,
			&i.Goal,
			&i.Collected,
			&i.AccountID,
			&i.IsActive,
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

const listActiveFundraises = `-- name: ListActiveFundraises :many
select id, type, goal, collected, account_id, is_active
from fundraises
where is_active = true
`

func (q *Queries) ListActiveFundraises(ctx context.Context) ([]Fundraise, error) {
	rows, err := q.db.Query(ctx, listActiveFundraises)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Fundraise{}
	for rows.Next() {
		var i Fundraise
		if err := rows.Scan(
			&i.ID,
			&i.Type,
			&i.Goal,
			&i.Collected,
			&i.AccountID,
			&i.IsActive,
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

const listFinishedFundraises = `-- name: ListFinishedFundraises :many
select id, type, goal, collected, account_id, is_active
from fundraises
where is_active = false
`

func (q *Queries) ListFinishedFundraises(ctx context.Context) ([]Fundraise, error) {
	rows, err := q.db.Query(ctx, listFinishedFundraises)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Fundraise{}
	for rows.Next() {
		var i Fundraise
		if err := rows.Scan(
			&i.ID,
			&i.Type,
			&i.Goal,
			&i.Collected,
			&i.AccountID,
			&i.IsActive,
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
