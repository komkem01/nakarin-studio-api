SET statement_timeout = 0;

--bun:split

create table product_images (
    id uuid primary key default gen_random_uuid(),
    product_id uuid not null references products(id),
    image_url text not null,
    alt_text varchar(255),
    sort_order int not null default 0,
    is_active boolean not null default true,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    deleted_at timestamptz
);
