SET statement_timeout = 0;

--bun:split

create table sub_districts (
    id uuid primary key default gen_random_uuid(),
    district_id uuid not null references districts(id),
    name varchar(100) not null,
    is_active boolean not null default true,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    deleted_at timestamptz,
    unique (district_id, name)
);