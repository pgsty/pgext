## Usage

Sources:

- [Official extension documentation](https://github.com/gojuno/lostgis/blob/85f26871adef0eb5f8fed43704a9d28b2fb5b80b/doc/lostgis.md)
- [Official extension SQL tree](https://github.com/gojuno/lostgis/tree/85f26871adef0eb5f8fed43704a9d28b2fb5b80b/sql)
- [Official PGXN release](https://pgxn.org/dist/lostgis/1.0.2/)

`lostgis` is a PostGIS utility collection for position/time/velocity tracks, OpenStreetMap-style opening hours, geometry repair and overlay, real-distance helpers for Web Mercator, grid cells, and map generalization. Version `1.0.2` is a 2017 SQL-only release; use individual functions after validating their assumptions rather than treating it as a complete GIS model.

### Core Workflow

PostGIS types must already be available even though the control file does not declare the dependency:

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION lostgis;

SELECT ST_GridCell(
  ST_SetSRID(ST_Point(4180000, 7500000), 3857),
  500
);

SELECT overlaps(
  timestamp '2017-08-13 13:00',
  'Mo-Fr 05:00-15:00; Sa 05:00-12:00'::text::opening_hours
);
```

The `overlaps()` input is a plain `timestamp`: convert the business location's time zone explicitly before testing opening hours. Its parser supports only simple weekday and hour/minute ranges.

### Object Index

- `tpv` stores a Web Mercator point, accuracy, heading, speed, timestamp, source, and OSM identifier.
- `opening_hours` stores parsed human text, validity flags, and a weekly minute mask.
- `project_position()` and `ST_AddTime()` operate on moving positions and tracks.
- `ST_Fast_Real_Buffer()`, `ST_Fast_Real_Length()`, and `ST_RealOffsetCurve()` approximate metre-based operations for projected geometries.
- `ST_Safe_Repair()`, `ST_Safe_Difference()`, and `ST_Safe_Intersection()` attempt recovery around invalid geometries.
- `ST_FilterSmallRings()`, `ST_LargestSubPolygon()`, `ST_LineAngleAtPoint()`, `ST_TimeLineMerge()`, `ST_AnglesEqual()`, `coslat()`, and `median()` provide additional helpers.

### Caveats

Most functions assume specific geometry types or SRID `3857`; confirm units, SRID, dimensionality, empty/NULL handling, and validity before use. The “safe” overlay functions may repair or snap data and therefore change topology instead of preserving exact input. The approximate real-distance helpers do not replace PostGIS geography for accuracy-sensitive work. Test representative polar/high-latitude data, antimeridian cases, invalid polygons, DST-aware opening hours, large arrays, and compatibility with the target PostGIS release.
