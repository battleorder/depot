set search_path = units, public, extensions;

create table units.members (
  -- ident
  id          uuid       primary key default uuid6(),
  unitid      uuid       not null,
  userid      uuid       not null,
  rankid      uuid       not null,
  displayname varchar(64),

  -- metadata
  isadmin   bool        not null default false,
  createdat timestamptz not null default now(),
  updatedat timestamptz not null default now(),

  -- constraints
  unique(unitid, userid),
  foreign key (unitid) references units.units(id) on delete cascade,
  foreign key (rankid) references units.ranks(id),
  foreign key (userid) references auth.users(id) on delete cascade
);

alter table units.members enable row level security;

create policy "Members are viewable by everyone"
  on units.members for select
  to authenticated, anon
  using ( true );

create policy "Admin members are creatable only by authenticated users"
  on units.members for insert
  to authenticated
  with check (
    (
      (select units.can_admin(units.members.unitid, auth.uid()))
      and units.members.isadmin = true
    ) or units.members.isadmin = false
  );

create policy "Admin members are escalated only by other admin members"
  on units.members for update
  to authenticated
  using (true)
  with check (
    (
      (select units.can_admin(units.members.unitid, auth.uid()))
      and units.members.isadmin = true
    ) or units.members.isadmin = false
  );

create view units.member_names
with (security_invoker = true)
as (
  select
    m.id as id,
    m.unitid,
    coalesce(m.displayname, u.email) as displayname,
    r.slug || '. ' || coalesce(m.displayname, u.email) as fullname
  from
    units.members m
    left join auth.users u on u.id = m.userid
    left join units.ranks r on r.id = m.rankid
);
