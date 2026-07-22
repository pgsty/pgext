## Usage

Sources:

- [Official version 0.7.6 README](https://github.com/paradedb/paradedb/blob/v0.7.6/pg_lakehouse/README.md)
- [Official version 0.7.6 control file](https://github.com/paradedb/paradedb/blob/v0.7.6/pg_lakehouse/pg_lakehouse.control)
- [Official version 0.7.6 source tree](https://github.com/paradedb/paradedb/tree/v0.7.6/pg_lakehouse)

`pg_lakehouse` version `0.7.6` is ParadeDB's historical analytical FDW for querying object-store files and Delta Lake tables through Apache DataFusion. It supports S3-compatible, Azure, GCS, and local storage with Parquet, CSV, JSON, and Avro files. This API belongs to the tagged subproject and should not be confused with current ParadeDB search documentation.

### Core Workflow

The extension uses executor hooks, so preload it and restart PostgreSQL:

```conf
shared_preload_libraries = 'pg_lakehouse'
```

Create an object-store wrapper, server, and foreign table whose columns match the remote file exactly:

```sql
CREATE EXTENSION pg_lakehouse;
CREATE FOREIGN DATA WRAPPER s3_wrapper
  HANDLER s3_fdw_handler VALIDATOR s3_fdw_validator;

CREATE SERVER s3_server FOREIGN DATA WRAPPER s3_wrapper
OPTIONS (region 'us-east-1', allow_anonymous 'true');

CREATE FOREIGN TABLE trips (
  "VendorID" integer,
  "tpep_pickup_datetime" timestamp,
  "trip_distance" double precision,
  "total_amount" double precision
)
SERVER s3_server
OPTIONS (
  path 's3://paradedb-benchmarks/yellow_tripdata_2024-01.parquet',
  extension 'parquet'
);

SELECT count(*) FROM trips;
```

If `path` names a partitioned directory it must end in `/`. DataFusion is case-sensitive, so quote mixed-case column names.

### Helpers and Formats

- `arrow_schema(server => ..., path => ..., extension => ...)`: inspect Arrow fields before choosing PostgreSQL column types.
- `CALL connect_table('schema.table')`: establish the object-store connection early and validate table credentials.
- Object stores: Amazon S3/S3-compatible services, Azure Blob/ADLS Gen2, Google Cloud Storage, and local files.
- Formats in 0.7.6: Parquet, CSV, JSON, Avro, and Delta Lake; the README lists ORC and Iceberg as not yet implemented.

### Caveats

Version `0.7.6` documents PostgreSQL 14–16 only. Preloading affects the whole server, and query pushdown changes planning/execution paths; validate every workload against the pinned binary. Schema/type mismatches fail at query time, especially for `date`, `timestamp`, and `timestamptz`.

Object-store credentials, network egress, file consistency, partition layout, and cold connection cost all affect behavior. Grant access to foreign servers/tables narrowly, avoid embedding long-lived secrets in broadly visible server options, and inspect `EXPLAIN` to confirm which operations are actually pushed to DataFusion.
