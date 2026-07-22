## Usage

Sources:

- [Official getting-started guide](https://github.com/tumluliu/rostgis/blob/5558428d8cf260e1f5e5f4286dca97c6983c5a42/docs/user-guide/GETTING_STARTED.md)
- [Exported SQL functions](https://github.com/tumluliu/rostgis/blob/5558428d8cf260e1f5e5f4286dca97c6983c5a42/src/lib.rs)
- [Current geometry implementation](https://github.com/tumluliu/rostgis/blob/5558428d8cf260e1f5e5f4286dca97c6983c5a42/src/functions.rs)
- [Upstream GiST failure report](https://github.com/tumluliu/rostgis/blob/5558428d8cf260e1f5e5f4286dca97c6983c5a42/pgrx_gist_cbor_issue.md)

`rostgis` version `0.1.0` is an experimental Rust/pgrx spatial extension that imitates a small part of the PostGIS naming surface. Use it only in an isolated development database: it is not a compatible replacement for PostGIS and its SQL object names can conflict with PostGIS objects.

### Core Workflow

```sql
CREATE EXTENSION rostgis;

SELECT rostgis_version();
SELECT ST_AsText(ST_MakePoint(1.0, 2.0));
SELECT ST_AsText(ST_GeomFromText('LINESTRING(0 0, 3 4)'));
SELECT ST_Distance(ST_MakePoint(0, 0), ST_MakePoint(3, 4));
SELECT ST_SRID(ST_SetSRID(ST_MakePoint(1, 2), 4326));
```

The source parser handles basic `POINT`, `LINESTRING`, and single-ring `POLYGON` WKT. Common exported names include `ST_MakePoint`, `ST_GeomFromText`, `ST_AsText`, `ST_AsGeoJSON`, `ST_X`, `ST_Y`, `ST_SRID`, `ST_SetSRID`, `ST_Distance`, `ST_Area`, `ST_Length`, `ST_Perimeter`, and `ST_Equals`.

### Incomplete Semantics

Several exported names are placeholders or approximations:

- `ST_GeomFromWKB` always reports that WKB parsing is not implemented, while `ST_AsWKB` prefixes WKT text rather than producing WKB.
- `ST_MakePointZ` discards Z and `ST_Z` returns NULL.
- `ST_Distance` calculates only point-to-point distance and returns zero for other geometry pairs.
- `ST_Intersects` and `ST_Contains` use bounding-box tests rather than exact geometry predicates.

Do not use these functions for correctness-sensitive spatial analysis.

### Index and Upgrade Boundary

Although the repository contains GiST support code and examples, its own issue report records that index creation fails because of pgrx/CBOR serialization. Do not claim or depend on working GiST indexes for this release. The extension control file requires superuser installation and is not relocatable. Treat stored custom `geometry` values as tied to the exact experimental build, and export data before upgrading or removing it.
