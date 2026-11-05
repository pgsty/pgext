## Usage

Sources:

- [Official spatial documentation](https://github.com/Snowflake-Labs/pg_lake/blob/v3.6.0/docs/spatial.md)
- [Official control file](https://github.com/Snowflake-Labs/pg_lake/blob/v3.6.0/pg_lake_spatial/pg_lake_spatial.control)

- [Upstream 3.6.0](https://github.com/Snowflake-Labs/pg_lake/releases/tag/v3.6.0)
- [3.5 to 3.6 SQL](https://github.com/Snowflake-Labs/pg_lake/blob/v3.6.0/pg_lake_spatial/pg_lake_spatial--3.5--3.6.sql)

`pg_lake_spatial` connects `pg_lake` analytics tables to PostGIS geometry and DuckDB spatial/GDAL readers. It can infer schemas from spatial files, query remote GeoParquet or GeoJSON directly, and load supported files into regular analytics tables. It requires both `pg_lake` and `postgis`.

### Core Workflow

Create the dependency chain, then define a foreign table whose columns are inferred from a remote file:

```sql
CREATE EXTENSION pg_lake_spatial CASCADE;

CREATE FOREIGN TABLE world ()
SERVER pg_lake
OPTIONS (path 's3://my-bucket/world.parquet');

SELECT ST_GeometryType(geom), count(*)
FROM world
GROUP BY 1;
```

For a persisted copy, create a regular analytics table with `load_from = '...'` or use `COPY` as described by `pg_lake`. For GDAL-supported archives, set the documented `format 'gdal'` and, where needed, `zip_path` options.

### Formats and Execution

The extension covers GeoParquet, GeoJSON/GeoJSONSeq, WKB/WKT, and formats readable through GDAL. Spatial expressions that DuckDB can execute are pushed into the analytics engine; unsupported or semantically incompatible operations fall back to PostGIS in PostgreSQL.

### Caveats

GeoParquet support currently expects WKB geometry. GDAL assets may be downloaded and cached, so remote access needs working network/object-store credentials and predictable cache storage. Writing through GDAL is not supported. Analytics tables can report SRID `0`; clients such as QGIS may need an explicit SRID assignment.

Spatial pushdown is incomplete. A fallback to PostGIS is correct but can add transfer and execution overhead, and joins between unlike table types can be less efficient. Verify plans and results for every important function, keep both dependencies at compatible versions, and benchmark with representative file sizes and coordinate systems.

### Upstream 3.6.0

Upstream 3.6.0 adds PostgreSQL 19 support and uses DuckDB 1.5.6. Deploy matching libraries and `pgduck_server`, keep `pg_extension_base` preloaded, restart loaded processes and apply the supplied SQL updates to version 3.6. This component's 3.5-to-3.6 SQL migration is empty; the matching runtime and its dependencies still need the coordinated release upgrade.
