-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUsers :many
SELECT * FROM users
ORDER BY created_at DESC;

-- name: SaveUser :exec
INSERT INTO users(
    name, login, email, password_hash, phone
) VALUES (
    $1, $2, $3, $4, $5
);

-- name: UpdateUserName :one
UPDATE users
SET name = $2,
    updated_at = CURRENT_TIMESTAMP(6)
WHERE id = $1
RETURNING id;

-- name: UpdateUserLogin :one
UPDATE users
SET login = $2,
    updated_at = CURRENT_TIMESTAMP(6)
WHERE id = $1
RETURNING id;

-- name: UpdateUserEmail :one
UPDATE users
SET email = $2,
    updated_at = CURRENT_TIMESTAMP(6)
WHERE id = $1
RETURNING id;

-- name: UpdateUserEmailVerified :one
UPDATE users
SET email_verified = $2,
    updated_at = CURRENT_TIMESTAMP(6)
WHERE id = $1
RETURNING id;

-- name: UpdateUserPhone :one
UPDATE users
SET phone = $2,
    updated_at = CURRENT_TIMESTAMP(6)
WHERE id = $1
RETURNING id;

-- name: UpdateUserPasswordHash :one
UPDATE users
SET password_hash = $2,
    updated_at = CURRENT_TIMESTAMP(6)
WHERE id = $1
RETURNING id;

-- name: RemoveUser :one
DELETE FROM users
WHERE id = $1
RETURNING id;