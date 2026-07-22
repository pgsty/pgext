## Usage

Sources:

- [Official README for version 0.3.7](https://github.com/paradedb/pg_analytics/blob/v0.3.7/README.md)
- [Extension control file](https://github.com/paradedb/pg_analytics/blob/v0.3.7/pg_analytics.control)
- [Version 0.3.7 Rust manifest](https://github.com/paradedb/pg_analytics/blob/v0.3.7/Cargo.toml)
- [SQL-visible helper APIs](https://github.com/paradedb/pg_analytics/tree/v0.3.7/src/api)

`pg_analytics` 0.3.7 embeds DuckDB in PostgreSQL and exposes object-store files and data-lake formats through read-only FDWs. ParadeDB discontinued and archived this project in March 2025; it receives no maintenance and should be treated as a legacy integration.

### Core Workflow

For hook-based query pushdown in every backend, upstream recommends adding the library to `postgresql.conf` and restarting PostgreSQL before creating the extension:

```text
shared_preload_libraries = 'pg_analytics'
```

Create a format-specific wrapper, server, and foreign table. An empty column list asks the extension to discover the Parquet schema:

```sql
CREATE EXTENSION pg_analytics;

CREATE FOREIGN DATA WRAPPER parquet_wrapper
HANDLER parquet_fdw_handler
VALIDATOR parquet_fdw_validator;

CREATE SERVER parquet_server
FOREIGN DATA WRAPPER parquet_wrapper;

CREATE FOREIGN TABLE trips ()
SERVER parquet_server
OPTIONS (files 's3://paradedb-benchmarks/yellow_tripdata_2024-01.parquet');

SELECT count(*) FROM trips;
```

Credentials and object-store settings belong in a user mapping rather than the table definition. For S3-compatible storage, relevant mapping options include the credential values plus `type`, `region`, `endpoint`, `use_ssl`, and `url_style` as required by the endpoint.

### Main Objects

- FDW handler/validator pairs are provided for `csv`, `json`, `parquet`, `delta`, `iceberg`, and `spatial` sources.
- Foreign-table option `files` identifies local paths, HTTP URLs, object-store URLs, or format-specific table locations.
- `sniff_csv` reports detected CSV settings and columns.
- `parquet_describe` and `parquet_schema` inspect Parquet relations.
- `duckdb_execute` runs a statement on the embedded DuckDB connection; `duckdb_settings` and `duckdb_extensions` expose its state.

Version 0.3.7 embeds DuckDB 1.1.0 and advertises PostgreSQL 13+ support. Upstream lists Parquet, CSV, JSON, geospatial files, Delta Lake, and Apache Iceberg, with local filesystems, HTTP, S3-compatible stores, Google Cloud Storage, Azure storage, and Hugging Face as sources.

### Operational Boundaries

The FDWs are read-only. Filesystem access and outbound requests occur from the PostgreSQL server process, and DuckDB query execution shares backend resources with transactional workloads. Restrict function execution, user mappings, file paths, and remote endpoints; do not expose `duckdb_execute` to untrusted roles.

Adding `shared_preload_libraries` requires a restart. The archived project documents no Windows support and no future compatibility fixes. Pin the exact 0.3.7 build, test planner pushdown and type conversion on the target PostgreSQL release, and plan migration to a maintained design rather than adopting it for new production systems.
