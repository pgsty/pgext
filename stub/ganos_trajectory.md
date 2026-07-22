## Usage

Sources:

- [Alibaba Cloud trajectory introduction](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/trajectory-introduction)
- [Alibaba Cloud Ganos basic concepts](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/basic-concepts)
- [Alibaba Cloud extension creation guide](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/create-extensions)
- [Alibaba Cloud RDS for PostgreSQL extension matrix](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/extensions-supported-by-apsaradb-rds-for-postgresql)

`ganos_trajectory` is Alibaba Cloud's Ganos component for storing and analyzing moving-object trajectories. A trajectory combines spatial points with ordered `TIMESTAMP` observations and can carry point attributes and events; it is provider-managed rather than a portable community extension.

### Enablement and Workflow

Alibaba Cloud documents `ganos_spatialref` and `ganos_geometry` as prerequisites. `CASCADE` is the compact provider-supported installation path; inspect the installed dependency set on the exact RDS instance.

```sql
CREATE EXTENSION ganos_trajectory CASCADE;

SELECT name, default_version, installed_version
FROM pg_available_extensions
WHERE name IN ('ganos_spatialref', 'ganos_geometry', 'ganos_trajectory');
```

The data model supports trajectory points, attributes, events, and spatiotemporal bounding boxes such as `BoxNDF`. Time helpers include `ST_UnixEpochToTS` and `ST_TsToUnixEpoch`. Supply observations in timestamp order and verify duplicate-time, missing-point, timezone, interpolation, and units behavior with provider examples before loading production data.

### Provider Boundaries

The catalog records Ganos 7.0, but the provider matrix lists different versions and availability across PostgreSQL engine versions. Extension creation can also be restricted by RDS edition and minor engine version while an already-installed copy remains usable. Query the current matrix and `pg_available_extensions` instead of assuming the catalog version can be created.

GanosBase and PostGIS cannot be installed in the same schema according to the provider's basic-concepts documentation. Plan schemas before enablement, and test dependency creation, privileges, upgrades, backup/restore, spatial indexes, temporal predicates, export formats, and migration to a non-Ganos PostgreSQL service in a disposable instance.
