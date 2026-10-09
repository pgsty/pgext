## Usage

Sources:

- [README.md](https://github.com/viniciuscva/qqespm_postgres_extension/blob/9c7d03a51a831e7066096c22e4888fe0736cb06f/README.md)
- [qqespm--1.0.sql](https://github.com/viniciuscva/qqespm_postgres_extension/blob/9c7d03a51a831e7066096c22e4888fe0736cb06f/qqespm--1.0.sql)
- [qqespm.control](https://github.com/viniciuscva/qqespm_postgres_extension/blob/9c7d03a51a831e7066096c22e4888fe0736cb06f/qqespm.control)

`qqespm` builds and executes spatial pattern queries over geo-textual objects. It is a PL/pgSQL extension using PostGIS; the control’s library pathname does not correspond to a shipped binary.

### Core Workflow

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION qqespm;
SELECT * FROM match_spatial_pattern(
  ARRAY[distance_constraint('school', 'pharmacy', 10, 10000, true, false)],
  'pois', ARRAY['school', 'pharmacy']);
```

### Operational Boundaries

Install PostGIS and prepare the expected POI table before querying: the default convention uses osm_id, geometry, centroid and keyword-bearing columns. `distance_constraint` and `connectivity_constraint` build JSONB constraints; `match_spatial_pattern` converts them to SQL and returns matching object identifiers. Generated query helpers expose the intermediate SQL. Distances in the documented examples are meters.

The README shows a PostgreSQL 16 installation path but supplies no supported-major matrix. No preload is needed; the control is relocatable and does not mark the extension trusted. Treat table names, keyword columns and query constraints as trusted inputs because the implementation constructs dynamic SQL. Check generated plans and appropriate PostGIS indexes on representative data. The project declares CC BY 4.0 in its README.
