SET statement_timeout = 0;

--bun:split

create table product_categories (
    id uuid primary key default gen_random_uuid(),
    name varchar(255) not null unique,
    description text,
    is_active boolean not null default true,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    deleted_at timestamptz
);
