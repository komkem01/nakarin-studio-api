SET statement_timeout = 0;

--bun:split

alter table products
    add column if not exists received_items text,
    add column if not exists note text;
