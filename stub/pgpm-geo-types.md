## Usage

Sources:

- [packages/geotypes/README.md](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/geotypes/README.md)
- [packages/geotypes/sql/pgpm-geo-types--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/geotypes/sql/pgpm-geo-types--0.47.0.sql)
- [packages/geotypes/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/geotypes/Makefile)
- [packages/geotypes/pgpm-geo-types.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/geotypes/pgpm-geo-types.control)

`pgpm-geo-types` adds four domains over PostGIS types: `geo_point`, `geo_polygon`, `geography_point` and `geography_polygon`. Their underlying point/polygon typmods use SRID 4326.

### Core Workflow

```sql
CREATE EXTENSION "pgpm-geo-types";
CREATE TABLE geo_demo (id integer PRIMARY KEY, location geo_point);
INSERT INTO geo_demo VALUES (1, ST_SetSRID(ST_MakePoint(-122.4194, 37.7749), 4326));
```

### Operational Boundaries

Geometry and geography retain their respective PostGIS calculation semantics. The SQL defines domains and Graphile naming comments; it does not add a general topology-validity check or a coordinate transformation. Convert coordinates before assigning a different SRID. Domain names are unqualified in the installation script, so use the installation schema when qualifying them. Dependencies are `plpgsql`, `citext`, `postgis`, `pgpm-types`, `pgpm-verify`. This SQL-only module needs no library preload or restart; no PostgreSQL-major support matrix is published for this revision. Its fixed schema layout is not relocatable. Install the matching dependency closure before enabling it.
