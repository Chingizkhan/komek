begin;

CREATE TYPE category_type AS (
    id UUID,
    name TEXT
);

create table if not exists "clients"(
    id uuid primary key default gen_random_uuid(),
    name text,
    phone text,
    email text,
    age numeric,
    city text,
    address text,
    description text,
    circumstances text,
    image_url text,
    created_at timestamp(6) not null default current_timestamp,
    updated_at timestamp(6) not null default current_timestamp
);

create table if not exists "categories"(
    id uuid primary key default gen_random_uuid(),
    name text,
    constraint unique_name unique (name)
);

create table if not exists "client_categories"(
    client_id uuid,
    category_id uuid,
    constraint fk_client_id foreign key (client_id) references clients(id),
    constraint fk_category_id foreign key (category_id) references categories(id)
);

create table if not exists "fundraises"(
    id uuid primary key default gen_random_uuid(),
    goal numeric not null,
    collected numeric not null default 0,
    account_id uuid not null,
    is_active bool not null default true,
    constraint fk_account_id foreign key (account_id) references account(id)
);

end;


-- sbory deneg (goal, collected, helpers from transfers)

-- goal
-- collected: number;
-- helpers: string[];
