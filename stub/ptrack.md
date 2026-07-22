## Usage

Sources:

- [Pinned official documentation](https://github.com/postgrespro/ptrack/blob/62bd42d7308ec7ed786abf2ddc1af24fc3dd72dd/README.md)
- [Pinned control file](https://github.com/postgrespro/ptrack/blob/62bd42d7308ec7ed786abf2ddc1af24fc3dd72dd/ptrack.control)

`ptrack` is a block-level change-tracking engine for incremental PostgreSQL backups, primarily consumed by `pg_probackup`. It marks changed relation pages in a map so a backup tool can copy only pages changed since an LSN. It may report false positives, but is designed not to miss changed pages except hint-bit-only changes.

### Prerequisites and Configuration

Installing the extension package alone is insufficient. Version 2.4 requires the matching `ptrack` patch applied to the PostgreSQL core build; the upstream repository provides separate patches for PostgreSQL 11–18. The patched server utilities know how to handle `ptrack.map.*` service files.

Configure the library and a nonzero map size, then restart and create the extension:

```ini
shared_preload_libraries = 'ptrack'
ptrack.map_size = 64
wal_level = 'replica'
```

```sql
CREATE EXTENSION ptrack;

SELECT ptrack_version();
SELECT ptrack_init_lsn();
```

`ptrack.map_size` is measured in MiB. The upstream sizing rule is approximately 1 MiB per 1 GiB of expected `PGDATA` size. The default `0` disables tracking and cleans up service files.

### Backup-Facing API

```sql
SELECT *
FROM ptrack_get_change_stat('0/285C8C8'::pg_lsn);

SELECT path, pagecount, pagemap
FROM ptrack_get_pagemapset('0/185C8C0'::pg_lsn);
```

- `ptrack_version()` returns the extension version.
- `ptrack_init_lsn()` returns the LSN of the last map initialization.
- `ptrack_get_change_stat(pg_lsn)` returns changed-file, page, and size totals since an LSN.
- `ptrack_get_pagemapset(pg_lsn)` returns each changed data-file path, block count, and bitmap for the backup tool.

### Operational Boundaries

Use `wal_level >= replica`; with `minimal`, crash recovery can miss changes made by operations that did not emit WAL. The upstream documentation identifies `pg_probackup` as the production-ready backup utility that fully supports this engine.

The map size cannot change at runtime. Resizing it on restart loses accumulated tracking state and should be paired with a new full backup. Allow up to twice `ptrack.map_size` in extra disk space because checkpoint durability uses a temporary map file. Setting the map too small increases false positives.

Upgrades can invalidate or require removal of older map files, especially across the 2.1→2.2 and 2.2→2.3 transitions. Stop the server, follow the release-specific upgrade sequence, upgrade the PostgreSQL core patch and extension binaries together, restart, then run `ALTER EXTENSION ptrack UPDATE`. Prove backup and restore behavior on the exact build before relying on it for recovery.
