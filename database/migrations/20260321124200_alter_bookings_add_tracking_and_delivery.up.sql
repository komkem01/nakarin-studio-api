SET statement_timeout = 0;

--bun:split

alter table bookings
    add column cancelled_reason text,
    add column internal_note text,
    add column tracking_attempt_count int not null default 0,
    add column last_tracking_at timestamptz,
    add column delivery_member_address_id uuid references member_addresses(id),
    add column delivery_first_name varchar(100),
    add column delivery_last_name varchar(100),
    add column delivery_phone varchar(20),
    add column delivery_no varchar(50),
    add column delivery_village varchar(100),
    add column delivery_street varchar(255),
    add column delivery_province_id uuid references provinces(id),
    add column delivery_district_id uuid references districts(id),
    add column delivery_sub_district_id uuid references sub_districts(id),
    add column delivery_zipcode_id uuid references zipcodes(id),
    add column delivery_note text;

create index idx_bookings_delivery_member_address_id on bookings(delivery_member_address_id);
create index idx_bookings_delivery_province_id on bookings(delivery_province_id);
create index idx_bookings_delivery_district_id on bookings(delivery_district_id);
create index idx_bookings_delivery_sub_district_id on bookings(delivery_sub_district_id);
create index idx_bookings_delivery_zipcode_id on bookings(delivery_zipcode_id);
