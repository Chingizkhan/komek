begin;

insert into client_category(id, name)
values
    ('1e15f5c2-31e9-4967-b47a-0650bb9b8f62', 'Животные'),
    ('daa9a412-e5df-4d25-9290-7e708a17bd93', 'Пожилые')
on conflict do nothing;

insert into fundraise_types (id, name)
values
    ('e4c61f9f-d249-459d-aea0-4b2206960fe8', 'Ежемесячный сбор')
on conflict do nothing;

end;