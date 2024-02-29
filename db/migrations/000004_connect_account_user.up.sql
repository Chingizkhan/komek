alter table "accounts" add foreign key ("owner") references "users"("id");

alter table "accounts" add constraint "owner_currency_key" unique ("owner", "currency");
