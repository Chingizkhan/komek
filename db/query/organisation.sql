-- name: GetOrganisation :one
SELECT * FROM organisation
WHERE id = $1
LIMIT 1;

-- name: ListOrganisation :many
SELECT * FROM organisation
ORDER BY created_at DESC;

-- name: SaveOrganisation :exec
INSERT INTO organisation (
    name, bin
) VALUES (
     $1, $2
);

-- name: UpdateOrganisationName :exec
UPDATE organisation
SET name = $2,
    updated_at = CURRENT_TIMESTAMP(6)
WHERE id = $1;

-- name: UpdateOrganisationBin :exec
UPDATE organisation
SET bin = $2,
    updated_at = CURRENT_TIMESTAMP(6)
WHERE id = $1;

-- name: DeleteOrganisation :exec
DELETE FROM organisation
WHERE id = $1;