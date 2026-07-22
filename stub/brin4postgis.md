## Usage

Sources:

- [Official README and experimental warning](https://github.com/2ndquadrant-it/brin4postgis/blob/fa3259493a7223bb7ffe0daee9d9cce7abc35d6c/README.md)
- [Version 0.0.1 control file](https://github.com/2ndquadrant-it/brin4postgis/blob/fa3259493a7223bb7ffe0daee9d9cce7abc35d6c/brin4postgis.control)
- [Operator and BRIN operator-class SQL](https://github.com/2ndquadrant-it/brin4postgis/blob/fa3259493a7223bb7ffe0daee9d9cce7abc35d6c/brin4postgis--0.0.1.sql)
- [Official regression workflow](https://github.com/2ndquadrant-it/brin4postgis/blob/fa3259493a7223bb7ffe0daee9d9cce7abc35d6c/sql/brin4postgis.sql)

`brin4postgis` 0.0.1 is an experimental pure-SQL operator class for BRIN indexing of PostGIS `box3d` values and 3D geometry bounding boxes. It can reduce the scan cost of spatially ordered, very large tables, but upstream explicitly says it is not production-ready.

### Create and Use the Index

Upstream documents PostgreSQL 9.5 or later and PostGIS 2.1.8 or later. The control file does not declare `postgis` as an extension dependency, so install it first:

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION brin4postgis;

CREATE INDEX observations_geom_brin
ON observations
USING brin (box3d(geom) brin_geometry_inclusion_ops_box3d);
```

Queries must use compatible `box3d` operators for the planner to consider the operator class:

```sql
SELECT id
FROM observations
WHERE box3d(geom) &&& ST_3DMakeBox(
  ST_MakePoint(-10, -10, -10),
  ST_MakePoint( 10,  10,  10)
);
```

Confirm index use and result equivalence against a sequential scan with `EXPLAIN` before relying on it.

### Main Objects

- `brin_geometry_inclusion_ops_box3d`: BRIN inclusion operator class for `box3d`.
- Spatial relationship operators including `&&&`, `@`, `~`, `<<`, `>>`, `<<|`, and `|>>`.
- Support functions such as `merge(box3d, box3d)`, `contains(box3d, box3d)`, and `isempty(box3d)`.

### Operational Boundaries

BRIN summaries are lossy bounding boxes, not exact geometries; qualifying heap rows still need rechecks. The extension stores summarized geometries as `box3d` with SRID zero, so original SRID metadata is not retained. Coordinate values remain in the original spatial reference system's units: do not combine mixed SRIDs in one indexed workflow without explicit transformation.

The SQL installs unqualified generic function and operator names in the chosen extension schema and depends on PostgreSQL's BRIN inclusion support functions. Its first and only upstream release dates from 2015. Pin the exact source, check for name conflicts and current PostGIS/core alternatives, and use it only after representative correctness and performance tests.
