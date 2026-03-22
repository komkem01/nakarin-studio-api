SET statement_timeout = 0;

--bun:split

alter table products
    add column if not exists category_id uuid;

alter table products
    add constraint fk_products_category_id
    foreign key (category_id)
    references product_categories(id)
    on update cascade
    on delete set null;

create index if not exists idx_products_category_id on products (category_id);
