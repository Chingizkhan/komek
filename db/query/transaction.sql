-- -- name: CreateTransfer :one
-- INSERT INTO transfers(
--     from_account_id,
--     to_account_id,
--     amount
-- ) VALUES (
--     $1, $2, $3
-- )
-- RETURNING *;
--
-- -- name: GetTransfer :one
-- SELECT *
-- FROM transfers
-- WHERE id = $1
-- LIMIT 1;
--
-- -- name: ListTransfers :many
-- SELECT *
-- FROM transfers
-- ORDER BY id
-- LIMIT $1
-- OFFSET $2;

-- name: CreateTransaction :one
INSERT INTO transaction(
    from_account_id, to_account_id, amount
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetTransactionByID :one
SELECT *
FROM transaction
WHERE id = $1;

-- name: GetTransactionByAccountID :many
SELECT *
FROM transaction
WHERE from_account_id = @account_id OR
      to_account_id = @account_id
ORDER BY created_at DESC;

-- name: GetTransactionsByAccounts :many
SELECT *
FROM transaction
WHERE from_account_id = $1 AND
      to_account_id = $2;

-- name: GetDonationsTotalAmountByAccountID :one
SELECT sum(amount)
FROM transaction
where from_account_id = $1;
