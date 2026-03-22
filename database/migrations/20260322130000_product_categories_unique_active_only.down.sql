SET statement_timeout = 0;

--bun:split

drop index if exists uq_product_categories_name_active_only;

alter table product_categories
    add constraint product_categories_name_key unique (name);
