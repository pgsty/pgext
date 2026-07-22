## Usage

Sources:

- [GanosBase Geometry model for ApsaraDB RDS](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/geometry-model)
- [Extensions supported by ApsaraDB RDS for PostgreSQL](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/extensions-supported-by-apsaradb-rds-for-postgresql)
- [RDS extension-creation restrictions](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/limits-on-the-creation-of-the-pg-cron-extension)

`ganos_geometry` is Alibaba Cloud's provider-only GanosBase spatial engine for ApsaraDB RDS for PostgreSQL. It supplies OpenGIS-compatible `geometry` and `geography` data types, spatial functions, operators, and indexes for 2D, 3D, and 4D vector data; it is not a portable community package.

### Core Workflow

Create it in a deliberate schema, then store data with an explicit geometry subtype and SRID:

```sql
CREATE EXTENSION ganos_geometry WITH SCHEMA public CASCADE;

CREATE TABLE roads (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name text NOT NULL,
  geom geometry(LINESTRING, 3857) NOT NULL
);

INSERT INTO roads (name, geom)
VALUES (
  'sample',
  ST_GeomFromText('LINESTRING(0 0, 100 100)', 3857)
);

CREATE INDEX roads_geom_gix ON roads USING GIST (geom);
VACUUM ANALYZE roads (geom);

SELECT id, name, ST_AsText(geom), ST_Length(geom)
FROM roads;
```

`CASCADE` installs declared GanosBase dependencies when the engine supports it. The official model also documents GiST and BRIN indexes; use `gist_geometry_ops_nd` for n-dimensional GiST indexing and `brin_geometry_inclusion_ops_3d` or `brin_geometry_inclusion_ops_4d` for 3D or 4D BRIN indexing.

### Main SQL Surface

- Constructors and serializers include `ST_GeomFromText` and `ST_AsText`.
- Measurements include `ST_Length` and `ST_Area`.
- Spatial relationship predicates include `ST_Contains`, `ST_Covers`, and `ST_Intersects`.
- The official SQL reference covers the much larger PostGIS-compatible function surface; verify individual signatures against the installed service version.

### Provider, Version, and Schema Boundaries

Availability is tied to the RDS engine edition, PostgreSQL major, and minor engine version. In the current Standard Edition table, `ganos_geometry` is 7.0 on PostgreSQL 17, 6.9 on PostgreSQL 12–16, 6.3 on PostgreSQL 10–11, and unavailable on PostgreSQL 18. Other editions differ. Check `pg_available_extensions` on the target instance rather than treating the catalog's 7.0 value as universal.

Current RDS guidance says not to install GanosBase and PostGIS extensions into the same schema because both manage `spatial_ref_sys`; use separate schemas or remove the conflicting extension. Some older minor engine versions also block new `CREATE EXTENSION` operations for security remediation. Existing installations can remain while new creation is blocked, so upgrade the instance before retrying instead of bypassing the provider restriction.

Alibaba Cloud controls binaries, privileges, backups, upgrades, and support. Test migrations, SRIDs, spatial indexes, query plans, dump/restore, and cross-version behavior on the exact RDS build before making the custom types part of a long-lived schema.
