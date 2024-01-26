-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUsers :many
SELECT * FROM users
ORDER BY created_at DESC;

-- name: SaveUser :one
INSERT INTO users(
    id, name, login, email, phone, created_at, updated_at
) VALUES (
    $1, $2, $3, $4, $5, CURRENT_TIMESTAMP(6), CURRENT_TIMESTAMP(6)
);

-- name: UpdateUser :one
UPDATE users
SET id = $1,
    name = $2,
    login = $3,
    email = $4,
    phone = $5,
    updated_at = CURRENT_TIMESTAMP(6)
WHERE id = $1
RETURNING id;

-- name: RemoveUser :one
DELETE FROM users
WHERE id = $1
RETURNING id;