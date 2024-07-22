set search_path = units, public, extensions;

create table units.ranks (
  -- ident
  id          uuid        primary key default uuid6(),
  unitid      uuid        not null,
  slug        varchar(32) not null,
  displayname varchar(64) not null,

  -- branding
  rankorder int  not null default 0,
  avatar    text,

  -- metadata
  createdat timestamptz not null default now(),
  updatedat timestamptz not null default now(),

  -- constraints
  unique(unitid, slug),
  unique(unitid, displayname),
  foreign key (unitid) references units.units(id) on delete cascade
);
alter table units.ranks enable row level security;

create policy "Ranks are viewable by everyone"
  on units.ranks for select
  to authenticated, anon
  using ( true );

create policy "Ranks are creatable by a unit's owner"
  on units.ranks for insert
  to authenticated, anon
  with check (units.can_admin(unitid, auth.uid()));
