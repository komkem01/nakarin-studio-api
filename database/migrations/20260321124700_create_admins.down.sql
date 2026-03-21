SET statement_timeout = 0;

--bun:split

drop index if exists idx_admins_is_active;
drop index if exists idx_admins_member_id;

drop table if exists admins;
