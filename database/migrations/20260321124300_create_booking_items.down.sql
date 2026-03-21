SET statement_timeout = 0;

--bun:split

drop index if exists idx_booking_items_product_id;
drop index if exists idx_booking_items_booking_id;

drop table if exists booking_items;
