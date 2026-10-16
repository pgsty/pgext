## Usage

Sources:

- [README.md](https://github.com/12vishnu10/pg_creation_tracker/blob/1e5dc91a17b8802fabd58b127fb30b8a5d93462d/README.md)
- [pg_creation_tracker--1.0.sql](https://github.com/12vishnu10/pg_creation_tracker/blob/1e5dc91a17b8802fabd58b127fb30b8a5d93462d/pg_creation_tracker--1.0.sql)
- [pg_creation_tracker.control](https://github.com/12vishnu10/pg_creation_tracker/blob/1e5dc91a17b8802fabd58b127fb30b8a5d93462d/pg_creation_tracker.control)

`pg_creation_tracker` 1.0 is a SQL/PLpgSQL extension that records table DDL and estimates timestamps for existing tables and databases. The GitHub release is labelled v1.0.0; the installed control and SQL version is 1.0.

### Core Workflow

```sql
CREATE EXTENSION pg_creation_tracker;
SELECT schema_name, object_name, creation_time, recorded_time, detection_method
FROM timestamp_tracking.object_creation_log
ORDER BY id;
SELECT database_name, creation_time, detection_method
FROM timestamp_tracking.database_creation_info;
```

### Operational Boundaries

Installation requires a superuser and creates the fixed `timestamp_tracking` schema, an event trigger and an initial backfill. No native library or preload is needed. Install separately in each database that needs tracking; installation in a template affects subsequently copied databases.

`object_creation_log` exposes `creation_time`, `recorded_time` and `detection_method`. The `ddl_command_end` trigger records table and foreign-table DDL with `statement_timestamp()` and keeps the first row for each object OID; it is not a complete immutable audit trail. `backfill_existing_tables()` attempts catalog commit timestamps and then file metadata. `database_creation_info` derives cluster database timestamps from directory metadata.

Historical values are estimates: the SQL reads file modification times, not birth times; catalog `xmin` may describe a later change, and partition roots use the earliest child modification time. A missing commit timestamp or unreadable path can yield NULL. Do not treat those values as proof of original creation time. The helper functions are `SECURITY DEFINER`; review execution grants and search-path security before deployment. Upstream supplies no complete supported-major matrix or explicit license file in the inspected release.
