begin;
  create table wh_credential_dimension (
    -- random id generated using encode(digest(gen_random_bytes(16), 'sha256'), 'base64')
    -- this is done to prevent conflicts with rows in other clusters
    -- which enables warehouse data from multiple clusters to be loaded into a
    -- single database instance
    key                                        wh_dim_id    primary key default wh_dim_id(),

    credential_library_id                      wh_public_id not null,
    credential_library_type                    wh_dim_text,
    credential_library_name                    wh_dim_text,
    credential_library_description             wh_dim_text,
    credential_library_vault_path              wh_dim_text,
    credential_library_vault_http_method       wh_dim_text,
    credential_library_vault_http_request_body wh_dim_text,

    credential_store_id                        wh_public_id not null,
    credential_store_type                      wh_dim_text,
    credential_store_name                      wh_dim_text,
    credential_store_description               wh_dim_text,
    credential_store_vault_namespace           wh_dim_text,
    credential_store_vault_address             wh_dim_text,

    target_id                                  wh_public_id not null,
    target_type                                wh_dim_text,
    target_name                                wh_dim_text,
    target_description                         wh_dim_text,
    target_default_port_number                 integer      not null,
    target_session_max_seconds                 integer      not null,
    target_session_connection_limit            integer      not null,

    project_id                                 wt_scope_id  not null,
    project_name                               wh_dim_text,
    project_description                        wh_dim_text,

    organization_id                            wt_scope_id  not null,
    organization_name                          wh_dim_text,
    organization_description                   wh_dim_text,

    current_row_indicator                      wh_dim_text,
    row_effective_time                         wh_timestamp,
    row_expiration_time                        wh_timestamp
  );

  create unique index wh_credential_dim_current_constraint
    on wh_credential_dimension (credential_library_id, credential_store_id, target_id)
    where current_row_indicator = 'Current';

  create table wh_credential_group (
    -- random id generated using encode(digest(gen_random_bytes(16), 'sha256'), 'base64')
    -- this is done to prevent conflicts with rows in other clusters
    -- which enables warehouse data from multiple clusters to be loaded into a
    -- single database instance
    key wh_dim_id primary key default wh_dim_id()
  );

  create table wh_credential_group_membership (
    credential_group_key wh_dim_id not null
      references wh_credential_group (key)
      on delete restrict
      on update cascade,
    credential_key       wh_dim_id not null
      references wh_credential_dimension (key)
      on delete restrict
      on update cascade,
    credential_purpose   wh_dim_text
  );

  insert into wh_credential_group
    (key)
  values
    ('no credentials');

  insert into wh_credential_dimension (
    key,
    credential_library_id, credential_library_type, credential_library_name,  credential_library_description, credential_library_vault_path,    credential_library_vault_http_method, credential_library_vault_http_request_body,
    credential_store_id,   credential_store_type,   credential_store_name,    credential_store_description,   credential_store_vault_namespace, credential_store_vault_address,
    target_id,             target_type,             target_name,              target_description,             target_default_port_number,       target_session_max_seconds,           target_session_connection_limit,
    project_id,            project_name,            project_description,
    organization_id,       organization_name,       organization_description,
    current_row_indicator, row_effective_time,      row_expiration_time
  ) values (
    'no credential',
    'not applicable',      'not applicable',        'not applicable',         'not applicable',               'not applicable',                 'not applicable',                     'not applicable',
    'not applicable',      'not applicable',        'not applicable',         'not applicable',               'not applicable',                 'not applicable',
    'not applicable',      'not applicable',        'not applicable',         'not applicable',               0,                                0,                                    0,
    'not applicable',      'not applicable',        'not applicable',
    'not applicable',      'not applicable',        'not applicable',
    'Current',             now(),                   'infinity'::timestamptz
  );

  insert into wh_credential_group_membership
    (credential_group_key, credential_key, credential_purpose)
  values
    ('no credentials', 'no credential', 'None');
commit;
