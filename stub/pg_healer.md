## Usage

Sources:

- [Official README](https://github.com/turnstep/pg_healer/blob/d313cccc2944a3c4a54757c5fa2693e9ae21d637/README.md)
- [Official control file](https://github.com/turnstep/pg_healer/blob/d313cccc2944a3c4a54757c5fa2693e9ae21d637/pg_healer.control)
- [Official installation SQL](https://github.com/turnstep/pg_healer/blob/d313cccc2944a3c4a54757c5fa2693e9ae21d637/pg_healer--1.0.sql)
- [Official implementation](https://github.com/turnstep/pg_healer/blob/d313cccc2944a3c4a54757c5fa2693e9ae21d637/pg_healer.c)
- [Author's design article](https://www.endpointdev.com/blog/2016/09/pghealer-repairing-postgres-problems/)

`pg_healer` is an abandoned proof of concept that hooks PostgreSQL errors and attempts to repair damaged heap pages by copying a block from a filesystem backup. It modifies relation files outside PostgreSQL's buffer and WAL machinery, so use it only while studying corruption on a disposable clone—not as a recovery tool.

### Core Workflow

The library must be preloaded. Install the untrusted `plperlu` language explicitly because the control file does not declare it as a dependency, then install the extension after restarting PostgreSQL.

```conf
shared_preload_libraries = 'pg_healer'
```

```sql
CREATE EXTENSION plperlu;
CREATE EXTENSION pg_healer;
REVOKE ALL ON FUNCTION pg_healer_cauldron() FROM PUBLIC;

SELECT * FROM pg_healer_config;
SELECT * FROM pg_healer_log ORDER BY happened DESC;
```

The error hook watches for an internal-data-corruption error, consults `pg_healer_config`, and tries to replace the reported block from a copy beneath `PGDATA/pg_healer`. The intended backup entry point is `pg_healer_cauldron()`, but it performs a raw filesystem copy and does not create an online-consistent PostgreSQL backup.

### Objects

- `pg_healer_config` enables or disables attempted repairs per relation.
- `pg_healer_log` records repair attempts.
- `pg_healer_cauldron()` invokes the Perl backup helper.
- `pg_healer_remove_from_buffer(regclass)` tries to evict a relation from shared buffers and requires superuser privileges in the C implementation.
- `pg_healer_corrupt(regclass, text)` is a deliberate corruption helper and must never be used on valuable data.

### Operational Boundaries

The reviewed version `1.0` is unsafe even as an automatic repair prototype. Its SQL declaration gives `pg_healer_corrupt` two arguments while the C function reads a third, and its block-number parser starts at the word `block` rather than at the following number. The backup helper's checksum validation is unfinished, while direct file writes are not coordinated with PostgreSQL buffers, checkpoints, WAL, or concurrent access.

Keep the preload setting and all extension objects off production systems. For real corruption, stop writes, preserve the original files, take a verified physical copy, and use supported recovery, backup, or specialist forensic procedures.
