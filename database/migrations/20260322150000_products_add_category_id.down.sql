SET statement_timeout = 0;

--bun:split

drop index if exists idx_products_category_id;

alter table products
    drop constraint if exists fk_products_category_id;

alter table products
    drop column if exists category_id;
