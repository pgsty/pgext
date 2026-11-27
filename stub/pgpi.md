## Usage

Sources:

- [README.md](https://github.com/s-hironobu/pg_plan_inspector/blob/7d929b64421f229800fe84838c36a414e7159de4/README.md)
- [pgpi.control](https://github.com/s-hironobu/pg_plan_inspector/blob/7d929b64421f229800fe84838c36a414e7159de4/extension/pgpi.control)
- [pgpi--0.4.sql](https://github.com/s-hironobu/pg_plan_inspector/blob/7d929b64421f229800fe84838c36a414e7159de4/extension/pgpi--0.4.sql)
- [Makefile](https://github.com/s-hironobu/pg_plan_inspector/blob/7d929b64421f229800fe84838c36a414e7159de4/extension/Makefile)
- [pgpi.c](https://github.com/s-hironobu/pg_plan_inspector/blob/7d929b64421f229800fe84838c36a414e7159de4/extension/pgpi.c)

`pgpi` is the PostgreSQL extension in the `pg_plan_inspector` research framework. Its control file and installation SQL use version 0.4; the framework calls this line 2.0. It records execution telemetry and can replace planner cardinality estimates with directives or answers from an external estimator.

### Enable and Inspect

The current source requires PostgreSQL 19. It supports Linux and macOS on x86_64/ARM64; runtime function detours are available only on Linux. On macOS the official hook route remains available. Enable preload and select the database that stores logs, then restart PostgreSQL:

```ini
shared_preload_libraries = 'pgpi'
pgpi.log_database = 'postgres'
```

Append to any existing preload list. Install the extension as a superuser in each database where it is needed:

```sql
CREATE EXTENSION pgpi;
SELECT * FROM pgpi_status();
SELECT pgpi_log_status();
```

The worker creates `pgpi.plans` and `pgpi.stats` in the logging database. SQL functions are relocatable, while these log tables remain in their own schema.

### Capture and Test Estimates

Intervention defaults to off and plan capture is separate. A privileged test session can capture a workload and flush the buffer before collecting it:

```sql
SET pgpi.log_plans = on;
SET pgpi.label = 'baseline';
SELECT count(*) FROM pg_class;
SELECT pgpi_log_flush(10000);
SET pgpi.log_plans = off;
```

`pgpi.log_plans` requires superuser permission or an explicit parameter grant. The Python client needs Python 3.10+ and dependencies for the selected estimator. It collects logs into a Parquet dataset, then serves estimates over a UNIX socket. Point `pgpi.callback` at that socket and use `pgpi.mode = 'shadow'` to record answers without applying them; `enforce` applies them. Static answers can instead be supplied through `pgpi.directives`.

### Objects and Boundaries

- `pgpi_requests` and `pgpi_clear` inspect or clear the current backend's request log. `pgpi_status` reports installed detours; `pgpi_log_status` exposes shared-buffer and worker statistics.
- `pgpi_log_flush` returns false on timeout or when not preloaded. `pgpi_stats_reset_seen` resets remembered statistics versions and is not executable by `PUBLIC` by default.
- The callback timeout defaults to 2000 ms. On failure, requests are suspended for five seconds and PostgreSQL estimates are used. An estimator can also abstain.
- The default retention is unlimited: configure `pgpi.log_retention` or `pgpi.log_max_rows` for bounded plan storage. `pgpi.log_database`, `pgpi.log_buffer_size`, and `pgpi.stats_max` require restart when changed.
- Logs can include SQL, bind parameters, and statistics. Restrict access accordingly. Capture adds execution instrumentation; estimates and detours are experimental and can change plans or performance. The project does not claim that better cardinality accuracy always makes queries faster.
