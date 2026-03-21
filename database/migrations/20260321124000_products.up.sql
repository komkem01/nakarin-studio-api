SET statement_timeout = 0;

--bun:split

create table products (
    id uuid primary key default gen_random_uuid(),
    name varchar(255) not null,
    description text,
    price numeric(12,2) not null default 0,
    is_active boolean not null default true,
    is_available boolean not null default true,
    sort_order int not null default 0,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    deleted_at timestamptz
);
