SET statement_timeout = 0;

--bun:split

alter table products
    drop column if exists suitable_for,
    drop column if exists on_site;
