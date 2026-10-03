## Usage

Sources:

- [Official README.md](https://github.com/GerardSmit/pgBackupDatabase/blob/e1bbaf1cb09657986e00bcfe350601bad25b9141/README.md)
- [Official pg_dbbackup.control](https://github.com/GerardSmit/pgBackupDatabase/blob/e1bbaf1cb09657986e00bcfe350601bad25b9141/pg_dbbackup.control)
- [Official pg_dbbackup--0.0.1.sql](https://github.com/GerardSmit/pgBackupDatabase/blob/e1bbaf1cb09657986e00bcfe350601bad25b9141/sql/pg_dbbackup--0.0.1.sql)

`pg_dbbackup` 0.0.1 provides per-database logical backups and recovery chains on PostgreSQL 17 and 18. It requires administrator installation, the fixed `dbbackup` schema, shared preload and logical WAL. Merge the preload setting with existing entries and restart before creating the extension.

### Create a Local Backup

```conf
shared_preload_libraries = 'pg_dbbackup'
wal_level = logical
```

```sql
CREATE EXTENSION pg_dbbackup;
SELECT dbbackup.pg_dbbackup_get_mode('app');
SELECT dbbackup.pg_dbbackup('app', '/var/backups/app-full.bak',
  type := 'full', compress := true);
SELECT * FROM dbbackup.pg_dbbackup_header('/var/backups/app-full.bak');
SELECT * FROM dbbackup.pg_dbbackup_verify('/var/backups/app-full.bak');
```

### Recovery Models and Storage

SIMPLE is the default recovery model: full backups plus cumulative differentials. FULL mode adds logical decoding and a DDL journal for log backups and point-in-time recovery; `dbbackup.pg_dbbackup_set_mode()` selects it. Paths are on the server filesystem, not the client. S3 targets use `dbbackup.create_s3_target()` and `dbbackup.pg_dbbackup_to_storage()`; the target stores non-secret configuration while AWS credentials come from the server environment. Local-file backups are rejected on hot standbys.

### Scheduling, Observation and Recovery

`dbbackup.create_schedule()` schedules a backup set. `dbbackup.pg_dbbackup_async()` queues work; `dbbackup.pg_dbbackup_status()` and `dbbackup.pg_dbbackup_wait()` report progress. Inspect jobs, artifacts and events through the corresponding tables in the extension schema. `dbbackup.pg_dbbackup_retention_plan()` previews expiry before applying retention. `dbbackup.pg_dbrestore()` restores a file chain; `dbbackup.pg_dbrestore_at()` restores to a selected time. Restore drops an existing target database after staging the replacement, so select and protect the target deliberately.

### Limits

This does not replace cluster-level physical recovery. FULL backup takes table SHARE locks; logical slots can retain WAL indefinitely if neglected. Keep replication-slot capacity and retention healthy, and configure failover logical-slot synchronization for HA. Protect backup files, passwords and cloud credentials; test recovery on isolated targets. The C library uses zstd, OpenSSL and libcurl; SQL helpers require PL/pgSQL.
