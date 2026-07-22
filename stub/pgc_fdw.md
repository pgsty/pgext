## Usage

Sources:

- [Official README](https://github.com/fengttt/pgc_fdw/blob/5993bb88886c0a27d2ee6086f68255ba9a8460c2/README.md)
- [Version 1.0 SQL objects](https://github.com/fengttt/pgc_fdw/blob/5993bb88886c0a27d2ee6086f68255ba9a8460c2/pgc_fdw--1.0.sql)
- [FDW option validation](https://github.com/fengttt/pgc_fdw/blob/5993bb88886c0a27d2ee6086f68255ba9a8460c2/option.c)
- [FoundationDB cache implementation](https://github.com/fengttt/pgc_fdw/blob/5993bb88886c0a27d2ee6086f68255ba9a8460c2/cache.c)

`pgc_fdw` is an old fork of `postgres_fdw` that caches remote query results in FoundationDB. It creates a foreign data wrapper named `pgc_fdw`, so it can coexist with the core `postgres_fdw`, but it tracks an old PostgreSQL source tree and should be treated as experimental.

### Create a Cached Foreign Table

FoundationDB server and client libraries must be installed, and the PostgreSQL operating-system account must be able to use the default FoundationDB cluster file.

```sql
CREATE EXTENSION pgc_fdw;

CREATE SERVER cached_remote
  FOREIGN DATA WRAPPER pgc_fdw
  OPTIONS (host 'db.internal', dbname 'app', cache_timeout '60');

CREATE USER MAPPING FOR CURRENT_USER
  SERVER cached_remote
  OPTIONS (user 'remote_user', password 'secret');

CREATE FOREIGN TABLE cached_items (id bigint, payload text)
  SERVER cached_remote
  OPTIONS (schema_name 'public', table_name 'items', cache_timeout '30');

SELECT * FROM cached_items;
```

`cache_timeout` is valid on a server or foreign table, is measured in seconds, defaults to `3600`, and can be set to `0` to disable caching. Negative values are rejected.

### Cache Operations and Risks

```sql
SELECT * FROM pgc_fdw_cache_info();
SELECT pgc_fdw_invalidate('cache-entry-sha');
```

`pgc_fdw_cache_info()` reports `sha`, timestamp, tuple count, and query. The extension also exposes low-level `pgc_fdw_set()` and `pgc_fdw_watch()` functions. Revoke cache mutation functions from `PUBLIC` unless callers explicitly need them.

Cached reads can remain stale for the full timeout, and the extension provides no general invalidation link to remote writes. Test transaction visibility, failover, FoundationDB outages and limits, credential handling, cache-key isolation, and manual invalidation. Because the fork depends on PostgreSQL internals and starts FoundationDB networking in each loading backend, build and stress-test it only against the exact PostgreSQL/FoundationDB combination intended for deployment.
