## Usage

Sources:

- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_geo/pg_geo.control)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_geo/pg_geo--1.0.sql)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_geo/meson.build)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_geo/geo_ops.c)

`pg_geo` `1.0` provides two-dimensional geometry and bounding-box indexing. This preview module is scoped to the PostgreSQL 18 fork in AlohaDB 21.1.0; stock PostgreSQL compatibility is not established.

### Workflow

```sql
CREATE EXTENSION pg_geo;
SELECT ST_DistanceSphere(ST_Point(-157.8,21.3), ST_Point(-156.3,20.8));
```

Installation requires a superuser; function use also requires the appropriate object privileges.

### Operational boundaries

geo_point, geo_bbox, geo_polygon and geo_circle support the supplied ST_ predicates and GiST bounding-box operations. This is a small two-dimensional implementation, not a PostGIS compatibility promise. Inspect accepted geometry syntax, coordinate assumptions and exact predicates before relying on query results; public function names may overlap other spatial extensions.

This module corresponds to the former upstream name `alohadb_geo`. Follow AlohaDB’s cluster/object conversion procedure and do not substitute an unrelated implementation with a similar name. Rehearse dependencies, privileges and backup restoration before upgrades.
