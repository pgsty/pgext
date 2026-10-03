## Usage

Sources:

- [Official v0.0.7 README](https://github.com/aiyou178/pg_eviltransform/blob/v0.0.7/README.md)
- [v0.0.7 release notes](https://github.com/aiyou178/pg_eviltransform/releases/tag/v0.0.7)
- [v0.0.7 control file](https://github.com/aiyou178/pg_eviltransform/blob/v0.0.7/pg_eviltransform.control)
- [v0.0.7 upgrade SQL](https://github.com/aiyou178/pg_eviltransform/blob/v0.0.7/pg_eviltransform--0.0.6--0.0.7.sql)

`pg_eviltransform` extends PostGIS with coordinate transformations involving China's GCJ-02 and BD-09 systems. Version `0.0.4` also adds exact Jenks natural-break classification through `ST_JenksBins` array and aggregate overloads.

### Coordinate Transformation

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION pg_eviltransform;

-- WGS84 to GCJ-02 using a readable coordinate-system name.
SELECT ST_EvilTransform(
    ST_SetSRID('POINT(120 30)'::geometry, 4326),
    'GCJ02'
);

-- BD-09 to Web Mercator.
SELECT ST_EvilTransform(
    ST_SetSRID('POINT(120.011070620552 30.0038830555128)'::geometry, 990002),
    3857
);
```

Custom SRIDs are `990001` for GCJ-02 and `990002` for BD-09. When neither endpoint uses a custom system, `ST_EvilTransform` delegates to PostGIS `ST_Transform`; otherwise it converts through WGS84 (`4326`) when necessary.

Version 0.0.7 adds the geographic CGCS2000 aliases `CGCS2000`, `CGCS-2000`, `4490`, and `EPSG:4490` to text overloads. PostGIS/PROJ chooses the datum transformation. For a projected CGCS2000 Gauss-Kruger zone, use its explicit integer EPSG SRID.

```sql
SELECT ST_EvilTransform(
    ST_SetSRID('POINT(120 30)'::geometry, 4326), 'CGCS2000'
);
```

### Jenks Natural Breaks

```sql
-- Array form; NULL elements are ignored.
SELECT ST_JenksBins(ARRAY[1, 2, NULL, 10, 11]::numeric[], 2);

-- Streaming aggregate form for a large table.
SELECT ST_JenksBins(value, 7)
FROM measurements;

-- Return lower rather than upper bin edges.
SELECT ST_JenksBins(value, 7, true)
FROM measurements;
```

Array inputs support `numeric`, `double precision`, `real`, `bigint`, `integer`, and `smallint`. Aggregate inputs are `numeric` or `double precision`; cast other numeric columns when needed.

### API Index and Caveats

- `ST_EvilTransform(geometry, integer|text)` and `ST_EvilTransform(geometry, text, integer|text)`: four overloads corresponding to the PostGIS `ST_Transform` interface.
- `ST_JenksBins(values[], breaks [, invert])`: classifies an array and returns `double precision[]` edges.
- `ST_JenksBins(value, breaks [, invert])`: streaming aggregate that avoids materializing `array_agg`.
- PostGIS is a runtime prerequisite and must be installed before `pg_eviltransform`.
- Jenks inputs must be finite and `breaks` must be at least one. `numeric` values are converted to finite `f64`, so returned edges are floating-point values.
- When the distinct value count does not exceed `breaks`, the result is the sorted set of unique values; no valid input rows return `NULL`.

### 0.0.7 Compatibility

The 0.0.6-to-0.0.7 upgrade SQL updates the coordinate-name resolution functions. Upstream supports PostgreSQL 14-18 and PostgreSQL 19 beta4 with pgrx 0.19.3. The `ST_JenksBins` interface remains unchanged; package/build metadata is maintained separately from the upstream release.
