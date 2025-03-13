-- name: GetClientByID :one
SELECT
    cl.*,
    jsonb_agg(jsonb_build_object(
        'id', ct.id,
        'name', ct.name
    )) as categories
FROM
    clients cl
JOIN
    client_category_map cl_ct on cl.id = cl_ct.client_id
JOIN
    client_category ct on cl_ct.category_id = ct.id
WHERE
    cl.id = $1
GROUP BY
    cl.id
LIMIT 1;

-- name: ListClients :many
SELECT
    cl.*,
    jsonb_agg(jsonb_build_object(
        'id', ct.id,
        'name', ct.name
    )) as categories
FROM
    clients cl
JOIN
    client_category_map cl_ct on cl.id = cl_ct.client_id
JOIN
    client_category ct on cl_ct.category_id = ct.id
GROUP BY
    cl.id
ORDER BY
    cl.created_at DESC;

-- name: SaveClient :one
INSERT INTO clients(
    name, phone, email, age, city, address, description, circumstances, image_url
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING *;

-- name: SaveClientCategories :many
INSERT INTO client_category(
    name
) VALUES (
    unnest(@names::text[])
) ON CONFLICT DO NOTHING
RETURNING *;

-- name: BindClientCategories :exec
INSERT INTO client_category_map(
    client_id, category_id
) VALUES (
    $1, unnest(@category_ids::uuid[])
);

-- -- name: UpdateUser :one
-- UPDATE users
-- SET name = coalesce(sqlc.narg(name), name),
--     login = coalesce(sqlc.narg(login), login),
--     email = coalesce(sqlc.narg(email), email),
--     email_verified = coalesce(sqlc.narg(email_verified), email_verified),
--     password_hash = coalesce(sqlc.narg(password_hash), password_hash),
--     phone = coalesce(sqlc.narg(phone), phone),
--     roles = coalesce(sqlc.narg(roles), roles),
--     password_changed_at = coalesce(sqlc.narg(password_changed_at), password_changed_at),
--     updated_at = CURRENT_TIMESTAMP(6)
-- WHERE id = sqlc.arg(id)
-- RETURNING *;
--
-- -- name: RemoveUser :one
-- DELETE FROM clients
-- WHERE id = $1
-- RETURNING id;
