DROP TYPE IF EXISTS operation_type;
CREATE TYPE operation_type AS ENUM (
    'refill',
    'withdraw',
    'hold',
    'clear',
    'commission'
);

DROP TYPE IF EXISTS account_status;
CREATE TYPE account_status AS ENUM (
    'active',
    'blocked',
    'closed'
);

create table if not exists account(
    id uuid primary key default gen_random_uuid(),
    owner uuid not null,
    balance bigint not null check ( balance >= 0 ) default 0,
    hold_balance bigint not null check ( hold_balance >= 0 ) default 0,
    country varchar(6) not null,
    currency varchar(6) not null,
    created_at timestamp(6) not null default current_timestamp(6),
    updated_at timestamp(6) not null default current_timestamp(6),
    status account_status not null default 'active'

--     constraint fk_owner_id foreign key (owner) references "user"(id)
);

create table if not exists transaction(
    id uuid PRIMARY KEY default gen_random_uuid(),
    from_account_id uuid not null,
    to_account_id uuid not null,
    amount bigint not null,
    created_at timestamp(6) not null default current_timestamp(6),
    constraint fk_from_account_id foreign key (from_account_id) references account(id),
    constraint fk_to_account_id foreign key (to_account_id) references account(id)
);

create table if not exists operation(
    id uuid PRIMARY KEY default gen_random_uuid(),
    transaction_id uuid not null,
    account_id uuid not null,
    type operation_type not null,
    amount bigint not null,
    balance_before bigint not null,
    balance_after bigint not null,
    created_at timestamp(6) not null default current_timestamp(6),
    constraint fk_transaction_id foreign key (transaction_id) references transaction(id),
    constraint fk_account_id foreign key (account_id) references account(id)
);
