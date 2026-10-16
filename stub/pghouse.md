## Usage

Sources:

- [README.md](https://github.com/myyrakle/pghouse/blob/a8f8b5e553586b29426d1eff1f1369e51eb3bb6d/README.md)
- [src/pg/am.rs](https://github.com/myyrakle/pghouse/blob/a8f8b5e553586b29426d1eff1f1369e51eb3bb6d/src/pg/am.rs)
- [src/pg/catalog.rs](https://github.com/myyrakle/pghouse/blob/a8f8b5e553586b29426d1eff1f1369e51eb3bb6d/src/pg/catalog.rs)
- [src/pg/api.rs](https://github.com/myyrakle/pghouse/blob/a8f8b5e553586b29426d1eff1f1369e51eb3bb6d/src/pg/api.rs)
- [src/pg/worker.rs](https://github.com/myyrakle/pghouse/blob/a8f8b5e553586b29426d1eff1f1369e51eb3bb6d/src/pg/worker.rs)
- [Cargo.toml](https://github.com/myyrakle/pghouse/blob/a8f8b5e553586b29426d1eff1f1369e51eb3bb6d/Cargo.toml)
- [pghouse.control](https://github.com/myyrakle/pghouse/blob/a8f8b5e553586b29426d1eff1f1369e51eb3bb6d/pghouse.control)

`pghouse` 0.1.0 is a PostgreSQL 18 prototype combining a heap-backed table access method with asynchronous columnar sidecar files. Ordinary SQL still reads the delegated heap; explicit scan helpers read the sidecar. It is not a completed durable columnar engine.

### Core Workflow

```sql
CREATE EXTENSION pghouse;
CREATE TABLE pghouse_demo (id bigint PRIMARY KEY, value text) USING pghouse;
SELECT pghouse_register_table('public.pghouse_demo', 'id');
INSERT INTO pghouse_demo VALUES (1, 'hello');
SELECT pghouse_run_maintenance_once();
SELECT pghouse_explain_scan('public.pghouse_demo');
SELECT pghouse_scan_rows('public.pghouse_demo');
```

### Operational Boundaries

The control requires superuser installation and forbids relocation; the extension creates the `pghouse` schema. No shared preload is required. `pghouse_register_table` installs DML capture for a table and its primary-key column; it supports granule size, compression, buffer thresholds and a storage path. `pghouse_run_maintenance_once` drains pending work; `pghouse_launch_worker` launches an optional dynamic worker in the current database.

`pghouse_scan_rows` returns sidecar rows as JSONB with optional projection, key bounds, limit and snapshot generation. `pghouse_explain_scan` reports its scan plan, and `pghouse_schedule_merge` queues compaction. Sidecar visibility depends on maintenance; these helpers are not interchangeable with ordinary transactional table reads.

The default file root is `$PGDATA/pghouse/db_<db_oid>/rel_<table_oid>/`. Granules are not WAL-aware; do not assume PostgreSQL replication, crash recovery or normal database backups protect these files. Keep authoritative data in the heap and use an isolated evaluation database. The source declares pgrx 0.17 and only a PostgreSQL 18 feature; no Pigsty package is implied.
