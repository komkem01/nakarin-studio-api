SET statement_timeout = 0;

--bun:split

alter table products
    drop column if exists received_items,
    drop column if exists note;
