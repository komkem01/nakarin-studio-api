SET statement_timeout = 0;

--bun:split

alter table product_categories
    drop constraint if exists product_categories_name_key;

create unique index if not exists uq_product_categories_name_active_only
    on product_categories (name)
    where deleted_at is null;
