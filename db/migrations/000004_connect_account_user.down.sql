alter table if exists "account" drop constraint if exists "owner_currency_key";

alter table if exists "account" drop constraint if exists "fk_owner_id";