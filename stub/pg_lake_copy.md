## Usage

Sources:

- [Version 3.5 control file](https://github.com/Snowflake-Labs/pg_lake/blob/v3.5.3/pg_lake_copy/pg_lake_copy.control)
- [Official data-lake import and export guide](https://github.com/Snowflake-Labs/pg_lake/blob/v3.5.3/docs/data-lake-import-export.md)
- [Official file-format reference](https://github.com/Snowflake-Labs/pg_lake/blob/v3.5.3/docs/file-formats-reference.md)
- [3.4 to 3.5 SQL migration](https://github.com/Snowflake-Labs/pg_lake/blob/v3.5.3/pg_lake_copy/pg_lake_copy--3.4--3.5.sql)

`pg_lake_copy` extends PostgreSQL `COPY` so queries, heap tables, external lake tables, and Iceberg tables can exchange Parquet, CSV, or newline-delimited JSON files with local paths, HTTP endpoints, and configured object stores. It adds behavior through hooks and has no standalone SQL function API.

The pg_lake release and package version is `3.5.3`; the SQL extension version is `3.5`. Keep the libraries and query server from the same release.

### Enable the Component

The normal entry point installs `pg_lake_copy` and its exact dependencies together:

```sql
CREATE EXTENSION pg_lake CASCADE;
```

Its control file requires `pg_lake_engine`, `pg_lake_iceberg`, and `pg_lake_table`. The deployment also needs `pg_extension_base` in `shared_preload_libraries` and a running `pgduck_server`.

### Export and Import

Format is inferred from the path suffix or selected explicitly:

```sql
COPY (
    SELECT event_id, event_time, payload
    FROM events
    WHERE event_time >= DATE '2026-07-01'
)
TO 's3://analytics-bucket/events/july.parquet'
WITH (format 'parquet');

COPY events_archive
FROM 's3://analytics-bucket/events/july.parquet'
WITH (format 'parquet');
```

CSV and compressed output use standard-looking `COPY` options extended for the lake writer:

```sql
COPY (SELECT * FROM daily_summary)
TO 's3://analytics-bucket/summary/daily.csv.gz'
WITH (format 'csv', header true, compression 'gzip');
```

The destination can be a PostgreSQL heap table or an Iceberg table; the source can likewise be any query supported by the installed pg_lake stack.

### Format and Runtime Boundaries

- Parquet is columnar and preserves supported typed values; CSV and newline-delimited JSON have format-specific inference and conversion options documented upstream.
- Object-store access runs through `pgduck_server`. Its credential chain, network access, and bucket permissions must permit the requested read or write.
- `COPY` is one statement and participates in the surrounding PostgreSQL transaction, but remote files and cleanup also depend on the pg_lake transaction/queue machinery. Inspect failed operations and orphan cleanup before retrying large exports.
- Version `3.5` adds no standalone SQL API in `pg_lake_copy`; the `3.4` to `3.5` migration is empty. COPY behavior still depends on matching lake libraries and query-server binaries.
