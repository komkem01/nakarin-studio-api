SET statement_timeout = 0;

--bun:split

alter table products
    add column if not exists suitable_for text,
    add column if not exists on_site text;
