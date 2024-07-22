drop schema if exists units cascade;
create schema units;

grant usage on schema units to anon, authenticated, service_role;
grant all on all tables in schema units to anon, authenticated, service_role;
grant all on all routines in schema units to anon, authenticated, service_role;
grant all on all sequences in schema units to anon, authenticated, service_role;
alter default privileges for role postgres in schema units grant all on tables to anon, authenticated, service_role;
alter default privileges for role postgres in schema units grant all on routines to anon, authenticated, service_role;
alter default privileges for role postgres in schema units grant all on sequences to anon, authenticated, service_role;

create or replace function uuid6() returns uuid as $$
declare
begin
	return uuid6(clock_timestamp());
end $$ language plpgsql;

create or replace function uuid6(p_timestamp timestamp with time zone) returns uuid as $$
declare

	v_time double precision := null;

	v_gregorian_t bigint := null;
	v_clock_sequence_and_node bigint := null;

	v_gregorian_t_hex_a varchar := null;
	v_gregorian_t_hex_b varchar := null;
	v_clock_sequence_and_node_hex varchar := null;

	c_epoch double precision := 12219292800; -- RFC-9562 epoch: 1582-10-15
	c_100ns_factor double precision := 10^7; -- RFC-9562 precision: 100 ns

	c_version bigint := x'0000000000006000'::bigint; -- RFC-9562 version: b'0110...'
	c_variant bigint := x'8000000000000000'::bigint; -- RFC-9562 variant: b'10xx...'

begin

	v_time := extract(epoch from p_timestamp);

	v_gregorian_t := trunc((v_time + c_epoch) * c_100ns_factor);
	v_clock_sequence_and_node := trunc(random() * 2^30)::bigint << 32 | trunc(random() * 2^32)::bigint;

	v_gregorian_t_hex_a := lpad(to_hex((v_gregorian_t >> 12)), 12, '0');
	v_gregorian_t_hex_b := lpad(to_hex((v_gregorian_t & 4095) | c_version), 4, '0');
	v_clock_sequence_and_node_hex := lpad(to_hex(v_clock_sequence_and_node | c_variant), 16, '0');

	return (v_gregorian_t_hex_a || v_gregorian_t_hex_b  || v_clock_sequence_and_node_hex)::uuid;

end $$ language plpgsql;

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
