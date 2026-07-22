## Usage

Sources:

- [Extensions supported by ApsaraDB RDS for PostgreSQL](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/extensions-supported-by-apsaradb-rds-for-postgresql)
- [ST_SrEqual](https://www.alibabacloud.com/help/doc-detail/176141.html)
- [ST_SrReg](https://www.alibabacloud.com/help/en/apsaradb-for-rds/latest/st-srreg)
- [ST_SrFromEsriWkt](https://www.alibabacloud.com/help/doc-detail/176143.html)

`ganos_spatialref` is the Alibaba Cloud GanosBase spatial-reference extension. It compares coordinate-system definitions semantically, registers custom spatial reference systems, and converts Esri WKT into OGC WKT. It is a managed GanosBase component, not a standalone community package.

### Core Workflow

On an ApsaraDB RDS for PostgreSQL instance that offers the extension:

```sql
CREATE EXTENSION ganos_spatialref;

SELECT ST_SrEqual(
  '+proj=longlat +datum=WGS84 +no_defs',
  '+proj=longlat +ellps=WGS84 +datum=WGS84 +no_defs'
);

SELECT srid, auth_name, auth_srid
FROM spatial_ref_sys
WHERE ST_SrEqual(
  srtext::cstring,
  '+proj=longlat +ellps=GRS80 +no_defs'
)
LIMIT 1;
```

`ST_SrEqual(sr1, sr2, strict DEFAULT true)` accepts OGC WKT or PROJ.4 strings. It compares projection, ellipsoid, and axis parameters rather than raw text; strict mode also compares reference-ellipsoid names where applicable.

### Registration and Conversion

`ST_SrReg(sr)` registers WKT/PROJ.4 and returns an existing or newly assigned SRID. `ST_SrReg(auth_name, auth_id, sr)` adds authority metadata. Registration changes `spatial_ref_sys`, so validate definitions and restrict execution to trusted spatial administrators. `ST_SrFromEsriWkt(sr)` converts Esri WKT to OGC WKT, including normalization such as removing the Esri `D_` datum-name prefix.

### Service and Version Boundaries

Availability depends on the RDS edition, PostgreSQL major, and minor engine. The cited provider matrix lists version 7.0 for PostgreSQL 17, 6.9 for several older majors, and 6.3 for still older combinations; it does not offer the extension for every supported PostgreSQL release. Confirm the current instance with `pg_available_extension_versions` before use. Coordinate-system registration, built-in definitions, and function behavior belong to the managed GanosBase release, so test transformations after engine upgrades and do not assume catalog version 7.0 everywhere.
