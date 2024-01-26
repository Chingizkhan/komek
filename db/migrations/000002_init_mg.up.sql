CREATE TABLE users IF NOT EXISTS (
                       id   uuid PRIMARY KEY,
                       name varchar(30) NOT NULL,
                       login varchar(20) NOT NULL,
                       email varchar(20) NOT NULL,
                       phone varchar(15) NOT NULL,
                       created_at timestamp(6) NOT NULL,
                       updated_at timestamp(6) NOT NULL
);

CREATE TABLE words IF NOT EXISTS (
                       fk_user_id uuid REFERENCES users(id),
                       value   text NOT NULL,
                       language varchar(5) NOT NULL,
                       translation text NOT NULL,
                       created_at timestamp(6) NOT NULL,
                       updated_at timestamp(6) NOT NULL
);
