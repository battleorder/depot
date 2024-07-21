drop schema if exists units cascade;
create schema units;

grant usage on schema units to anon, authenticated, service_role;
grant all on all tables in schema units to anon, authenticated, service_role;
grant all on all routines in schema units to anon, authenticated, service_role;
grant all on all sequences in schema units to anon, authenticated, service_role;
alter default privileges for role postgres in schema units grant all on tables to anon, authenticated, service_role;
alter default privileges for role postgres in schema units grant all on routines to anon, authenticated, service_role;
alter default privileges for role postgres in schema units grant all on sequences to anon, authenticated, service_role;

create or replace function units.member_is_admin(unit_id uuid, user_id uuid)
returns boolean
as $$
  begin
    return exists (
      select 1 from units.members
      where
        unitid = unit_id
        and userid = user_id
        and m.is_admin = true
    );
  end;
$$
language plpgsql;

create or replace function units.user_is_owner(unit_id uuid, user_id uuid)
returns boolean
as $$
  begin
    return exists (
      select 1 from units.units
      where
        id = unit_id
        and ownerid = user_id
    );
  end;
$$
language plpgsql;

create or replace function units.can_admin(unit_id uuid, user_id uuid)
returns boolean
as $$
  begin
    return coalesce(units.user_is_owner(unit_id, user_id), units.member_is_admin(unit_id, user_id));
  end;
$$
language plpgsql;
