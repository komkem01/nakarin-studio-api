SET statement_timeout = 0;

--bun:split

create table booking_details (
    id uuid primary key default gen_random_uuid(),
    booking_id uuid not null references bookings(id),
    first_name varchar(100) not null,
    last_name varchar(100),
    phone varchar(20) not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    deleted_at timestamptz
);