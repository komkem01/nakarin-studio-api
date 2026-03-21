SET statement_timeout = 0;

--bun:split

alter table products
    drop column if exists prep_time;
