## Usage

Sources:

- [Official pg_hardstorage.control](https://github.com/cybertec-postgresql/pg_hardstorage/blob/b47541b7e1cea69ce6ec63b26e154eb25fc4ca91/ext/pg_hardstorage_extension/pg_hardstorage.control)
- [Official pg_hardstorage--1.0.sql](https://github.com/cybertec-postgresql/pg_hardstorage/blob/b47541b7e1cea69ce6ec63b26e154eb25fc4ca91/ext/pg_hardstorage_extension/pg_hardstorage--1.0.sql)
- [Official pg_hardstorage_db_install-extension.md](https://github.com/cybertec-postgresql/pg_hardstorage/blob/b47541b7e1cea69ce6ec63b26e154eb25fc4ca91/docs/reference/cli/pg_hardstorage_db_install-extension.md)

`pg_hardstorage` SQL extension 1.0 exposes backup status through tables, views and writer functions. It is the database integration surface of the separate Go backup CLI; creating this extension does not start backups or WAL streaming.

### Enable and Inspect

```sql
CREATE EXTENSION pg_hardstorage;
SELECT * FROM pg_hardstorage.backups;
SELECT * FROM pg_hardstorage.health;
SELECT * FROM pg_hardstorage.rpo;
```

### Writer and Reader Roles

With extension files installed, a superuser creates the extension in its fixed schema. Installation also creates the cluster-wide `pg_hardstorage_writer` NOLOGIN role if absent. Readers receive schema usage and SELECT on the three public views. `pg_hardstorage.upsert_backup()`, `pg_hardstorage.upsert_health()` and `pg_hardstorage.upsert_rpo()` are SECURITY DEFINER functions with a fixed search path; execution and direct state-table writes are restricted to the writer role. Grant membership only to the trusted state publisher.

### Population and Maintenance

In this source revision no built-in backup data path calls the upsert functions, so the views remain empty until an operator publishes state. The CLI also exposes `pg_hardstorage db install-extension` with dry-run and SQL-printing options; that path installs its bundled SQL surface and should not be confused with a registered extension upgrade. The SQL extension uses PL/pgSQL during installation, needs no PostgreSQL shared library or preload, and has no independently documented PostgreSQL-major support matrix. CLI release numbers differ from SQL version 1.0.
