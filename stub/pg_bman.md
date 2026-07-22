## Usage

Sources:

- [Official upstream documentation](https://github.com/s-hironobu/pg_bman/blob/3d812f26abed778c44645259eaac2ede6e398a9d/README.md)
- [Official extension SQL](https://github.com/s-hironobu/pg_bman/blob/3d812f26abed778c44645259eaac2ede6e398a9d/pg_bman--1.0.sql)
- [Official extension implementation](https://github.com/s-hironobu/pg_bman/blob/3d812f26abed778c44645259eaac2ede6e398a9d/pg_bman.c)

`pg_bman` 1.0 is one component of a 2014-era remote backup toolkit. Its optional server extension lets a separate `pg_archivebackup` client list and fetch archived WAL segments over libpq; the `pg_bman` shell script coordinates full and incremental backup repositories. It is not a current PostgreSQL backup solution and should be used only to understand or recover a pinned legacy deployment.

### Extension API

The SQL extension installs two functions, revokes their default access from `PUBLIC`, and also enforces a C-level superuser check:

```sql
CREATE EXTENSION pg_bman;

SELECT *
FROM pg_show_archives('/var/lib/postgresql/data/archives');

SELECT octet_length(
  pg_get_archive(
    '/var/lib/postgresql/data/archives/000000010000000000000001'
  )
);
```

`pg_show_archives(text)` returns names that look like 24-hex-digit WAL segments. `pg_get_archive(text)` reads up to one WAL segment into a `bytea`. These are internal transport primitives, not backup-consistency or restore-validation APIs.

### Legacy Backup Flow

Upstream configures PostgreSQL archiving, creates a backup-server repository, and points the external scripts at `pg_basebackup`, the archive directory, and a libpq connection. The command surface includes `pg_bman BACKUP FULL`, `pg_bman BACKUP INCREMENTAL`, `pg_bman SHOW`, and `pg_bman RESTORE`. Restore output must then be copied back to a stopped PostgreSQL data directory by external tools.

### Do Not Use Unverified on Modern PostgreSQL

The documented configuration targets PostgreSQL 9.3-era WAL and recovery behavior, including `recovery.conf` and manual archive copying. Modern PostgreSQL recovery signaling, WAL layout, permissions, checksums, timelines, tablespaces, and backup manifests differ materially. The source's filename checks are also old and should not be treated as a general secure file API even though calls require superuser.

For any retained legacy deployment, isolate the backup role and network, keep `PUBLIC` revoked, validate repository permissions, and perform full restore drills to a separate host. A successful `BACKUP` command or WAL read is not proof of a restorable cluster. For supported PostgreSQL, select a maintained physical-backup tool with current-version restore tests, retention, encryption, timeline handling, and integrity validation.
