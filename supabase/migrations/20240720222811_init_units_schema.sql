drop schema if exists units cascade;
create schema units;

grant usage on schema units to anon, authenticated, service_role;
grant all on all tables in schema units to anon, authenticated, service_role;
grant all on all routines in schema units to anon, authenticated, service_role;
grant all on all sequences in schema units to anon, authenticated, service_role;
alter default privileges for role postgres in schema units grant all on tables to anon, authenticated, service_role;
alter default privileges for role postgres in schema units grant all on routines to anon, authenticated, service_role;
alter default privileges for role postgres in schema units grant all on sequences to anon, authenticated, service_role;
