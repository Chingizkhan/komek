CREATE TABLE IF NOT EXISTS users(
    id uuid PRIMARY KEY default gen_random_uuid(),
    name varchar(30),
    login varchar(20) NOT NULL,
    email varchar(20),
    email_verified bool DEFAULT FALSE,
    password_hash varchar(100) NOT NULL,
    phone varchar(15),
    created_at timestamp(6) default current_timestamp(6),
    updated_at timestamp(6) default current_timestamp(6)
);