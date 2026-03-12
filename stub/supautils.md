


## Usage

> [supautils: Extension that secures a cluster on a cloud environment](https://github.com/supabase/supautils)

`supautils` is a loadable library that securely allows creating event triggers, publications, and extensions for non-superusers. It is completely managed by configuration -- no tables, functions, or security labels are added to your database.

### Configuration

Add to `postgresql.conf`:

```ini
shared_preload_libraries = 'supautils'
supautils.privileged_role = 'your_privileged_role'
```

Or enable per-role:

```sql
ALTER ROLE role1 SET session_preload_libraries TO 'supautils';
```

### Key GUC Parameters

| Parameter | Description |
|-----------|-------------|
| `supautils.privileged_role` | Proxy role for superuser operations |
| `supautils.superuser` | The actual superuser (defaults to bootstrap user) |
| `supautils.privileged_extensions` | Extensions allowed for non-superuser installation |
| `supautils.privileged_role_allowed_configs` | Superuser-only settings the privileged role may change |
| `supautils.reserved_roles` | Roles protected from mutation by CREATEROLE users |
| `supautils.reserved_memberships` | Role memberships restricted from being granted |
| `supautils.constrained_extensions` | JSON defining resource constraints for extensions |
| `supautils.extensions_parameter_overrides` | JSON overriding CREATE EXTENSION parameters |
| `supautils.policy_grants` | JSON granting RLS policy management to non-owners |
| `supautils.drop_trigger_grants` | JSON granting trigger drop permission to non-owners |

### Non-Superuser Publications

```sql
SET ROLE privileged_role;
CREATE PUBLICATION p FOR ALL TABLES;
DROP PUBLICATION p;
```

### Privileged Extensions

```ini
supautils.privileged_extensions = 'hstore'
```

Non-superusers can then create extensions that normally require superuser:

```sql
CREATE EXTENSION hstore;
```

### Reserved Roles

```ini
supautils.reserved_roles = 'connector, storage_admin'
```

Users with CREATEROLE cannot ALTER or DROP these roles.

### Table Ownership Bypass (RLS Policy Management)

```ini
supautils.policy_grants = '{ "my_role": ["public.not_my_table"] }'
```

Allows `my_role` to manage RLS policies on tables it does not own.
