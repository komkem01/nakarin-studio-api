SET statement_timeout = 0;

--bun:split

create table member_booking (
    id uuid primary key default gen_random_uuid(),
    member_id uuid not null references members(id),
    booking_id uuid not null references bookings(id),
    created_at timestamptz not null default now(),
    unique (member_id, booking_id)
);