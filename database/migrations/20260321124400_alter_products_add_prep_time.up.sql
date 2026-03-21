SET statement_timeout = 0;

--bun:split

alter table products
    add column prep_time int not null default 0;
