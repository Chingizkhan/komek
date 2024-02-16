create table if not exists "organisation"(
    id uuid primary key default gen_random_uuid(),
    name text not null,
    bin varchar(30) not null,
    created_at timestamp(6) default current_timestamp(6),
    updated_at timestamp(6) default current_timestamp(6)
)