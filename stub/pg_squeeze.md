## Usage

Sources:

- [pg_squeeze REL1_9_6 README](https://github.com/cybertec-postgresql/pg_squeeze/blob/REL1_9_6/README.md)
- [pg_squeeze release notes](https://github.com/cybertec-postgresql/pg_squeeze/blob/REL1_9_6/NEWS)
- [pg_squeeze extension control](https://github.com/cybertec-postgresql/pg_squeeze/blob/REL1_9_6/pg_squeeze.control)

`pg_squeeze` removes bloat from a table and its indexes while allowing concurrent reads and writes. It copies live tuples to new storage and applies concurrent changes through logical decoding, avoiding the long exclusive lock of `VACUUM FULL`. Use it only after sizing replication slots, disk space, and the table's replica identity.

### Configure and Install

```ini
max_replication_slots = 1  # or add one to the existing requirement
shared_preload_libraries = 'pg_squeeze'
output_plugin_libraries = 'pg_squeeze'
wal_level = logical       # required on PostgreSQL versions before 19
```

Add `pg_squeeze` to any existing library lists. On PostgreSQL releases with `output_plugin_libraries`, that list must permit `pg_squeeze` so logical decoding can start. Restart PostgreSQL, then create the extension:

```sql
CREATE EXTENSION pg_squeeze;
```

The table must have an identity index. A primary key works with the default replica identity; otherwise select a suitable unique index with `ALTER TABLE ... REPLICA IDENTITY USING INDEX`.

### Run an Ad-Hoc Squeeze

```sql
SELECT squeeze.squeeze_table('public', 'pgbench_accounts');

SELECT squeeze.squeeze_table(
  'public',
  'large_table',
  'large_table_cluster_idx',
  'target_tablespace'
);
```

The function starts background work and is not transactional in the ordinary SQL-function sense. Monitor the operation rather than assuming a surrounding `ROLLBACK` cancels it.

### Schedule Tables and Monitor Work

```sql
INSERT INTO squeeze.tables (tabschema, tabname, schedule)
VALUES ('public', 'events', ('{30}', '{22}', NULL, NULL, '{3,5}'));

SELECT * FROM squeeze.get_active_workers();
SELECT * FROM squeeze.log ORDER BY finished DESC;
SELECT * FROM squeeze.errors;
```

The schedule tuple contains minutes, hours, days of month, months, and days of week. Registration also supports thresholds and placement options such as `free_space_extra`, `min_size`, `vacuum_max_age`, `max_retry`, `clustering_index`, relation/index tablespaces, and `skip_analyze`.

Start or stop scheduled processing as a superuser:

```sql
SELECT squeeze.start_worker();
SELECT squeeze.stop_worker();
```

For automatic startup, configure a worker role with superuser privileges:

```ini
squeeze.worker_autostart = 'my_database'
squeeze.worker_role = 'postgres'
```

### Version 1.9.6 and Operational Caveats

- Version 1.9.6 includes the 1.9.5 fixes for exclusion constraints and PostgreSQL 19 compatibility. Its configuration also accounts for `output_plugin_libraries`; check that setting when upgrading PostgreSQL.
- Version 1.9.4 fixes unsafe quoting in dynamically constructed `ANALYZE`, log, and error statements, including a superuser SQL-injection path. Upgrade earlier 1.9 builds promptly.
- A full-table squeeze needs free disk space of roughly twice the combined size of the target table and its indexes.
- Disruptive DDL, `VACUUM FULL`, `CLUSTER`, or `TRUNCATE` can make an in-progress squeeze abort. Coordinate schema changes and use `max_retry` deliberately.
- Like other online rewrite tools, `pg_squeeze` changes row visibility and has documented MVCC caveats for concurrent sessions that retain old snapshots.
- Configure `pg_squeeze` in `shared_preload_libraries` on the new cluster before `pg_upgrade` or dump/restore of a database containing the extension.
- On PostgreSQL versions before 19, keep `wal_level = logical`. The upstream PostgreSQL 19 configuration no longer requires that setting.
