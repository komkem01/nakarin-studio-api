SET statement_timeout = 0;

--bun:split

drop index if exists idx_payments_status;
drop index if exists idx_payments_booking_id;

drop table if exists payments;

drop type if exists payment_status;
drop type if exists payment_channel;
