-- name: GetWord :one
SELECT * FROM words
WHERE value = $1 AND
  fk_user_id = $2
LIMIT 1;

-- name: ListWords :many
SELECT * FROM words
WHERE fk_user_id = $1
ORDER BY created_at DESC;

-- name: SaveWord :exec
INSERT INTO words (
    fk_user_id, value, language, translation, created_at, updated_at
) VALUES (
     $1, $2, $3, $4, CURRENT_TIMESTAMP(6), CURRENT_TIMESTAMP(6)
);

-- name: UpdateWord :one
UPDATE words
SET value = $2,
    language = $3,
    translation = $4,
    updated_at = CURRENT_TIMESTAMP(6)
WHERE value = $1 AND
    fk_user_id = $2
RETURNING value;

-- name: DeleteWord :one
DELETE FROM words
WHERE value = $1 AND
    fk_user_id = $2
RETURNING value;