


## Usage

> [pg_squeeze: A tool to remove unused space from a relation.](https://github.com/cybertec-postgresql/pg_squeeze)

`pg_squeeze` requires `wal_level = logical` and must be added to `shared_preload_libraries`. It removes bloat from tables while allowing concurrent reads and writes, using logical decoding instead of triggers.

### Register a Table for Scheduled Processing

Insert into `squeeze.tables` to enable periodic bloat checks:

```sql
INSERT INTO squeeze.tables (tabschema, tabname, schedule)
VALUES ('public', 'foo', ('{30}', '{22}', NULL, NULL, '{3, 5}'));
```

The `schedule` field uses a crontab-like format: `(minutes, hours, days_of_month, months, days_of_week)`. The above checks table `foo` every Wednesday and Friday at 22:30.

Optional columns: `free_space_extra` (min % extra free space to trigger, default 50), `min_size` (min MB, default 8), `vacuum_max_age` (max time since last VACUUM, default 1h), `max_retry` (retry count, default 0), `clustering_index` (sort tuples by this index), `rel_tablespace`, `ind_tablespaces`, `skip_analyze`.

### Ad-hoc Squeeze

```sql
SELECT squeeze.squeeze_table('public', 'pgbench_accounts');
SELECT squeeze.squeeze_table('public', 'mytable', 'my_cluster_idx', 'target_tablespace');
```

### Start / Stop Workers

```sql
SELECT squeeze.start_worker();   -- start scheduler + squeeze workers
SELECT squeeze.stop_worker();    -- stop all workers for current database
```

Auto-start on cluster boot via `postgresql.conf`:

```
squeeze.worker_autostart = 'my_database your_database'
squeeze.worker_role = postgres
```

### Monitoring

- **`squeeze.log`** -- one entry per successfully squeezed table (with `started`, `finished`, `ins_initial`, `ins`, `upd`, `del`)
- **`squeeze.errors`** -- errors during squeezing
- **`squeeze.get_active_workers()`** -- shows currently active squeeze workers and their progress

### Configuration

- **`squeeze.max_xlock_time`** -- max exclusive lock time in ms (default unlimited)
- **`squeeze.workers_per_database`** -- number of concurrent squeeze workers (default 1)
