## Usage

Sources:

- [Alibaba Cloud RDS extension matrix](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/extensions-supported-by-apsaradb-rds-for-postgresql)
- [Alibaba Cloud RDS grid-model guide](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/grid-model)
- [Alibaba Cloud RDS ST_AsGrid reference](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/st-asgrid)

ganos_geomgrid is a provider-managed GanosBase component on Alibaba Cloud ApsaraDB RDS for PostgreSQL. It supplies GeoSOT geographic grids and H3 hexagonal grids for multi-resolution encoding, indexing, aggregation, and spatial relationship queries. The provider matrix lists version 7.0 for newer supported engines and older 6.x versions for other PostgreSQL majors.

### Basic use

Create it on an RDS instance and let CASCADE install the provider-managed dependencies:

```sql
CREATE EXTENSION ganos_geomgrid WITH SCHEMA public CASCADE;

SELECT ST_AsText(g)
FROM unnest(
    ST_AsGrid(
        ST_GeomFromText(
            'POINT(116.31522216796875 39.910277777777778)',
            4490
        ),
        15
    )
) AS g;
```

ST_AsGrid returns the GeoSOT cells intersecting a geometry at precision 1 through 32. Alibaba Cloud documents SRID 4490 as the native coordinate system; other supported inputs are transformed to it automatically. The extension also exposes the geomgrid and h3grid types plus conversion, hierarchy, spatial-relation, and index operations.

### Caveats

- This is an Alibaba Cloud service extension, not a separately published open-source package. Availability, exact version, privileges, and dependency behavior depend on the RDS PostgreSQL major version and current instance policy.
- Confirm support in the provider matrix for the exact engine version before deployment. If CREATE EXTENSION is rejected, upgrade the instance minor version or contact Alibaba Cloud support.
- Grid precision controls both accuracy and result size. High precision over large or complex geometries can produce many cells; benchmark storage, index size, and query cost.
- GeoSOT processing is centered on CGC2000/SRID 4490. Validate transformations and datum assumptions before mixing data from other spatial reference systems.
