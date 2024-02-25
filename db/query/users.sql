-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY created_at DESC;

-- name: SaveUser :one
INSERT INTO users(
    name, login, email, password_hash, phone, roles
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET name = coalesce(sqlc.narg(name), name),
    login = coalesce(sqlc.narg(login), login),
    email = coalesce(sqlc.narg(email), email),
    email_verified = coalesce(sqlc.narg(email_verified), email_verified),
    password_hash = coalesce(sqlc.narg(password_hash), password_hash),
    phone = coalesce(sqlc.narg(phone), phone),
    roles = coalesce(sqlc.narg(roles), roles),
    password_changed_at = coalesce(sqlc.narg(password_changed_at), password_changed_at),
    updated_at = CURRENT_TIMESTAMP(6)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: RemoveUser :one
DELETE FROM users
WHERE id = $1
RETURNING id;

-- name: FindUsers :many
SELECT *
FROM users
WHERE (@name::varchar = '' OR name = @name)
AND (@login::varchar = '' OR login = @login)
AND (@email::varchar = '' OR email = @email)
-- AND email_verified = $4
-- AND phone = $5
ORDER BY created_at DESC;