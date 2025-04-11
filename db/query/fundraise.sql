-- name: GetFundraiseByID :one
select *
from fundraises
where id = $1
limit 1;

-- name: GetFundraiseByAccountID :one
select *
from fundraises
where account_id = $1
order by created_at
limit 1;

-- name: ListActiveFundraises :many
select *
from fundraises
where is_active = true
order by created_at;

-- name: GetFundraisesByAccountID :many
select *
from fundraises
where account_id = $1
and is_active
order by created_at;

-- name: ListFinishedFundraises :many
select *
from fundraises
where is_active = false
order by created_at;

-- name: CreateFundraiseType :one
insert into fundraise_types(
    name
) values (
    $1
) returning *;

-- name: CreateFundraise :one
insert into fundraises(
    goal, collected, type, account_id, is_active
) values (
    $1, $2, $3, $4, sqlc.narg('is_active')
) returning *;

-- name: DonateFundraise :exec
update fundraises
set
    collected = collected +sqlc.arg(amount)
where id = sqlc.arg(id);

-- name: SetFundraiseStatus :exec
update fundraises
set
    is_active = $2
where id = $1;