SET statement_timeout = 0;

--bun:split

create table booking_items (
    id uuid primary key default gen_random_uuid(),
    booking_id uuid not null references bookings(id),
    product_id uuid not null references products(id),
    product_name varchar(255) not null,
    unit_price_at_booking numeric(12,2) not null default 0,
    quantity int not null default 1,
    line_total numeric(12,2) not null default 0,
    note text,
    sort_order int not null default 0,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    deleted_at timestamptz,
    constraint booking_items_quantity_positive check (quantity > 0),
    constraint booking_items_unit_price_non_negative check (unit_price_at_booking >= 0),
    constraint booking_items_line_total_non_negative check (line_total >= 0)
);

create index idx_booking_items_booking_id on booking_items(booking_id);
create index idx_booking_items_product_id on booking_items(product_id);
