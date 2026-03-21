SET statement_timeout = 0;

--bun:split

create table admins (
    id uuid primary key default gen_random_uuid(),
    member_id uuid unique references members(id),
    username varchar(100) not null unique,
    password_hash text not null,
    display_name varchar(255),
    last_login_at timestamptz,
    is_active boolean not null default true,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    deleted_at timestamptz
);

create index idx_admins_member_id on admins(member_id);
create index idx_admins_is_active on admins(is_active);
