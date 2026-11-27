## Usage

Sources:

- [Official source](https://api.pgxn.org/src/pg_repack_safe/pg_repack_safe-1.0.0/README.md)
- [Official source](https://api.pgxn.org/src/pg_repack_safe/pg_repack_safe-1.0.0/pg_repack_safe.control)
- [Official source](https://api.pgxn.org/src/pg_repack_safe/pg_repack_safe-1.0.0/pg_repack_safe--1.0.sql)
- [Official source](https://api.pgxn.org/src/pg_repack_safe/pg_repack_safe-1.0.0/doc/installation.md)
- [Official source](https://api.pgxn.org/src/pg_repack_safe/pg_repack_safe-1.0.0/doc/limits.md)
- [Official source](https://api.pgxn.org/src/pg_repack_safe/pg_repack_safe-1.0.0/doc/removal.md)

`pg_repack_safe` 1.0.0 guards PostgreSQL 19's built-in concurrent REPACK against reads using snapshots predating the rewrite. Instead of silently seeing an empty rewritten table, affected transactions receive SQLSTATE 40001 and must retry. It does not guard the separate pg_repack extension or every table-rewrite command.

### Enable and Check

Install the matching library on the primary and every standby. Add it to the existing preload list, restart standbys first, then the primary:

```ini
shared_preload_libraries = 'pg_repack_safe'
```

As a superuser, enable it in each protected database:

```sql
CREATE EXTENSION pg_repack_safe;
SELECT repack_safe.status();
```

The result must be true. Merely creating the SQL extension without preload provides no protection. The distribution/library version is 1.0.0; the intentional SQL extension version is 1.0. It requires `plpgsql`, is not trusted or relocatable, and uses the fixed `repack_safe` schema.

### Operational Boundaries

The extension writes custom WAL using resource manager 149. Every replaying server needs a compatible library; missing it can stop replay. Keep the same library major on primary and standbys, and include the library and configuration in recovery planning. Applications must retry the whole failed transaction.

Some reads outside the executor, including amcheck and pgrowlocks, are outside the guard. A foreign-key check can fail with SQLSTATE 23503 instead of the retry code, and upstream documents a standby replay gap. Treat these as limits, not a general MVCC repair.

### Removal

Do not simply remove the preload library while protected rewrite records remain. Stop concurrent repacking, follow the superuser pruning and VACUUM TRUNCATE procedure in the removal guide, and verify the relevant relation storage reaches zero before dropping the extension and retiring the library. Coordinate this with standby replay and backups.
