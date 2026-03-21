SET statement_timeout = 0;

--bun:split

create table member_addresses (
    id uuid primary key default gen_random_uuid(),
    member_id uuid not null references members(id),
    first_name varchar(100) not null,
    last_name varchar(100),
    phone varchar(20) not null,
    no varchar(100),
    village varchar(255),
    street varchar(255),
    province_id uuid not null references provinces(id),
    district_id uuid not null references districts(id),
    sub_district_id uuid not null references sub_districts(id),
    zipcode_id uuid not null references zipcodes(id),
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    deleted_at timestamptz
);