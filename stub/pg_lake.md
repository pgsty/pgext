## Usage

Sources:

- [Official pg_lake README](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/README.md)
- [Version 3.4 control file](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_lake/pg_lake.control)
- [Official build and startup guide](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/docs/building-from-source.md)
- [Official project documentation index](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/docs/README.md)

`pg_lake` is the top-level extension for Snowflake's PostgreSQL lakehouse stack. It installs the table, Iceberg, copy, query-engine, extension-base, and map components needed to query object-store files and create transactional Iceberg tables. The PostgreSQL extensions orchestrate planning and transactions while a separate local `pgduck_server` process executes vectorized work with DuckDB.

### Start the Stack

Version `3.4` supports PostgreSQL 16 through 18. Preload the common extension infrastructure, restart PostgreSQL, and start `pgduck_server` on the database host:

```conf
shared_preload_libraries = 'pg_extension_base'
```

```shell
pgduck_server --cache_dir /var/cache/pg_lake
```

Create the complete dependency tree in the target database:

```sql
CREATE EXTENSION pg_lake CASCADE;
SELECT lake.version();
```

Configure object-store credentials for `pgduck_server`, then choose the managed Iceberg location:

```sql
SET pg_lake_iceberg.default_location_prefix =
    's3://analytics-bucket/warehouse';
```

### Core Workflows

Create and modify a transactional Iceberg table:

```sql
CREATE TABLE measurements (
    station_name text NOT NULL,
    measured_at timestamptz NOT NULL,
    value double precision
) USING iceberg;

INSERT INTO measurements VALUES
    ('Istanbul', now(), 18.5),
    ('Haarlem', now(), 9.3);
```

Import or export Parquet, CSV, or newline-delimited JSON through `COPY`:

```sql
COPY (SELECT * FROM measurements)
TO 's3://analytics-bucket/export/measurements.parquet';

COPY measurements
FROM 's3://analytics-bucket/import/measurements.parquet';
```

Query files without loading them into PostgreSQL:

```sql
CREATE FOREIGN TABLE external_events ()
SERVER pg_lake
OPTIONS (path 's3://analytics-bucket/events/*.parquet');

SELECT count(*) FROM external_events;
```

### Component Index

- `pg_lake`: meta-extension and `lake.version()`.
- `pg_lake_table`: data-lake FDW, Iceberg table syntax, file utilities, and table catalogs.
- `pg_lake_iceberg`: Iceberg metadata, snapshots, manifests, and catalog integration.
- `pg_lake_copy`: `COPY` interception for object-store files and lake formats.
- `pg_lake_engine`: shared query rewrite, type conversion, cleanup, and `pgduck_server` client layer.
- `pg_extension_base`: preload and lifecycle-worker infrastructure.
- `pg_map`: generated PostgreSQL map types used for nested lake data.

### Operational Caveats

- `pgduck_server` is required for lake queries and must have working object-store credentials and local socket connectivity from PostgreSQL.
- S3 and compatible credentials are resolved by the DuckDB secrets/credential chain. Grant only the bucket permissions required by the workload.
- Iceberg writes create Parquet files per statement. Batch inserts and run regular `VACUUM` to avoid many small files.
- The PostgreSQL extensions, `pgduck_server`, object-store data, and Iceberg catalog form one deployment unit. Back up and upgrade them as separate evidence layers; creating the extension alone does not prove the external services are usable.
