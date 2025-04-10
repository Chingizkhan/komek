-- name: GetFundraiseByID :one
select *
from fundraises
where id = $1
limit 1;

-- name: GetFundraiseByAccountID :one
select *
from fundraises
where account_id = $1
limit 1;

-- name: ListActiveFundraises :many
select *
from fundraises
where is_active = true;

-- name: GetFundraisesByAccountID :many
select *
from fundraises
where account_id = $1
and is_active;

-- name: ListFinishedFundraises :many
select *
from fundraises
where is_active = false;

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