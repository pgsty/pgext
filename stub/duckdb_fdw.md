## Usage

Sources:

- [2.0.2 README](https://github.com/alitrack/duckdb_fdw/blob/4c444d19f80eace50bc09fa3ec766e8ca2cf667d/README.md)
- [2.0.2 control file](https://github.com/alitrack/duckdb_fdw/blob/4c444d19f80eace50bc09fa3ec766e8ca2cf667d/duckdb_fdw.control)
- [2.0.2 SQL functions and privileges](https://github.com/alitrack/duckdb_fdw/blob/4c444d19f80eace50bc09fa3ec766e8ca2cf667d/duckdb_fdw--2.0.2.sql)
- [Server, mapping and table options](https://github.com/alitrack/duckdb_fdw/blob/4c444d19f80eace50bc09fa3ec766e8ca2cf667d/option.c)
- [Scan, write and diagnostic functions](https://github.com/alitrack/duckdb_fdw/blob/4c444d19f80eace50bc09fa3ec766e8ca2cf667d/duckdb_fdw.c)
- [Connection and transaction handling](https://github.com/alitrack/duckdb_fdw/blob/4c444d19f80eace50bc09fa3ec766e8ca2cf667d/connection.c)
- [Schema import implementation](https://github.com/alitrack/duckdb_fdw/blob/4c444d19f80eace50bc09fa3ec766e8ca2cf667d/import.c)
- [2.0.1 to 2.0.2 upgrade](https://github.com/alitrack/duckdb_fdw/blob/4c444d19f80eace50bc09fa3ec766e8ca2cf667d/duckdb_fdw--2.0.1--2.0.2.sql)

`duckdb_fdw` 2.0.2 exposes DuckDB tables and data sources as PostgreSQL foreign tables through DuckDB's native C API. Use it to query a DuckDB database, read Parquet, or access optional DuckDB catalogs. Foreign tables support reads and inserts; PostgreSQL UPDATE and DELETE through the FDW are not supported.

### Enable and Map a DuckDB Table

Install the extension as a superuser. A new installation does not require preloading or a PostgreSQL restart. Use a persistent DuckDB file in a directory accessible to the PostgreSQL service account; the example assumes that directory already exists.

```sql
CREATE EXTENSION duckdb_fdw;
CREATE SERVER duckdb_server
  FOREIGN DATA WRAPPER duckdb_fdw
  OPTIONS (database '/var/lib/postgresql/duckdb/analytics.duckdb');

SELECT duckdb_execute(
  'duckdb_server',
  'CREATE TABLE measurements (sensor_id INTEGER, reading DOUBLE)'
);

CREATE FOREIGN TABLE measurements (
  sensor_id integer,
  reading double precision
) SERVER duckdb_server OPTIONS (table 'measurements');

INSERT INTO measurements VALUES (1, 21.5), (2, 19.8);
SELECT * FROM measurements WHERE reading > 20;

SELECT extversion FROM pg_extension WHERE extname = 'duckdb_fdw';
SELECT duckdb_fdw_version();
```

The catalog query reports the extension version; despite its name, `duckdb_fdw_version()` reports the linked DuckDB engine version. A user mapping is optional for a local file without credentials. Database-file permissions are those of the PostgreSQL operating-system account.

### Import Tables and Read Parquet

```sql
CREATE SCHEMA duckdb_import;
IMPORT FOREIGN SCHEMA main
  FROM SERVER duckdb_server INTO duckdb_import;
```

Use DuckDB's schema name, normally `main`, rather than PostgreSQL's `public`. To read an existing Parquet file, declare matching PostgreSQL column types:

```sql
CREATE FOREIGN TABLE parquet_measurements (
  sensor_id integer,
  reading double precision
) SERVER duckdb_server
  OPTIONS (table '/var/lib/postgresql/duckdb/measurements.parquet');

SELECT * FROM parquet_measurements WHERE sensor_id = 1;
```

### Options and Privileges

| Object | Purpose |
| --- | --- |
| Server `database` | DuckDB file path, or a connection-local in-memory database |
| Server `force_readonly` | Open file databases read-only and reject FDW writes and non-read-only helper statements |
| Server `extensions`, `attach_catalogs` | Load optional DuckDB extensions and attach catalogs |
| Table `table` | Map a DuckDB table or a Parquet path |
| Column `column_name` | Map a PostgreSQL column to a differently named DuckDB column |
| `duckdb_execute(server, statement)` | Execute DuckDB SQL; returns void, not a result table |
| `duckdb_create_s3_secret(...)` | Create an S3 secret in the DuckDB connection |
| `duckdb_fdw_preflight()` | Inspect runtime compatibility before using DuckDB |

Grant application roles only the foreign-server, schema and foreign-table privileges they need. EXECUTE on the arbitrary-SQL and secret helpers is revoked from PUBLIC on installation; grant it only to trusted roles.

For S3, MotherDuck or Quack credentials, prefer user mappings. Server options are visible in PostgreSQL catalogs and are unsuitable for secrets. Remote examples additionally depend on the relevant DuckDB extensions, network access and credentials; their availability is not implied by installing the FDW.

### Connection and Transaction Boundaries

- Connections close at transaction end by default. With `database ':memory:'`, keep setup and queries in the same explicit transaction or the temporary database disappears.
- `duckdb_fdw.keep_connection` defaults to off. Enabling it retains connections after commit, but a writable file connection holds the DuckDB file lock across transactions and can block other PostgreSQL sessions from opening that file.
- Foreign-table inserts follow the top-level PostgreSQL commit or rollback. Savepoint rollback does not undo individual DuckDB writes: if the outer transaction commits, writes from an aborted subtransaction can remain.
- DuckDB does not participate in PostgreSQL two-phase commit. Do not use prepared transactions or assume crash-atomic commits across PostgreSQL and DuckDB. Arbitrary SQL sent through the helper also does not acquire PostgreSQL transaction semantics automatically.

### Compatibility and Upgrade

The Pigsty 2.0.2 packages target PostgreSQL 15-18. Upstream's README advertises PostgreSQL 13-18, which is broader than this package range. Version 2.0.2 accepts DuckDB 1.5.x and 2.0.x runtimes; it does not change the SQL function set from 2.0.1.

Runtime execution is guarded against loading `pg_duckdb` in the same backend. Keep the two extensions in separate PostgreSQL instances. When switching away from a preloaded `pg_duckdb`, remove its preload entry and restart PostgreSQL; replacing files does not unload libraries from existing processes.

```sql
SELECT duckdb_fdw_runtime_compatibility_status();
SELECT duckdb_fdw_runtime_fingerprint();
SELECT duckdb_fdw_preflight();
```

After installing matching files, use fresh backend sessions and update each database:

```sql
ALTER EXTENSION duckdb_fdw UPDATE TO '2.0.2';
```
