-- name: GetFundraiseByID :one
select *
from fundraises
where id = $1
limit 1;

-- name: ListActiveFundraises :many
select *
from fundraises
where is_active = true;

-- name: ListFinishedFundraises :many
select *
from fundraises
where is_active = false;

-- name: GetFundraisesByAccountID :many
select *
from fundraises
where account_id = $1;

-- name: CreateFundraise :one
insert into fundraises(
    goal, collected, account_id, is_active
) values (
    $1, $2, $3, sqlc.narg('is_active')
) returning *;