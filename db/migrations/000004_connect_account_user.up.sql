-- alter table "account" add constraint "fk_owner_id" foreign key ("owner") references "user"("id");

-- alter table "account" add constraint "owner_currency_key" unique ("owner", "currency");
