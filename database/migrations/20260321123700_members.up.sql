SET statement_timeout = 0;

--bun:split

create table members (
    id uuid primary key default gen_random_uuid(),
    gender_id uuid not null references genders(id),
    prefix_id uuid references prefixes(id),
    role member_role not null default 'customer',
    first_name varchar(100) not null,
    last_name varchar(100),
    phone varchar(20) not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    deleted_at timestamptz
);