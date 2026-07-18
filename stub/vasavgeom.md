## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/accarniel/VagueGeometry/blob/96c4be4cd382b93398c56e0db2a0a01bc409ce5d/README.md)
- [Version 1.1 SQL API](https://github.com/accarniel/VagueGeometry/blob/96c4be4cd382b93398c56e0db2a0a01bc409ce5d/vasavgeom--1.1.sql)
- [Extension control file](https://github.com/accarniel/VagueGeometry/blob/96c4be4cd382b93398c56e0db2a0a01bc409ce5d/vasavgeom.control)
- [Project wiki](https://github.com/accarniel/VagueGeometry/wiki)

`vasavgeom` implements Vague Spatial Algebra on top of PostGIS. A `vaguegeometry` contains a definite kernel and a possible conjecture; topological predicates return `vaguebool`, while measures such as area or distance can return `vaguenumeric` intervals.

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION vasavgeom;
WITH v AS (
  SELECT VG_MakeVagueGeom(
    ST_GeomFromText('POLYGON((0 0,0 1,1 1,1 0,0 0))'),
    ST_GeomFromText('POLYGON((-1 -1,-1 2,2 2,2 -1,-1 -1))')
  ) AS geom
)
SELECT VG_AsText(geom) FROM v;
```

The API includes constructors and serializers, kernel/conjecture extraction, set operations, vague topological predicates, and minimum/maximum spatial measures. Kernel and conjecture must have compatible geometry type and SRID, and application code must preserve the algebra's uncertainty semantics rather than casting results prematurely to ordinary boolean or numeric values.

Version 1.1 is old native PostGIS-dependent code with no current compatibility matrix. Validate it against the exact PostgreSQL, PostGIS, GEOS, and compiler versions, especially before storing the custom binary type in long-lived tables or indexes.
