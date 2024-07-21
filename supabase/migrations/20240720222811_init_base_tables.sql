create schema if not exists units;

grant usage on schema units to anon, authenticated, service_role;
grant all on all tables in schema units to anon, authenticated, service_role;
grant all on all routines in schema units to anon, authenticated, service_role;
grant all on all sequences in schema units to anon, authenticated, service_role;
alter default privileges for role postgres in schema units grant all on tables to anon, authenticated, service_role;
alter default privileges for role postgres in schema units grant all on routines to anon, authenticated, service_role;
alter default privileges for role postgres in schema units grant all on sequences to anon, authenticated, service_role;

create table units.units (
  -- ident
  id          uuid        primary key default uuid_generate_v4(),
  slug        varchar(32) unique not null,
  displayname varchar(64) unique not null,

  -- branding
  tagline     text,
  description text,
  avatar      text,

  -- metadata
  createdat timestamptz default now(),
  updatedat timestamptz default now()
);

create table units.ranks (
  -- ident
  id          uuid        primary key default uuid_generate_v4(),
  unitid      uuid        not null,
  slug        varchar(32) not null,
  displayname varchar(64) not null,

  -- branding
  rankorder int  not null default 0,
  avatar    text,

  -- metadata
  createdat timestamptz default now(),
  updatedat timestamptz default now(),

  -- constraints
  unique(unitid, slug),
  unique(unitid, displayname),
  foreign key (unitid) references units.units(id)
);

create table units.members (
  -- ident
  id          uuid       primary key default uuid_generate_v4(),
  unitid      uuid       not null,
  userid      uuid       not null,
  rankid      uuid       not null,
  displayname varchar(64),

  -- metadata
  isadmin   bool        default false,
  createdat timestamptz default now(),
  updatedat timestamptz default now(),

  -- constraints
  unique(unitid, userid),
  unique(unitid, displayname),
  foreign key (unitid) references units.units(id),
  foreign key (rankid) references units.ranks(id),
  foreign key (userid) references auth.users(id)
);

create view units.member_names as (
  select
    m.id as id,
    unitid,
    coalesce(m.displayname, u.email) as displayname
  from
    units.members m
    left join auth.users u on u.id = m.userid
);
