## Usage

Sources:

- [Official data-lake file query guide](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/docs/query-data-lake-files.md)
- [Official Iceberg table guide](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/docs/iceberg-tables.md)
- [Version 3.4 control file](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_lake_table/pg_lake_table.control)
- [FDW, server, utility, and access-method SQL](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_lake_table/pg_lake_table--3.0.sql)

`pg_lake_table` exposes object-store files as PostgreSQL foreign tables and provides the `USING iceberg` table syntax. It owns the `pg_lake` and `pg_lake_iceberg` foreign servers, file inspection/cache utilities, table catalogs, and transaction hooks; Iceberg metadata encoding is delegated to `pg_lake_iceberg`.

### Query External Files

Install the complete stack, including the required `pg_lake_engine`, `pg_lake_iceberg`, and `btree_gist` dependencies:

```sql
CREATE EXTENSION pg_lake CASCADE;
```

An empty column list asks pg_lake to infer the file schema:

```sql
CREATE FOREIGN TABLE external_events ()
SERVER pg_lake
OPTIONS (
    path 's3://analytics-bucket/events/*.parquet',
    filename 'true'
);

SELECT count(*) FROM external_events;
```

Create a writable Iceberg table through the extension-provided table access method:

```sql
CREATE TABLE managed_events (
    event_time timestamptz,
    payload jsonb
) USING iceberg;
```

### File and Table Utility Index

- `lake_file.list(pattern)`: lists matching object paths, sizes, modification times, and ETags.
- `lake_file.size(path)` and `lake_file.exists(path)`: inspect one remote object.
- `lake_file.preview(url, format, compression)`: returns inferred column names and types.
- `lake_file.delete(url)`: deletes a remote object; restrict it to roles that are allowed to remove data.
- `lake_file_cache.add(path, refresh)`, `remove(path)`, and `list()`: manage the local file cache for members of `lake_read`.
- `lake_iceberg.table_size(regclass)`: totals the current data-file sizes of an Iceberg table.
- Foreign servers `pg_lake` and `pg_lake_iceberg`: read-file and writable-Iceberg entry points.
- Access methods `iceberg` and `pg_lake_iceberg`: aliases used to translate `CREATE TABLE ... USING` into the extension's foreign-table representation.

```sql
SELECT path, file_size
FROM lake_file.list('s3://analytics-bucket/events/**/*.parquet');

SELECT *
FROM lake_file.preview('s3://analytics-bucket/events/sample.parquet');
```

### Operational Caveats

- `pg_extension_base` must be preloaded and `pgduck_server` must be running with credentials for every referenced location.
- `lake_read`, `lake_write`, and `lake_read_write` control access to servers, schemas, and utilities. Grant the narrowest role required by each application.
- External tables are references to files, not imported copies. File replacement, cross-region access, and cache invalidation can change latency or results independently of PostgreSQL catalog state.
- Iceberg inserts are optimized for batches rather than single rows. Use a staging heap table for high-rate row-at-a-time ingestion and periodically flush batches.
- Internal `lake_table.*` catalogs track files, field IDs, partitions, and recovery state. Do not modify them directly.
