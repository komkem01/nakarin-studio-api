SET statement_timeout = 0;

--bun:split

create table booking_status_logs (
    id uuid primary key default gen_random_uuid(),
    booking_id uuid not null references bookings(id),
    from_status booking_status,
    to_status booking_status not null,
    changed_by uuid,
    changed_by_role member_role,
    reason text,
    changed_at timestamptz not null default now()
);

create index idx_booking_status_logs_booking_id on booking_status_logs(booking_id);
create index idx_booking_status_logs_changed_at on booking_status_logs(changed_at);
