

## Usage

> [pg_polyline: Google Encoded Polyline encoding & decoding for PostgreSQL](https://github.com/yihong0618/pg_polyline)

Fast Google Encoded Polyline encoding & decoding as a PostgreSQL extension. Built with `pgrx`.

```sql
CREATE EXTENSION pg_polyline;

-- Encode an array of points to a polyline string
SELECT polyline_encode(
  ARRAY[point(-120.2, 38.5), point(-120.95, 40.7), point(-126.453, 43.252)], 6
);
--          polyline_encode
-- ----------------------------------
--  _izlhA~rlgdF_{geC~ywl@_kwzCn`{nI

-- Decode a polyline string back to points
SELECT polyline_decode('_ibE_seK_seK_seK', 6);
--       polyline_decode
-- ---------------------------
--  {"(0.2,0.1)","(0.4,0.3)"}
```

The second parameter is the precision (number of decimal places).
