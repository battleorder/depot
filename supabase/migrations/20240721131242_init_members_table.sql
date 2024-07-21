set search_path = units, public, extensions;

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

alter table units.members enable row level security;

create policy "Members are viewable by everyone"
  on units.members for select
  to authenticated, anon
  using ( true );

create policy "Members are creatable only by authenticated users"
  on units.members for insert
  to authenticated
  with check (
    exists (
      select 1 from units.units
      where
        units.units.id = units.members.unitid
        and units.units.ownerid = (select auth.uid())
    )
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
