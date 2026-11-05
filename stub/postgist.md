## Usage

Sources:

- [Upstream source 1](https://github.com/postgis-t/postgis-t/blob/f5ecd265fd7ae266a1038d0d1946e334f5657db0/src/postgist/postgist--0.1.0.sql)
- [Upstream source 2](https://github.com/postgis-t/postgis-t/blob/f5ecd265fd7ae266a1038d0d1946e334f5657db0/src/postgist/postgist.control)

`postgist` 0.1.0 combines time-stamped measurements and geometry into PostGIS-based temporal composite types.

### Core Workflow

`timeseries` stores measurements, their range, period and spatial hull; `trajectory` stores time-stamped geometry. `tst_timeseries` and `tst_trajectory` aggregate rows. Accessors `tst_hull`, `tst_period`, `tst_range` and `tst_data` expose components.

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION postgist;
SELECT tst_timeseries(value, observed_at, geom)
FROM (VALUES
  (1::numeric, timestamp '2026-01-01', ST_Point(0,0)),
  (2::numeric, timestamp '2026-01-02', ST_Point(1,1))
) AS samples(value, observed_at, geom);
```

### Operational Boundaries

Requires PostGIS and installation privileges for extension objects. This is legacy 0.1.0 SQL/PLpgSQL source without a verified current PostgreSQL-major support range. Its bounding-box overlap helpers only identify possible intersections, not exact spatiotemporal intersection. Validate ordering, coordinate systems and interpolation behavior for the dataset before relying on results.
