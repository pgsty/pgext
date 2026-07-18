## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/gportenko/pg_rtiles/blob/b92acdbf435f8cec12df9d2ab6223a31662e417c/README.md)
- [Rust implementation](https://github.com/gportenko/pg_rtiles/tree/b92acdbf435f8cec12df9d2ab6223a31662e417c/pg_rtiles/src)
- [Cargo manifest](https://github.com/gportenko/pg_rtiles/blob/b92acdbf435f8cec12df9d2ab6223a31662e417c/pg_rtiles/Cargo.toml)

`rtiles` is a pgrx/PostGIS extension for generating Mapbox Vector Tiles for self-hosted interactive maps. Layer definitions live in `rtiles.layers` and specify a source table, geometry column, selected fields, optional clustering, and allowed database roles.

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION rtiles;
INSERT INTO rtiles.layers
  (name, source, geom_field, fields, clustered, roles)
VALUES
  ('roads', 'public.road', 'geom', ARRAY['class', 'name'], false,
   ARRAY['map_reader']);
```

The companion `tileserver` connects with `DB_HOST`, `DB_PORT`, `DB_DBNAME`, `DB_USER`, and `DB_PASSWORD`. Keep its credentials in a secret store, use a least-privilege login, enforce TLS and network controls, and verify that database role checks cannot be bypassed by the server's connection role.

Layer metadata drives dynamic table and column access. Restrict writes to `rtiles.layers`, quote and validate every source identifier, and expose only curated fields to avoid leaking private columns or expensive expressions. Benchmark tile size, spatial filtering, clustering, invalid geometries, large zoom ranges, cancellation, concurrent requests, and statement timeouts. The reviewed README records PostgreSQL 16 testing; verify each pgrx and PostGIS ABI combination before upgrading.
