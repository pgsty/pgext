## Usage

Sources:

- [README.md](https://github.com/profundium/pg_local_cache/blob/940f796de41bf4ad22de3f8efc32dcf6d6b791b9/README.md)
- [pg_local_cache.control](https://github.com/profundium/pg_local_cache/blob/940f796de41bf4ad22de3f8efc32dcf6d6b791b9/pg_local_cache.control)
- [sql/pg_local_cache--3.0.0.sql](https://github.com/profundium/pg_local_cache/blob/940f796de41bf4ad22de3f8efc32dcf6d6b791b9/sql/pg_local_cache--3.0.0.sql)
- [sql/pg_local_cache--2.0.4--3.0.0.sql](https://github.com/profundium/pg_local_cache/blob/940f796de41bf4ad22de3f8efc32dcf6d6b791b9/sql/pg_local_cache--2.0.4--3.0.0.sql)
- [docs/UPGRADING.md](https://github.com/profundium/pg_local_cache/blob/940f796de41bf4ad22de3f8efc32dcf6d6b791b9/docs/UPGRADING.md)
- [CHANGELOG.md](https://github.com/profundium/pg_local_cache/blob/940f796de41bf4ad22de3f8efc32dcf6d6b791b9/CHANGELOG.md)

`pg_local_cache` 3.0.0 caches whole rows by their complete primary key in bounded shared memory. Cached reads now use authenticated RESP MGET; SQL `local_cache.mget` is removed and ordinary SELECT queries are not rewritten. PostgreSQL remains the durable source of truth.

### Core Workflow

```ini
shared_preload_libraries = 'pg_local_cache'
pg_local_cache.database = 'app'
pg_local_cache.role = 'local_cache_worker'
pg_local_cache.bind_address = '127.0.0.1'
pg_local_cache.port = 6380
pg_local_cache.auth_token_file = '/secure/path/token'
```

```sql
CREATE EXTENSION pg_local_cache;
CREATE TABLE public.items (id bigint PRIMARY KEY, value text);
INSERT INTO public.items VALUES (42, 'example');
SELECT local_cache.attach_table('public.items'::regclass);
SELECT local_cache.health();
```

After authenticating the RESP connection:

```text
MGET CRUD:app.public.items:{"id":42} CRUD:app.public.items:{"id":7}
```

### Operational Boundaries

Append the library to existing preload entries, provision a dedicated worker role and its documented metadata/table grants, secure the token file, and restart. A superuser creates the extension in fixed schema `local_cache`. Attach only supported permanent primary-key tables. RESP MGET preserves order and duplicate keys and returns null for missing rows. Workers use the configured database role, not the client’s SQL privileges, transaction or snapshot. Keep joins, projections, row locks and session-sensitive reads in ordinary SQL.

`local_cache.attach_table` installs invalidation triggers; `local_cache.detach_table` removes a mapping; `local_cache.reconcile_table` revalidates it after DDL or privilege changes. `local_cache.health`, `local_cache.stats` and `local_cache.metrics` report readiness and resource use. Normal PostgreSQL writes invalidate affected cache entries. RLS, partitioned and inherited tables are outside the documented workload. There is no TTL or distributed cache coordination.

Supports PostgreSQL 14–18 on one writable primary. Non-loopback listeners require `pg_local_cache.tls` or explicit `pg_local_cache.allow_plaintext_network` opt-in. TLS uses separate certificate/key settings; a CA setting enables mutual TLS. The reloadable `pg_local_cache.enabled` switch disables caching.

Before upgrading from 2.x, move SQL MGET callers to the RESP authorization model. Install matching 3.0.0 files, configure listener security, restart, confirm the library version and listener readiness, then run `ALTER EXTENSION pg_local_cache UPDATE` in each affected database. Dependencies on removed functions or the replaced metrics return type can block migration; the script deliberately avoids CASCADE. No downgrade script is supplied.
