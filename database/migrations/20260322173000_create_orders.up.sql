SET statement_timeout = 0;

--bun:split

create table orders (
    id uuid primary key default gen_random_uuid(),
    booking_id uuid not null references bookings(id),
    order_no text not null unique,
    status text not null default 'new',
    total_amount numeric(12,2) not null default 0,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    deleted_at timestamptz,
    unique (booking_id)
);
