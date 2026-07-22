## Usage

Sources:

- [Official README and metric inventory](https://github.com/andrearosado/pg_landmetrics/blob/20ea7350111e3b8bfbd15e921b8f239503893801/README.md)
- [Metric type definitions](https://github.com/andrearosado/pg_landmetrics/blob/20ea7350111e3b8bfbd15e921b8f239503893801/sql/types/metric.sql)
- [`p_area` implementation and examples](https://github.com/andrearosado/pg_landmetrics/blob/20ea7350111e3b8bfbd15e921b8f239503893801/sql/functions/p_Area.sql)

`pg_landmetrics` 0.0.1 provides FRAGSTATS-style landscape metrics for PostGIS vector data, with overloads for `geometry` and `geography`. Use it to calculate metrics for individual patches, classes of patches, or an entire landscape. Most results are composite metric values, so read their fields rather than expecting a bare number.

### Core Workflow

Install PostGIS first, then create `pg_landmetrics`:

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION pg_landmetrics;
```

Patch functions accept one geometry at a time. For example, patch area returns a `metric` composite whose `value` is hectares when the default divisor is used:

```sql
SELECT id,
       (p_area(geom)).id    AS metric_id,
       (p_area(geom)).value AS hectares
FROM habitat_patches;
```

Class and landscape functions are aggregates. Group class metrics by the category represented by each row, and run landscape metrics over the full input set:

```sql
SELECT habitat_class, (c_totalarea(geom)).value AS class_area
FROM habitat_patches
GROUP BY habitat_class;

SELECT (l_numpatches(geom)).value AS patch_count
FROM habitat_patches;
```

### Metric Index

- Patch functions use the `p_` prefix: `p_area`, `p_corearea`, `p_coreareaindex`, `p_euclideanearestneighbourdistance`, `p_numcoreareas`, `p_perimeter`, `p_perimarearatio`, and `p_shapeindex`.
- Class aggregates use `c_`: total/core area, percentage of landscape, total edge, edge density, patch count, and patch density.
- Landscape aggregates use `l_`: total area/edge, edge and patch density, patch count/richness, richness density, and Shannon or Simpson diversity.
- `metric` stores a metadata identifier and numeric value; `metric_labeled` and `metric_labeled_pair` carry labeled collections used by selected metrics.

### Requirements and Caveats

PostGIS is required. For `geometry`, metric units depend on the geometry's coordinate reference system: calculations described in metres or hectares require an appropriate projected CRS, while longitude/latitude degrees do not yield those units. Several core-area and nearest-neighbour functions also need additional parameters or a meaningful collection of patches; consult the function-specific upstream pages before selecting thresholds. This is an early 0.0.1 project, so verify each chosen metric against a known dataset and inspect the composite result definition instead of assuming every function has the same signature.
