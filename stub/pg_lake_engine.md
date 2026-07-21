## Usage

Sources:

- [Official pg_lake architecture overview](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/README.md#architecture)
- [Version 3.4 control file](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_lake_engine/pg_lake_engine.control)
- [Base SQL objects](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_lake_engine/pg_lake_engine--3.0.sql)
- [Version 3.4 cleanup-queue change](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_lake_engine/pg_lake_engine--3.3--3.4.sql)

`pg_lake_engine` is the shared execution layer used by the pg_lake table, copy, and Iceberg extensions. It rewrites eligible PostgreSQL work for `pgduck_server`, maps PostgreSQL and DuckDB values, and tracks remote files that must be removed after aborts or table changes. It is an internal dependency rather than a standalone analytics interface.

### Deployment Boundary

Install it through the top-level extension so its dependency graph and preload directives remain aligned:

```conf
shared_preload_libraries = 'pg_extension_base'
```

```sql
CREATE EXTENSION pg_lake CASCADE;
```

`pg_lake_engine` requires `pg_extension_base` and `pg_map`. Query execution also requires a running local `pgduck_server`; creating only this extension does not provide a usable lake table.

### User-Visible Objects

- Roles `lake_read`, `lake_write`, and `lake_read_write`: shared privilege groups consumed by the other components.
- `to_postgres(any)`: returns its input while forcing that expression to be evaluated in PostgreSQL instead of pushed to the lake engine.
- `to_date(double precision)`: converts a days-since-Unix-epoch value commonly found in Parquet to a PostgreSQL `date`.
- `lake_engine.deletion_queue`: tracks committed orphan-file cleanup; readable by `lake_write`.
- `lake_engine.in_progress_files`: tracks files produced by transactions that have not committed.
- `lake_engine.flush_deletion_queue(regclass)` and `flush_in_progress_queue()`: privileged cleanup functions used by maintenance paths.

```sql
SELECT to_postgres(application_only_function(payload))
FROM external_events;
```

Use `to_postgres()` only when an expression cannot or should not be pushed down; pulling data back into PostgreSQL may substantially increase transfer and execution cost.

### Internal State and Caveats

- The `__lake__internal__nsp__` functions are planner/deparser placeholders and are not a supported direct SQL API.
- Do not manually update or delete queue rows. Cleanup functions need the extension's object-store credentials and privilege roles and should be invoked only as documented by operational tooling.
- Version `3.4` adds `resolve_metadata` to the deletion queue so Iceberg metadata can be expanded into exact referenced files during `VACUUM`, moving object-store traversal off the `DROP` path.
- Roles are cluster-wide objects and can outlive an extension instance in one database; review memberships separately when removing pg_lake.
