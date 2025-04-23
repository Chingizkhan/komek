begin;

create table if not exists "donation"(
    id uuid primary key default gen_random_uuid(),
    fundraise_id uuid not null,
    transaction_id uuid not null,
    client_id uuid not null,
    created_at timestamp(6) not null default current_timestamp(6),

    constraint fk_fundraise_id foreign key (fundraise_id) references fundraises(id),
    constraint fk_transaction_id foreign key (transaction_id) references transaction(id),
    constraint fk_client_id foreign key (client_id) references clients(id)
);

end;