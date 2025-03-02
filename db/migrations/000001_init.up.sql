CREATE TABLE IF NOT EXISTS "user"(
    id uuid PRIMARY KEY default gen_random_uuid(),
    name varchar(30),
    login varchar(30) UNIQUE,
    email varchar(30) UNIQUE,
    email_verified bool DEFAULT FALSE,
    password_hash varchar(100) NOT NULL,
    phone varchar(15) UNIQUE,
    roles text NOT NULL,
    password_changed_at timestamp(6) default null,
    created_at timestamp(6) NOT NULL default current_timestamp(6),
    updated_at timestamp(6) NOT NULL default current_timestamp(6)
);