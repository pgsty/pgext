## Usage

Sources:

- [Official extension control file](https://github.com/CaghanTU/pg_flashback/blob/7f0a26d763372859826f8e21d9b1594dce8dc414/pg_flashback.control)
- [Official upstream documentation](https://github.com/CaghanTU/pg_flashback/blob/7f0a26d763372859826f8e21d9b1594dce8dc414/README.md)
- [Official Rust package manifest](https://github.com/CaghanTU/pg_flashback/blob/7f0a26d763372859826f8e21d9b1594dce8dc414/Cargo.toml)

`pg_flashback` 0.1.0 is a Rust and pgrx extension for table-level point-in-time restore and read-only historical queries. It captures DML with triggers or asynchronous logical-WAL consumption, records DDL through a utility hook, and reconstructs a target table from snapshots plus delta events. It is not a substitute for cluster backups or PostgreSQL PITR.

### Core Workflow

Preload the module and restart PostgreSQL. WAL capture also requires logical WAL; trigger mode can operate without that setting.

```conf
wal_level = logical
shared_preload_libraries = 'pg_flashback'
```

```sql
CREATE EXTENSION pg_flashback;

SELECT flashback_track('public.orders');
SELECT flashback_checkpoint('public.orders');

SELECT flashback_restore(
  'public.orders',
  now() - interval '30 seconds'
);
```

`flashback_track(table)` creates capture objects, a base snapshot, and schema metadata. `flashback_untrack(table)` stops tracking and removes that history. `flashback_restore(table, timestamptz)` replaces table state at a requested time; an array overload restores multiple tables in foreign-key order.

### Query and Recovery APIs

`flashback_query(table, timestamptz, filter_clause)` reconstructs a temporary historical state and returns `SETOF record`, so callers must provide a column definition list. The filter is only a `WHERE` predicate. `flashback_recover_deleted(table, timestamptz)` re-inserts rows that are missing now and requires a primary key. `flashback_restore_parallel(table, timestamptz, num_workers)` supplies parallel-query hints rather than parallel delta replay.

Use `flashback_retention_status()`, `flashback.pg_stat_flashback`, `flashback_history(table, interval)`, and `flashback.restore_log` to inspect retention, pending work, change history, and restore auditing.

### Configuration and Safety

`pg_flashback.capture_mode` selects `auto`, `wal`, or `trigger`; `pg_flashback.enabled` is the global capture switch. Worker database, interval, batch, row-size, retention, slot, and restore-memory settings control the remaining behavior. WAL mode depends on an asynchronously consumed logical slot, while trigger mode adds synchronous DML overhead. A restore performs a shadow-table rebuild and atomic swap with a brief exclusive lock; test dependencies, RLS, sequences, partitions, inheritance, crash recovery, worker lag, retention, and disk growth before relying on it.

The official repository and pinned files currently return HTTP 404. This document reflects the reviewed upstream snapshot; source availability and maintenance status must be re-established before adoption.
