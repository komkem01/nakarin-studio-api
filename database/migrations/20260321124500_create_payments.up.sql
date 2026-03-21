SET statement_timeout = 0;

--bun:split

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'payment_channel') THEN
        CREATE TYPE payment_channel AS ENUM ('bank_transfer', 'promptpay', 'cash', 'credit_card', 'other');
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'payment_status') THEN
        CREATE TYPE payment_status AS ENUM ('pending', 'paid', 'failed', 'refunded');
    END IF;
END
$$;

create table payments (
    id uuid primary key default gen_random_uuid(),
    booking_id uuid not null references bookings(id),
    channel payment_channel not null default 'bank_transfer',
    amount numeric(12,2) not null default 0,
    deposit_amount numeric(12,2) not null default 0,
    status payment_status not null default 'pending',
    proof_url text,
    note text,
    paid_at timestamptz,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    deleted_at timestamptz,
    constraint payments_amount_non_negative check (amount >= 0),
    constraint payments_deposit_amount_non_negative check (deposit_amount >= 0)
);

create index idx_payments_booking_id on payments(booking_id);
create index idx_payments_status on payments(status);
