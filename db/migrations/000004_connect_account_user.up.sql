alter table "accounts" add foreign key ("owner") references "users"("login");

alter table "accounts" add constraint "owner_currency_key" unique ("owner", "currency");
