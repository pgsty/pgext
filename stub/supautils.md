## Usage

Sources:

- [v3.4.4 README](https://github.com/supabase/supautils/blob/v3.4.4/README.md)
- [v3.4.4 release](https://github.com/supabase/supautils/releases/tag/v3.4.4)
- [Version restriction implementation](https://github.com/supabase/supautils/blob/v3.4.4/src/extensions.c)

`supautils` is a loadable library that unlocks selected superuser-only PostgreSQL features for non-superusers through configuration. Upstream emphasizes that it adds no tables, functions, or security labels to the database.

### Load it

Cluster-wide:

```ini
shared_preload_libraries = 'supautils'
supautils.privileged_role = 'your_privileged_role'
```

Per role:

```sql
ALTER ROLE role1 SET session_preload_libraries TO 'supautils';
```

### Privileged role capabilities

The README documents a privileged proxy role that can create publications, foreign data wrappers, event triggers, and privileged extensions without granting `SUPERUSER`.

```sql
SET ROLE privileged_role;
CREATE PUBLICATION p FOR ALL TABLES;
DROP PUBLICATION p;
```

For event triggers, the README says privileged-role triggers run for non-superusers, skip superusers, and also skip reserved roles. It also documents one limitation: those triggers do not fire while creating publications, foreign data wrappers, or extensions.

### Important configuration knobs

- `supautils.superuser`
- `supautils.privileged_role`
- `supautils.privileged_role_allowed_configs`
- `supautils.privileged_extensions`
- `supautils.extension_custom_scripts_path`
- `supautils.constrained_extensions`
- `supautils.extensions_parameter_overrides`
- `supautils.policy_grants`
- `supautils.drop_trigger_grants`
- `supautils.reserved_roles`
- `supautils.reserved_memberships`
- `supautils.hint_roles`
- `supautils.log_skipped_evtrigs`

### Useful examples

Allow a non-superuser to create specific privileged extensions:

```ini
supautils.privileged_extensions = 'hstore'
```

Allow a role to manage RLS policies on tables it does not own:

```ini
supautils.policy_grants = '{ "my_role": ["public.not_my_table"] }'
```

Force an extension into a specific schema on `CREATE EXTENSION`:

```ini
supautils.extensions_parameter_overrides = '{ "pg_cron": { "schema": "pg_catalog" } }'
```

Protect managed-service roles from `CREATEROLE` users:

```ini
supautils.reserved_roles = 'connector, storage_admin'
supautils.reserved_memberships = 'pg_read_server_files'
```

### Version Selection and Operational Boundaries

`supautils.restrict_extension_versions` controls explicit version clauses for non-superusers: `off` allows them, `warn` ignores them and selects the control-file default with a warning, and `error` rejects them. This applies to both extension creation and upgrades; superusers and the configured proxy superuser are exempt. Omitting an explicit version remains allowed subject to normal privilege checks.

Cluster preload requires a restart; role-specific session preload applies to new connections. Do not run CREATE EXTENSION for supautils itself. Source release 3.4.4 is a library update and has no SQL extension-update step. It avoids ACCESS EXCLUSIVE locks during allowlisted-table policy checks and restores the caller's role on every exit from an elevated region.

Review allowed extensions and custom scripts as trusted code because their operations run with delegated superuser privileges. Enhanced privilege hints do not work for views on PostgreSQL 18 according to the tagged README. Test role transitions, event-trigger ownership and reserved-role protections before broadening grants.
