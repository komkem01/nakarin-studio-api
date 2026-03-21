SET statement_timeout = 0;

--bun:split

drop index if exists idx_bookings_delivery_zipcode_id;
drop index if exists idx_bookings_delivery_sub_district_id;
drop index if exists idx_bookings_delivery_district_id;
drop index if exists idx_bookings_delivery_province_id;
drop index if exists idx_bookings_delivery_member_address_id;

alter table bookings
    drop column if exists delivery_note,
    drop column if exists delivery_zipcode_id,
    drop column if exists delivery_sub_district_id,
    drop column if exists delivery_district_id,
    drop column if exists delivery_province_id,
    drop column if exists delivery_street,
    drop column if exists delivery_village,
    drop column if exists delivery_no,
    drop column if exists delivery_phone,
    drop column if exists delivery_last_name,
    drop column if exists delivery_first_name,
    drop column if exists delivery_member_address_id,
    drop column if exists last_tracking_at,
    drop column if exists tracking_attempt_count,
    drop column if exists internal_note,
    drop column if exists cancelled_reason;
