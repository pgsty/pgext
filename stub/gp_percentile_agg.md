## Usage

Sources:

- [gpdb-doc/markdown/ref_guide/modules/gp_percentile_agg.html.md](https://github.com/open-gpdb/gpdb/blob/c6e63e9a53d01ed6ea3a93e5e280f38e8b9d808b/gpdb-doc/markdown/ref_guide/modules/gp_percentile_agg.html.md)
- [gpcontrib/gp_percentile_agg/gp_percentile_agg--1.0.0.sql](https://github.com/open-gpdb/gpdb/blob/c6e63e9a53d01ed6ea3a93e5e280f38e8b9d808b/gpcontrib/gp_percentile_agg/gp_percentile_agg--1.0.0.sql)
- [gpcontrib/gp_percentile_agg/Makefile](https://github.com/open-gpdb/gpdb/blob/c6e63e9a53d01ed6ea3a93e5e280f38e8b9d808b/gpcontrib/gp_percentile_agg/Makefile)
- [LICENSE](https://github.com/open-gpdb/gpdb/blob/c6e63e9a53d01ed6ea3a93e5e280f38e8b9d808b/LICENSE)
- [gpcontrib/gp_percentile_agg/gp_percentile_agg.control](https://github.com/open-gpdb/gpdb/blob/c6e63e9a53d01ed6ea3a93e5e280f38e8b9d808b/gpcontrib/gp_percentile_agg/gp_percentile_agg.control)

`gp_percentile_agg` supplies aggregates used by Greenplum GPORCA to optimize percentile and median queries. Version 1.0.0 is a Greenplum 6 module; applications normally keep using the built-in ordered-set aggregate syntax.

### Core Workflow

```sql
CREATE EXTENSION gp_percentile_agg;
SET optimizer_enable_orderedagg = on;
CREATE TABLE measurements (value double precision);
INSERT INTO measurements VALUES (10), (20), (30);
SELECT percentile_cont(0.5) WITHIN GROUP (ORDER BY value)
FROM measurements;
```

### Operational Boundaries

With GPORCA in use, enable `optimizer_enable_orderedagg` in the query session. The documented optimization requires a direct, noncomputed sort column, a constant scalar fraction rather than an array, and no `GROUP BY`. This is an eligibility boundary, not a guarantee that every percentile query runs faster.

The module defines `gp_percentile_cont` overloads for floating-point, interval, timestamp, and timestamp-with-time-zone values, plus polymorphic `gp_percentile_disc`. These helpers expect sorted input and the total-row and peer-group counts; use the ordinary `percentile_cont`, `percentile_disc`, or `median` interface for application queries.

Register the module as a superuser in each database. Its C library loads on demand without preload configuration. The extension is relocatable. The official guide requires dropping and recreating the module when upgrading an installation that used Greenplum 6.21.x to the corrected implementation in 6.22.0; review dependent objects before such maintenance. No upstream PostgreSQL major compatibility is claimed here.
