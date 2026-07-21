## Usage

Sources:

- [Official Iceberg table guide](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/docs/iceberg-tables.md)
- [Version 3.4 control file](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_lake_iceberg/pg_lake_iceberg.control)
- [Iceberg metadata SQL API](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_lake_iceberg/pg_lake_iceberg--3.0.sql)
- [Version 3.4 catalog FDW SQL](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_lake_iceberg/pg_lake_iceberg--3.3--3.4.sql)

`pg_lake_iceberg` implements Iceberg metadata, snapshots, manifests, partition specifications, and catalog integration inside PostgreSQL. The familiar `CREATE TABLE ... USING iceberg` syntax is exposed by the dependent `pg_lake_table` component; users normally install both through `pg_lake`.

### Create and Inspect an Iceberg Table

```sql
CREATE EXTENSION pg_lake CASCADE;

SET pg_lake_iceberg.default_location_prefix =
    's3://analytics-bucket/warehouse';

CREATE TABLE events (
    event_time timestamptz NOT NULL,
    user_id bigint NOT NULL,
    payload jsonb
) USING iceberg
WITH (partition_by = 'day(event_time), bucket(32, user_id)');

SELECT table_namespace, table_name, metadata_location
FROM iceberg_tables
WHERE table_name = 'events';
```

Inspect an Iceberg metadata file and its referenced state:

```sql
SELECT lake_iceberg.metadata(metadata_location)
FROM iceberg_tables
WHERE table_name = 'events';

SELECT f.*
FROM iceberg_tables AS t
CROSS JOIN LATERAL lake_iceberg.files(t.metadata_location) AS f
WHERE t.table_name = 'events';
```

### Metadata and Catalog API

- `iceberg_tables`: `pg_catalog` view combining local managed tables and external catalog entries.
- `iceberg_namespace_properties`: catalog namespace properties.
- `lake_iceberg.metadata(uri)`: raw Iceberg metadata JSON.
- `lake_iceberg.files(uri)`: manifest path, content kind, data-file path/format, spec ID, record count, and file size.
- `lake_iceberg.snapshots(uri)`: sequence number, snapshot ID, timestamp, and manifest-list path.
- `lake_iceberg.data_file_stats(uri)`: per-file sequence and lower/upper bounds; execution is granted to `lake_read` rather than `PUBLIC`.
- `iceberg_catalog`: version 3.4 FDW for named PostgreSQL, object-store, or REST catalog configurations.

Define a user-managed REST catalog server and keep credentials in a user mapping:

```sql
CREATE SERVER my_polaris TYPE 'rest'
FOREIGN DATA WRAPPER iceberg_catalog
OPTIONS (rest_endpoint 'https://polaris.example.com');

CREATE USER MAPPING FOR app_role SERVER my_polaris
OPTIONS (client_id 'app', client_secret 'secret');

CREATE TABLE catalog_events (id bigint)
USING iceberg
WITH (catalog = 'my_polaris');
```

### Catalog and Storage Caveats

- User-created catalog servers require their own `USER MAPPING` credentials and do not fall back to the built-in REST catalog credential GUCs.
- The built-in `postgres`, `object_store`, and `rest` catalogs map to immutable extension-owned servers. Configure them through the documented GUCs rather than altering those servers.
- External modifications to `iceberg_tables` are blocked by default because changing metadata behind pg_lake can break transaction and query-engine consistency.
- Iceberg writes should be batched. Each statement can add Parquet files and snapshots; regular `VACUUM` compacts small files and expires data according to table/GUC policy.
- Iceberg has narrower representations for some PostgreSQL values. The default `out_of_range_values = 'error'` preserves integrity; `clamp` silently changes out-of-range temporal values and replaces some unsupported values with `NULL`.
