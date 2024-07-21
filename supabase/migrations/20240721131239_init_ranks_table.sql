set search_path = units, public, extensions;

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
alter table units.ranks enable row level security;

create policy "Ranks are viewable by everyone"
  on units.ranks for select
  to authenticated, anon
  using ( true );

create policy "Ranks are creatable by a unit's owner"
  on units.ranks for insert
  to authenticated, anon
  with check (
    exists (
      select 1 from units.units
      where units.units.id = units.ranks.unitid
      and units.units.ownerid = (select auth.uid())
    )
  );


