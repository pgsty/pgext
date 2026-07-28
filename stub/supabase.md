## Usage

Sources:

- [Official upstream README](https://github.com/constructive-io/supabase-test-suite/blob/44ceff6be0b62e1d1a60524acc0c65c1b8343726/packages/supabase/README.md)
- [Official extension control file (supabase.control)](https://github.com/constructive-io/supabase-test-suite/blob/44ceff6be0b62e1d1a60524acc0c65c1b8343726/packages/supabase/supabase.control)
- [Official extension SQL (supabase--0.0.1.sql)](https://github.com/constructive-io/supabase-test-suite/blob/44ceff6be0b62e1d1a60524acc0c65c1b8343726/packages/supabase/sql/supabase--0.0.1.sql)

`supabase` — Supabase-focused SQL, tests, and helpers for building robust applications with Row-Level Security. Use it when implementing the corresponding security, audit, or access-control workflow. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION supabase;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `auth.email()` is an extension function and returns `text`.
- `auth.role()` is an extension function and returns `text`.
- `auth.uid()` is an extension function and returns `uuid`.
- `extensions.grant_pg_cron_access()` is an extension function and returns `event_trigger`.
- `extensions.grant_pg_graphql_access()` is an extension function and returns `event_trigger`.
- `extensions.grant_pg_net_access()` is an extension function and returns `event_trigger`.
- `extensions.pgrst_ddl_watch()` is an extension function and returns `event_trigger`.
- `extensions.pgrst_drop_watch()` is an extension function and returns `event_trigger`.
- `extensions.set_graphql_placeholder()` is an extension function and returns `event_trigger`.
- `graphql_public.graphql("operationName" text default null, query text default null, variables jsonb default null, extensions jsonb default null)` is an extension function and returns `jsonb`.
- `pgbouncer.get_auth(p_usename text)` is an extension function and returns `TABLE`.
- `storage.add_prefixes(_bucket_id text, _name text)` is an extension function and returns `void`.
- `storage.delete_prefix(_bucket_id text, _name text)` is an extension function and returns `boolean`.
- `storage.delete_prefix_hierarchy_trigger()` is an extension function and returns `trigger`.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- Install the confirmed extension dependencies first: `plpgsql`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
