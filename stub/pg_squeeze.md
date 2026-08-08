## Usage

Sources:

- [pg_squeeze REL1_9_4 release](https://github.com/cybertec-postgresql/pg_squeeze/releases/tag/REL1_9_4)
- [pg_squeeze REL1_9_4 README](https://github.com/cybertec-postgresql/pg_squeeze/blob/REL1_9_4/README.md)
- [pg_squeeze release notes](https://github.com/cybertec-postgresql/pg_squeeze/blob/REL1_9_4/NEWS)

`pg_squeeze` removes bloat from a table and its indexes while allowing concurrent reads and writes. It copies live tuples to new storage and applies concurrent changes through logical decoding, avoiding the long exclusive lock of `VACUUM FULL`. Use it only after sizing replication slots, disk space, and the table's replica identity.

### Configure and Install

```conf
max_replication_slots = 1  # or add one to the existing requirement
shared_preload_libraries = 'pg_squeeze'
wal_level = logical       # required on PostgreSQL versions before 19
```

Restart PostgreSQL, then create the extension:

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

For automatic startup:

```conf
squeeze.worker_autostart = 'my_database'
squeeze.worker_role = 'postgres'
```

### Version 1.9.4 and Operational Caveats

- Version 1.9.4 fixes unsafe quoting in dynamically constructed `ANALYZE`, log, and error statements, including a superuser SQL-injection path. Upgrade earlier 1.9 builds promptly.
- A full-table squeeze needs free disk space of roughly twice the combined size of the target table and its indexes.
- Disruptive DDL, `VACUUM FULL`, `CLUSTER`, or `TRUNCATE` can make an in-progress squeeze abort. Coordinate schema changes and use `max_retry` deliberately.
- Like other online rewrite tools, `pg_squeeze` changes row visibility and has documented MVCC caveats for concurrent sessions that retain old snapshots.
- Configure `pg_squeeze` in `shared_preload_libraries` on the new cluster before `pg_upgrade` or dump/restore of a database containing the extension.
- Current Pigsty packages cover PostgreSQL 14-18. For those versions, keep `wal_level = logical`; upstream's relaxed PostgreSQL 19 rule does not apply to this package matrix yet.
