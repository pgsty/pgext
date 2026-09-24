## Usage

Sources:

- [README](https://github.com/MEGALab/pgeverything/blob/517ff96f6c398b38008ccdb4586d2ab691921893/README.md)
- [Control](https://github.com/MEGALab/pgeverything/blob/517ff96f6c398b38008ccdb4586d2ab691921893/extensions/pgcache/pgcache.control)
- [SQL](https://github.com/MEGALab/pgeverything/blob/517ff96f6c398b38008ccdb4586d2ab691921893/extensions/pgcache/pgcache--0.1.0.sql)

`pgcache` is PGEverything's SQL key/value cache with JSON values and optional expiration. This 0.1.0 source snapshot is used in the upstream PostgreSQL 16 container; it stores disposable cache data, not durable application state.

### Core Workflow

Install the SQL/control files and enable the extension in the target database. PL/pgSQL is needed for deletion and cleanup helpers. There is no shared library or preload step.

```sql
CREATE EXTENSION pgcache;
SELECT cache_set('user:1', '{"name":"alice"}'::jsonb, ttl => 60);
SELECT cache_get('user:1');
SELECT cache_incr('hits');
SELECT cache_del('user:1');
SELECT pgcache.purge_expired();
```

### API and Storage

| Function | Behavior |
| --- | --- |
| `cache_set(text,jsonb,integer)` | Insert or replace a value; TTL is seconds, NULL means no expiration |
| `cache_get(text)` | Return JSON, or NULL for a missing/expired entry |
| `cache_del(text)` | Delete a key and report whether a row existed |
| `cache_incr(text,bigint)` | Atomically increment an integer JSON scalar; default delta is 1, and an update clears its TTL |
| `pgcache.purge_expired()` | Delete expired rows and return the count |

The store is the unlogged table `pgcache.store`. Crash recovery truncates unlogged data, and it is not a normal physical-replication data source. Reads hide expired rows but do not remove them; schedule cleanup separately if needed. `pg_cron` is an optional external scheduler, not an installation dependency.

### Privileges and Limits

These are invoker-rights functions: callers still need appropriate schema/table privileges. The control file sets `superuser=false` but does not mark the extension trusted; ordinary DDL privileges must suffice for installation. The fixed internal schema makes the extension non-relocatable. No explicit license or support matrix beyond the upstream PostgreSQL 16 deployment is supplied.
