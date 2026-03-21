SET statement_timeout = 0;

--bun:split

create table bookings (
    id uuid primary key default gen_random_uuid(),
    booking_no varchar(50) not null unique,
    status booking_status not null default 'pending',
    payment payment_type not null default 'deposit',
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    deleted_at timestamptz
);