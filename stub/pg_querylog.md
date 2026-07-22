## Usage

Sources:

- [Official README](https://github.com/adjust/pg_querylog/blob/f639adb6dc1686044b76d82139df1c3a49076053/README.md)
- [Installed SQL objects](https://github.com/adjust/pg_querylog/blob/f639adb6dc1686044b76d82139df1c3a49076053/init.sql)
- [Shared-memory and GUC implementation](https://github.com/adjust/pg_querylog/blob/f639adb6dc1686044b76d82139df1c3a49076053/pg_querylog.c)

`pg_querylog` captures each backend's current or most recently completed query and rendered parameters in shared memory. Use it for short-lived diagnostics on old, explicitly tested PostgreSQL versions; captured text can contain credentials, personal data, and application secrets.

### Load and Query

Add the library to `shared_preload_libraries` and restart PostgreSQL, then install its SQL objects in a restricted schema:

```conf
shared_preload_libraries = 'pg_querylog'
pg_querylog.buffer_size = '16kB'
```

```sql
CREATE SCHEMA querylog;
REVOKE ALL ON SCHEMA querylog FROM PUBLIC;
CREATE EXTENSION pg_querylog SCHEMA querylog;

SET pg_querylog.enabled = on;
SELECT * FROM querylog.get_queries(false, true);
SELECT * FROM querylog.running_queries;
```

`get_queries(only_running, skip_overflow)` returns `pid`, `query`, `params`, `start_time`, `end_time`, `running`, and `overflow`. `running_queries` filters to active entries. An overflow flag means the per-backend buffer could not retain the complete item.

### Configuration and Security

`pg_querylog.buffer_size` is a superuser setting with a 10 kB minimum and is allocated per backend; plan approximately that value times `max_connections`, plus overhead. `pg_querylog.enabled` is also superuser-only. The README calls it disabled by default, while the reviewed C source initializes it to true; set it explicitly and verify with `SHOW pg_querylog.enabled`.

The README also permits `session_preload_libraries`, but warns that every starting backend rewrites the shared enabled value. Prefer cluster preload for deterministic behavior. Restrict both schema and function/view privileges, define a short diagnostic window, and disable collection afterward. This unmaintained code hooks executor internals and lists only old PostgreSQL releases, so test crashes, prepared statements, parameter rendering, concurrent backends, truncation, and memory use on the exact target build.
