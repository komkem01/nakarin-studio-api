SET statement_timeout = 0;

--bun:split

create table zipcodes (
    id uuid primary key default gen_random_uuid(),
    sub_district_id uuid not null references sub_districts(id),
    name varchar(10) not null,
    is_active boolean not null default true,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    deleted_at timestamptz
);