## Usage

Sources:

- [3.7.0 pg_lake_engine.control](https://github.com/Snowflake-Labs/pg_lake/blob/v3.7.0/pg_lake_engine/pg_lake_engine.control)
- [3.7.0 v3.7.0](https://github.com/Snowflake-Labs/pg_lake/releases/tag/v3.7.0)

- [Upstream source](https://github.com/Snowflake-Labs/pg_lake/releases/tag/v3.6.2)

- [Current upstream](https://github.com/Snowflake-Labs/pg_lake/releases/tag/v3.6.1)
- [Current upstream](https://github.com/Snowflake-Labs/pg_lake/blob/v3.6.1/docs/table-options.md)
- [Current upstream](https://github.com/Snowflake-Labs/pg_lake/blob/v3.6.1/docs/iceberg-catalogs.md)

- [Official pg_lake architecture overview](https://github.com/Snowflake-Labs/pg_lake/blob/v3.6.0/README.md#architecture)
- [Version 3.6 control file](https://github.com/Snowflake-Labs/pg_lake/blob/v3.6.0/pg_lake_engine/pg_lake_engine.control)
- [Base SQL objects](https://github.com/Snowflake-Labs/pg_lake/blob/v3.6.0/pg_lake_engine/pg_lake_engine--3.0.sql)
- [3.5 cleanup retry-state migration](https://github.com/Snowflake-Labs/pg_lake/blob/v3.6.0/pg_lake_engine/pg_lake_engine--3.4--3.5.sql)

- [Upstream 3.6.0](https://github.com/Snowflake-Labs/pg_lake/releases/tag/v3.6.0)
- [3.5 to 3.6 SQL](https://github.com/Snowflake-Labs/pg_lake/blob/v3.6.0/pg_lake_engine/pg_lake_engine--3.5--3.6.sql)

`pg_lake_engine` is the shared execution layer used by the pg_lake table, copy, and Iceberg extensions. It rewrites eligible PostgreSQL work for `pgduck_server`, maps PostgreSQL and DuckDB values, and tracks remote files that must be removed after aborts or table changes. It is an internal dependency rather than a standalone analytics interface.

The reviewed upstream release is `3.6.0`, with SQL extension version `3.6` and PostgreSQL 16–19 support. Existing packaged versions are tracked separately. Keep libraries and the query server from the same release.

### Deployment Boundary

Install it through the top-level extension so its dependency graph and preload directives remain aligned:

```ini
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
- Version `3.5` adds `last_attempt_at` to `lake_engine.deletion_queue`. Cleanup retries are spaced by `pg_lake_engine.vacuum_file_remove_retry_interval`; `pg_lake_engine.vacuum_file_remove_max_retries` limits attempts instead of counting unrelated VACUUM passes.
- Roles are cluster-wide objects and can outlive an extension instance in one database; review memberships separately when removing pg_lake.

### Upstream 3.6.0

Upstream 3.6.0 adds PostgreSQL 19 support and uses DuckDB 1.5.6. Deploy matching libraries and `pgduck_server`, keep `pg_extension_base` preloaded, restart loaded processes and apply the supplied SQL updates to version 3.6. The engine migration adds internal UUID pushdown wrappers. They are planner/deparser plumbing, not a general-purpose application API.

### Upstream 3.6.1 runtime update

Deploy the matching 3.6.1 `pgduck_server` with the extension libraries. This patch backports a Brotli decoder security correction and reduces the query server's S3 uploader maximum file size to 50 GB. SQL remains 3.6; an unchanged SQL version does not prove the patched query server is installed. The C COPY-query export interface is internal integration work, not a new SQL function in this component.

### Upstream 3.6.2

The 3.6.2 patch corrects byte-by-byte CSV-value transmission in pgduck_server. Deploy the matching query-server binary with the extension libraries and reconnect/restart processes holding old code. The component SQL version remains 3.6; an unchanged extversion does not prove the query-server fix is installed. Package availability remains a separate catalog fact.

### Upstream 3.7.0

The current reviewed release is 3.7.0 and this component's control version is 3.7. Earlier version sections describe historical behavior and package examples. Upgrade the complete family with matching libraries and `pgduck_server`, restart processes holding old code, and apply the supplied extension updates to SQL version 3.7. This does not establish package availability.
