## Usage

Sources:

- [Alibaba Cloud point-cloud model guide](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/point-cloud-model)
- [Alibaba Cloud Ganos basic concepts](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/basic-concepts)
- [Alibaba Cloud RDS for PostgreSQL extension matrix](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/extensions-supported-by-apsaradb-rds-for-postgresql)

`ganos_pointcloud_geometry` is the Alibaba Cloud Ganos integration between point-cloud values and spatial geometry. Provider documentation pairs it with `ganos_pointcloud`, whose `pcpoint` and `pcpatch` types are described by rows in `pointcloud_formats`.

```sql
CREATE EXTENSION ganos_pointcloud WITH SCHEMA public CASCADE;
CREATE EXTENSION ganos_pointcloud_geometry WITH SCHEMA public CASCADE;
SELECT name, default_version, installed_version
FROM pg_available_extensions
WHERE name IN ('ganos_pointcloud', 'ganos_pointcloud_geometry');
```

Follow the provider guide to define a point-cloud schema before storing data. Schema IDs, SRIDs, dimension types, scale/offset values, and compression settings determine how binary points are interpreted; changing them without migration can corrupt meaning. Use `pcpatch` and compression for large data sets only after benchmarking precision and query behavior.

This is provider-managed Ganos software, not a portable public extension. The catalog records 7.0, while the provider matrix varies versions and availability by PostgreSQL engine; RDS minor-version policy can also restrict new creation while existing installations remain usable. GanosBase and PostGIS cannot be installed in the same schema according to the provider's basic-concepts documentation, so plan schemas before enablement. Validate `ST_MakePoint`, patch statistics, geometry conversion, spatial indexes, backup/restore, and export portability in a disposable instance. Plan for migration to environments where this component is absent.
