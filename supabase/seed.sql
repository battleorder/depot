-- create buckets
insert into storage.buckets(id, name, public) values 
  ('units_avatars', 'units_avatars', true),
  ('users_avatars', 'users_avatars', true);

-- create test users
INSERT INTO
    auth.users (
        instance_id,
        id,
        aud,
        role,
        email,
        encrypted_password,
        email_confirmed_at,
        recovery_sent_at,
        last_sign_in_at,
        raw_app_meta_data,
        raw_user_meta_data,
        created_at,
        updated_at,
        confirmation_token,
        email_change,
        email_change_token_new,
        recovery_token
    ) (
        select
            '00000000-0000-0000-0000-000000000000',
            uuid_generate_v4 (),
            'authenticated',
            'authenticated',
            'user' || (ROW_NUMBER() OVER ()) || '@battleorder.me',
            crypt ('user123!', gen_salt ('bf')),
            current_timestamp,
            current_timestamp,
            current_timestamp,
            '{"provider":"email","providers":["email"]}',
            '{}',
            current_timestamp,
            current_timestamp,
            '',
            '',
            '',
            ''
        FROM
            generate_series(1, 10)
    );

-- test user email identities
INSERT INTO
    auth.identities (
        id,
        user_id,
        provider_id,
        identity_data,
        provider,
        last_sign_in_at,
        created_at,
        updated_at
    ) (
        select
            uuid_generate_v4 (),
            id,
            id,
            format('{"sub":"%s","email":"%s"}', id :: text, email) :: jsonb,
            'email',
            current_timestamp,
            current_timestamp,
            current_timestamp
        from
            auth.users
    );

-- test units
INSERT INTO
  units.units (
    slug,
    displayname,
    tagline,
    description,
    avatar,
    createdat,
    updatedat
  ) (
    select
      'unit' || row_number() over () || 'ms',
      'Unit ' || row_number() over () || ' Milsim',
      'We are Unit ' || (row_number() over ()) || ' and we rock!',
      'Welcome to _Unit ' || (row_number() over ()) || '_. **Happy to have you.**',
      'unit' || (row_number() over ()) || '.webp',
      current_timestamp,
      current_timestamp
    from
      generate_series(1, 3)
  );

-- test ranks
INSERT INTO
  units.ranks (
    unitid,
    slug,
    displayname,
    rankorder,
    avatar,
    createdat,
    updatedat
  ) (
    select
      id,
      'Pvt',
      'Private',
      0,
      'default_low.webp',
      current_timestamp,
      current_timestamp
    from
      units.units
  );

INSERT INTO
  units.ranks (
    unitid,
    slug,
    displayname,
    rankorder,
    avatar,
    createdat,
    updatedat
  ) (
    select
      id,
      'PFC',
      'Private First Class',
      1,
      'default_low.webp',
      current_timestamp,
      current_timestamp
    from
      units.units
  );

INSERT INTO
  units.ranks (
    unitid,
    slug,
    displayname,
    rankorder,
    avatar,
    createdat,
    updatedat
  ) (
    select
      id,
      'Spc',
      'Specialist',
      2,
      'default_mid.webp',
      current_timestamp,
      current_timestamp
    from
      units.units
  );
