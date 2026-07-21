## Usage

Sources:

- [Official pg_extension_base README](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_extension_base/README.md)
- [Version 3.4 control file](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_extension_base/pg_extension_base.control)
- [SQL API definition](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_extension_base/pg_extension_base--1.6.sql)
- [pg_lake build and preload guide](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/docs/building-from-source.md)

`pg_extension_base` is Snowflake's infrastructure extension for other PostgreSQL extensions. It provides control-file-driven library preloading, database-scoped lifecycle workers, dependency-aware updates, and short-lived attached workers. Application users normally install it as a dependency of `pg_lake`; extension developers use its C and SQL APIs directly.

### Enable the Infrastructure

`pg_extension_base` must be preloaded before extensions can use its startup machinery:

```conf
shared_preload_libraries = 'pg_extension_base'
```

Restart PostgreSQL, then create it in a database that hosts dependent extensions:

```sql
CREATE EXTENSION pg_extension_base;

SELECT * FROM extension_base.list_preload_libraries();
SELECT * FROM extension_base.list_base_workers();
```

### Extension-Developer Workflow

A dependent extension can request startup loading with a control-file directive:

```text
requires = 'pg_extension_base'
module_pathname = '$libdir/my_extension'
#!shared_preload_libraries
```

To attach a database lifecycle worker, define an internal C function and register it during extension installation:

```sql
CREATE FUNCTION my_extension.main_worker(internal)
RETURNS internal
LANGUAGE C
AS 'MODULE_PATHNAME', 'my_extension_main_worker';

SELECT extension_base.register_worker(
    'my_extension_main',
    'my_extension.main_worker'
);
```

The base infrastructure starts the worker after server startup, `CREATE EXTENSION`, or creation of a database from a template, and attempts to stop it for `DROP EXTENSION` or `DROP DATABASE`.

### SQL API Index

- `extension_base.list_preload_libraries()`: reports extension/library pairs discovered for startup loading.
- `extension_base.register_worker(name, regproc)` and `deregister_worker(name)`: manage lifecycle-worker registrations; public execution is revoked.
- `extension_base.list_base_workers()`: reports database, extension, PID, and restart state for base workers.
- `extension_base.list_database_starters()`: reports per-database starter processes.
- `extension_base.run_attached(command, dbname)`: runs a command in a short-lived worker, optionally in another database, and returns command IDs and tags.

### Operational Boundaries

- Lifecycle worker functions run as superuser and outside a transaction. They must start their own transactions, check interrupts, and avoid trusting user-controlled SQL.
- Worker termination around failed `DROP` operations is best effort; worker code must tolerate the extension briefly disappearing or a stop being reversed.
- `run_attached` is for bounded work that may commit independently of the caller. It is not a detached job queue for long-running tasks.
- Version `3.4` changes no user-facing SQL objects relative to `3.3`; its upgrade script is intentionally empty.
