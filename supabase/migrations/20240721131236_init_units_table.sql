set search_path = units, public, extensions;

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
  ownerid uuid not null,
  createdat timestamptz not null default now(),
  updatedat timestamptz not null default now(),

  -- constraints
  foreign key(ownerid) references auth.users(id)
);

alter table units.units enable row level security;

create policy "Units are viewable by everyone"
  on units.units for select
  to authenticated, anon
  using ( true );

create policy "Units are creatable only by authenticated users"
  on units.units for insert
  to authenticated
  with check (
    (select auth.uid()) = units.units.ownerid
  );
