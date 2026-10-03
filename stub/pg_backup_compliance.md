## Usage

Sources:

- [README.md](https://github.com/klouddb/pg_backup_compliance/blob/40876bd5d60c33dfc3451375cb472bba516e5cf1/README.md)
- [pg_backup_compliance.control](https://github.com/klouddb/pg_backup_compliance/blob/40876bd5d60c33dfc3451375cb472bba516e5cf1/pg_backup_compliance.control)
- [sql/pg_backup_compliance--1.0.sql](https://github.com/klouddb/pg_backup_compliance/blob/40876bd5d60c33dfc3451375cb472bba516e5cf1/sql/pg_backup_compliance--1.0.sql)
- [pg_backup_compliance.c](https://github.com/klouddb/pg_backup_compliance/blob/40876bd5d60c33dfc3451375cb472bba516e5cf1/pg_backup_compliance.c)
- [pg_backup_compliance_capture.c](https://github.com/klouddb/pg_backup_compliance/blob/40876bd5d60c33dfc3451375cb472bba516e5cf1/pg_backup_compliance_capture.c)

`pg_backup_compliance` 1.0 observes backup-related sessions and exposes activity through SQL views. It monitors activity; a recorded success does not prove that the resulting backup can be restored.

### Enablement

Merge the library into the existing preload list and restart, then create the SQL objects as a superuser in each database used for inspection. The README states PostgreSQL 13+ support, but the pinned capture hook uses the PostgreSQL 14+ ProcessUtility signature without a PG13 branch; do not assume PG13 compatibility.

```ini
shared_preload_libraries = 'pg_backup_compliance'
```

```sql
CREATE EXTENSION pg_backup_compliance;
SELECT application_name, backup_type, status, start_time, end_time
FROM pg_backup_compliance ORDER BY start_time DESC LIMIT 20;
SELECT * FROM pg_backup_compliance_failed;
```

### Objects and Configuration

`pg_backup_compliance` unions live and archived operation rows in the reviewed SQL. `pg_backup_compliance_running`, `pg_backup_compliance_failed`, `pg_backup_compliance_last_24h` and the monthly/quarterly views filter those records. The views grant access to `pg_monitor`. `pg_backup_compliance_info()` reports cache counters; `pg_backup_compliance_reset()` is restricted to the superuser.

The main settings are `pg_backup_compliance.enabled`, `pg_backup_compliance.save`, `pg_backup_compliance.max_entries` and `pg_backup_compliance.track_apps`; the last matches application-name prefixes.

### Limitations

A standalone `pg_dump` from the same host and user can be merged into an overlapping `pg_dumpall` record. Some pgBackRest failures are not classified accurately. Application names are client-supplied, and cache eviction, persistence and archive retention must be evaluated for the deployment; the README and SQL describe history differently. Do not use the views as a complete retention or restore guarantee. Upstream provides copyright notices but no explicit license grant was found.
