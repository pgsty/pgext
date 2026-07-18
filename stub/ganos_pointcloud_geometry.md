## Usage

Sources:

- [Alibaba Cloud point-cloud model guide](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/point-cloud-model)
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

This is provider-managed Ganos 7.0 software, not a portable public extension. Availability and permitted operations can vary by RDS engine version, edition, region, and policy. Validate `ST_MakePoint`, patch statistics, geometry conversion, spatial indexes, backup/restore, and export portability in a disposable instance. Plan for migration to environments where this component is absent.
