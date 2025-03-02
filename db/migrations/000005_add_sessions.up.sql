begin;

CREATE TABLE if not exists "session" (
    "id" uuid PRIMARY KEY,
    "user_id" uuid not null,
    "refresh_token" varchar not null,
    "user_agent" varchar not null,
    "client_ip" varchar not null,
    "is_blocked" bool not null default false,
    "expires_at" timestamptz not null,
    "created_at" timestamptz not null default (now())
);

alter table "session" add constraint "fk_user_id" foreign key ("user_id") references "user" ("id");

end;