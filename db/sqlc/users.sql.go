// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: users.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const findUsers = `-- name: FindUsers :many
SELECT id, name, login, email, email_verified, password_hash, phone, roles, password_changed_at, created_at, updated_at
FROM users
WHERE ($1::varchar = '' OR name = $1)
AND ($2::varchar = '' OR login = $2)
AND ($3::varchar = '' OR email = $3)
ORDER BY created_at DESC
`

type FindUsersParams struct {
	Name  string `json:"name"`
	Login string `json:"login"`
	Email string `json:"email"`
}

// AND email_verified = $4
// AND phone = $5
func (q *Queries) FindUsers(ctx context.Context, arg FindUsersParams) ([]User, error) {
	rows, err := q.db.Query(ctx, findUsers, arg.Name, arg.Login, arg.Email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Login,
			&i.Email,
			&i.EmailVerified,
			&i.PasswordHash,
			&i.Phone,
			&i.Roles,
			&i.PasswordChangedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const getUserByAccount = `-- name: GetUserByAccount :one
SELECT id, name, login, email, email_verified, password_hash, phone, roles, password_changed_at, created_at, updated_at FROM users as u
WHERE u.id = (
    SELECT a.owner FROM accounts as a
    WHERE a.id = $1 LIMIT 1
)
`

func (q *Queries) GetUserByAccount(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRow(ctx, getUserByAccount, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Login,
		&i.Email,
		&i.EmailVerified,
		&i.PasswordHash,
		&i.Phone,
		&i.Roles,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, name, login, email, email_verified, password_hash, phone, roles, password_changed_at, created_at, updated_at FROM users
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email pgtype.Text) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Login,
		&i.Email,
		&i.EmailVerified,
		&i.PasswordHash,
		&i.Phone,
		&i.Roles,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, name, login, email, email_verified, password_hash, phone, roles, password_changed_at, created_at, updated_at FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUserByID(ctx context.Context, id pgtype.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Login,
		&i.Email,
		&i.EmailVerified,
		&i.PasswordHash,
		&i.Phone,
		&i.Roles,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByLogin = `-- name: GetUserByLogin :one
SELECT id, name, login, email, email_verified, password_hash, phone, roles, password_changed_at, created_at, updated_at FROM users
WHERE "login" = $1 LIMIT 1
`

func (q *Queries) GetUserByLogin(ctx context.Context, login pgtype.Text) (User, error) {
	row := q.db.QueryRow(ctx, getUserByLogin, login)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Login,
		&i.Email,
		&i.EmailVerified,
		&i.PasswordHash,
		&i.Phone,
		&i.Roles,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByPhone = `-- name: GetUserByPhone :one
SELECT id, name, login, email, email_verified, password_hash, phone, roles, password_changed_at, created_at, updated_at FROM users
WHERE phone = $1 LIMIT 1
`

func (q *Queries) GetUserByPhone(ctx context.Context, phone pgtype.Text) (User, error) {
	row := q.db.QueryRow(ctx, getUserByPhone, phone)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Login,
		&i.Email,
		&i.EmailVerified,
		&i.PasswordHash,
		&i.Phone,
		&i.Roles,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, name, login, email, email_verified, password_hash, phone, roles, password_changed_at, created_at, updated_at FROM users
ORDER BY created_at DESC
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Login,
			&i.Email,
			&i.EmailVerified,
			&i.PasswordHash,
			&i.Phone,
			&i.Roles,
			&i.PasswordChangedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const removeUser = `-- name: RemoveUser :one
DELETE FROM users
WHERE id = $1
RETURNING id
`

func (q *Queries) RemoveUser(ctx context.Context, id pgtype.UUID) (pgtype.UUID, error) {
	row := q.db.QueryRow(ctx, removeUser, id)
	err := row.Scan(&id)
	return id, err
}

const saveUser = `-- name: SaveUser :one
INSERT INTO users(
    name, login, email, password_hash, phone, roles
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING id, name, login, email, email_verified, password_hash, phone, roles, password_changed_at, created_at, updated_at
`

type SaveUserParams struct {
	Name         pgtype.Text `json:"name"`
	Login        pgtype.Text `json:"login"`
	Email        pgtype.Text `json:"email"`
	PasswordHash string      `json:"password_hash"`
	Phone        pgtype.Text `json:"phone"`
	Roles        string      `json:"roles"`
}

func (q *Queries) SaveUser(ctx context.Context, arg SaveUserParams) (User, error) {
	row := q.db.QueryRow(ctx, saveUser,
		arg.Name,
		arg.Login,
		arg.Email,
		arg.PasswordHash,
		arg.Phone,
		arg.Roles,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Login,
		&i.Email,
		&i.EmailVerified,
		&i.PasswordHash,
		&i.Phone,
		&i.Roles,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET name = coalesce($1, name),
    login = coalesce($2, login),
    email = coalesce($3, email),
    email_verified = coalesce($4, email_verified),
    password_hash = coalesce($5, password_hash),
    phone = coalesce($6, phone),
    roles = coalesce($7, roles),
    password_changed_at = coalesce($8, password_changed_at),
    updated_at = CURRENT_TIMESTAMP(6)
WHERE id = $9
RETURNING id, name, login, email, email_verified, password_hash, phone, roles, password_changed_at, created_at, updated_at
`

type UpdateUserParams struct {
	Name              pgtype.Text      `json:"name"`
	Login             pgtype.Text      `json:"login"`
	Email             pgtype.Text      `json:"email"`
	EmailVerified     pgtype.Bool      `json:"email_verified"`
	PasswordHash      pgtype.Text      `json:"password_hash"`
	Phone             pgtype.Text      `json:"phone"`
	Roles             pgtype.Text      `json:"roles"`
	PasswordChangedAt pgtype.Timestamp `json:"password_changed_at"`
	ID                pgtype.UUID      `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.Name,
		arg.Login,
		arg.Email,
		arg.EmailVerified,
		arg.PasswordHash,
		arg.Phone,
		arg.Roles,
		arg.PasswordChangedAt,
		arg.ID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Login,
		&i.Email,
		&i.EmailVerified,
		&i.PasswordHash,
		&i.Phone,
		&i.Roles,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
