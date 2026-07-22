## Usage

Sources:

- [Pinned upstream README](https://github.com/frost242/ssi_failures/blob/8b6b38b1f688632c3520d1b6aaac24feb7734c84/README.md)
- [Pinned version 0.1 installation SQL](https://github.com/frost242/ssi_failures/blob/8b6b38b1f688632c3520d1b6aaac24feb7734c84/pg_stat_ssi_failures--0.1.sql)
- [Pinned C implementation](https://github.com/frost242/ssi_failures/blob/8b6b38b1f688632c3520d1b6aaac24feb7734c84/pg_stat_ssi_failures.c)
- [Pinned extension control file](https://github.com/frost242/ssi_failures/blob/8b6b38b1f688632c3520d1b6aaac24feb7734c84/pg_stat_ssi_failures.control)

`pg_stat_ssi_failures` version `0.1` maintains one cluster-wide counter for errors with SQLSTATE `serialization_failure`. It is intended to expose a cumulative metric for monitoring serializable-transaction conflicts, not to identify the statement, database, role, or transaction that failed.

### Core Workflow

The module must reserve shared memory at server start. Add it to `shared_preload_libraries` and restart PostgreSQL before creating the SQL extension:

```conf
shared_preload_libraries = 'pg_stat_ssi_failures'
```

```sql
CREATE EXTENSION pg_stat_ssi_failures;

SELECT pg_stat_ssi_failures();
SELECT pg_stat_ssi_failures_reset();
```

For interval monitoring, store the last observed value in the monitoring system and subtract it from the current cumulative count. Handle a lower value as a reset or restart rather than as a negative failure rate.

### Important Objects

- `pg_stat_ssi_failures() RETURNS bigint`: reads the shared cluster-wide counter.
- `pg_stat_ssi_failures_reset() RETURNS void`: resets that counter to zero for every database in the instance.

### Operational Notes

The hook counts every emitted error carrying PostgreSQL's serialization-failure SQLSTATE, not a per-query subset, and the value has no labels. A clean shutdown writes the counter to a server-side statistics file; the source deliberately removes that file after startup and excludes it from backup and replication behavior. Crash recovery, restore, promotion, reset, or extension removal can therefore create a discontinuity.

Revoke `pg_stat_ssi_failures_reset()` from application roles because the installation SQL adds no privilege guard. The reviewed reset implementation writes shared state while holding only a shared LWLock, so avoid concurrent resets and treat the function as an administrator-only maintenance operation. This old source uses PostgreSQL internal hooks and lock APIs; verify compilation and hook chaining on the exact target release, and confirm that other preloaded modules do not interfere before production monitoring.
