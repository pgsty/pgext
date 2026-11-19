## Usage

Sources:

- [Official source](https://github.com/ken19990709/pg_query_progress/blob/b0009ff4de03de70ac83ffd23ed1e696e9b81bc7/README.md)
- [Official source](https://github.com/ken19990709/pg_query_progress/blob/b0009ff4de03de70ac83ffd23ed1e696e9b81bc7/pg_query_progress.control)
- [Official source](https://github.com/ken19990709/pg_query_progress/blob/b0009ff4de03de70ac83ffd23ed1e696e9b81bc7/pg_query_progress--1.0.sql)
- [Official source](https://github.com/ken19990709/pg_query_progress/blob/b0009ff4de03de70ac83ffd23ed1e696e9b81bc7/pg_query_progress.c)
- [Official source](https://github.com/ken19990709/pg_query_progress/blob/b0009ff4de03de70ac83ffd23ed1e696e9b81bc7/pgqp_compat.h)
- [Official source](https://github.com/ken19990709/pg_query_progress/blob/b0009ff4de03de70ac83ffd23ed1e696e9b81bc7/LICENSE)

`pg_query_progress` reports cost-weighted estimates for running SELECT statements through a view/function and an optional NOTICE progress bar. The control and installation SQL version is 1.0; the repository's 4.0.0 labels describe its source distribution.

### Enable and monitor

Choose the branch for the server major: PG15, PG16, PG17 or PG18. The reviewed main branch targets PostgreSQL 20 development, not PostgreSQL 19. Linux is the tested platform. Shared preload and a server restart are mandatory before the monitoring objects are used.

```ini
shared_preload_libraries = 'pg_query_progress'
```

```sql
CREATE EXTENSION pg_query_progress;
GRANT EXECUTE ON FUNCTION pg_query_progress() TO pg_monitor;
GRANT SELECT ON pg_query_progress TO pg_monitor;
SET pg_query_progress.bar = on;
SELECT pid, state, progress, tuples_processed, age_ms
FROM pg_query_progress WHERE state = 'running';
```

Preserve other preload entries. Poll from a separate session to observe long queries. The function and view revoke PUBLIC access because query text can contain sensitive information; grant both function execution and view access only to intended monitoring roles. A plain `LOAD` does not initialize shared memory: the view returns no rows and warns.

### Configuration and limits

`pg_query_progress.track` is superuser-controlled. Session settings include `pg_query_progress.bar`, `pg_query_progress.bar_ansi`, `pg_query_progress.bar_interval_ms`, `pg_query_progress.bar_min_query_ms` and `pg_query_progress.max_query_len`. The bar uses NOTICE, so client message filtering can hide it.

`progress` is a fraction from 0 to 1, not a measured completion percentage. Bad row estimates and blocking plans distort progress and ETA. DML, DDL and utilities are not tracked. There is one tracked portal per session; another query can mark an open cursor as terminated. Shared-memory records disappear after a server crash or restart. Parallel workers contribute estimates, but blocked parallel sections may produce few intermediate updates.
