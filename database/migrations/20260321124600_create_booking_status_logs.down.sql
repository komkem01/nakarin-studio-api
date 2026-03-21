SET statement_timeout = 0;

--bun:split

drop index if exists idx_booking_status_logs_changed_at;
drop index if exists idx_booking_status_logs_booking_id;

drop table if exists booking_status_logs;
