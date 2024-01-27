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

-- name: UpdateUser :one
UPDATE users
SET name = $2,
    login = $3,
    email = $4,
    email_verified = $5,
    phone = $6,
    password_hash = $7,
    updated_at = CURRENT_TIMESTAMP(6)
WHERE id = $1
RETURNING id;

-- name: RemoveUser :one
DELETE FROM users
WHERE id = $1
RETURNING id;