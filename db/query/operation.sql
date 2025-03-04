-- -- name: CreateEntry :one
-- INSERT INTO entries(
--     account_id,
--     amount
-- ) VALUES (
--     $1, $2
-- )
-- RETURNING *;
--
-- -- name: GetEntry :one
-- SELECT *
-- FROM entries
-- WHERE id = $1
-- LIMIT 1;
--
-- -- name: ListEntries :many
-- SELECT *
-- FROM entries
-- ORDER BY id
-- LIMIT $1
-- OFFSET $2;

-- name: GetOperationsByTransactionID :many
SELECT *
FROM operation
WHERE transaction_id = $1
ORDER BY created_at;

-- name: CreateOperation :one
INSERT INTO operation(
    transaction_id, account_id, type, amount, balance_before, balance_after
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;