-- name: CreateDonation :exec
INSERT INTO donation(fundraise_id, transaction_id, client_id)
VALUES ($1, $2, $3);

-- name: GetDonationByID :one
SELECT
    d.*,
    tr.from_account_id,
    tr.to_account_id,
    tr.amount,
    cl.name as client_name,
    cl.image_url as client_image
FROM
    donation d
JOIN
    transaction tr on tr.id = d.transaction_id
JOIN
    clients cl on cl.id = d.client_id
WHERE d.id = $1;

-- name: GetDonationByTransactionID :one
SELECT
    d.*,
    tr.from_account_id,
    tr.to_account_id,
    tr.amount,
    cl.name as client_name,
    cl.image_url as client_image
FROM
    donation d
        JOIN
    transaction tr on tr.id = d.transaction_id
        JOIN
    clients cl on cl.id = d.client_id
WHERE d.transaction_id = $1;