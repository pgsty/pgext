## Usage

Sources:

- [Alibaba Cloud Ganos Raster model](https://help.aliyun.com/en/rds/apsaradb-rds-for-postgresql/raster-model)
- [Alibaba Cloud raster loading guide](https://help.aliyun.com/en/rds/apsaradb-rds-for-postgresql/load-raster-data)
- [ApsaraDB RDS for PostgreSQL extension matrix](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/extensions-supported-by-apsaradb-rds-for-postgresql)

`ganos_raster` is Alibaba Cloud's provider-managed GanosBase raster extension for ApsaraDB RDS for PostgreSQL. It stores and analyzes gridded imagery and similar raster data; it is not a portable community extension with a public SQL source tree.

### Core Workflow

```sql
CREATE EXTENSION Ganos_Raster WITH SCHEMA public CASCADE;

CREATE TABLE raster_table (
    id integer,
    raster_obj raster
);

INSERT INTO raster_table
VALUES (1, ST_ImportFrom('chunk_table', '/home/beijing.tif'));

SELECT ST_Height(raster_obj), ST_Width(raster_obj)
FROM raster_table
WHERE id = 1;
```

`ST_ImportFrom` also accepts an Alibaba OSS URL. The RDS instance must be authorized for the object, and the OSS bucket must be in the same region.

### Object Index

- `raster` is the stored raster type.
- `ST_ImportFrom(chunk_table, path)` imports a local or OSS-hosted raster file.
- `ST_Height` and `ST_Width` report raster dimensions.
- `ST_BuildPyramid`, `ST_BestPyramidLevel`, `ST_Clip`, and `ST_ClipDimension` support pyramid generation and spatial extraction.

Use Alibaba's Raster SQL reference for the full provider-version-specific surface.

### Operational Notes

Availability and version depend on the ApsaraDB PostgreSQL engine major and minor release. Catalog version `7.0` corresponds to a provider release, while other supported engine majors can expose different Ganos Raster versions; check `pg_available_extension_versions` and the current Alibaba matrix on the target instance.

`CASCADE` installs provider dependencies. Keep OSS credentials out of reusable SQL, logs, and query history because the documented URI form can contain an AccessKey ID and secret. Import paths are interpreted at the managed service boundary, not the client machine. Test storage growth, pyramid maintenance, backup behavior, and access controls with the chosen RDS release before production use.
