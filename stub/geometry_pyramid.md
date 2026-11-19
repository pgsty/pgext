## Usage

Sources:

- [Version 0.1.3 guide](https://github.com/mapseekai/geometry_pyramid/blob/v0.1.3/README.md)
- [Release and packaged installation SQL](https://github.com/mapseekai/geometry_pyramid/releases/tag/v0.1.3)
- [Control file](https://github.com/mapseekai/geometry_pyramid/blob/v0.1.3/geometry_pyramid.control)
- [Catalog definitions](https://github.com/mapseekai/geometry_pyramid/blob/v0.1.3/sql/bootstrap.sql)
- [Lifecycle functions](https://github.com/mapseekai/geometry_pyramid/blob/v0.1.3/src/pyramid.rs)
- [Tile access](https://github.com/mapseekai/geometry_pyramid/blob/v0.1.3/src/tile.rs)
- [Storage checks](https://github.com/mapseekai/geometry_pyramid/blob/v0.1.3/src/storage.rs)
- [Worker execution context](https://github.com/mapseekai/geometry_pyramid/blob/v0.1.3/src/worker.rs)

`geometry_pyramid` builds materialized vector tiles over PostGIS geometry tables and returns MVT bytes through SQL. Version 0.1.3 supports PostgreSQL 15–18 and requires PostGIS 3.1 or later with MVT/protobuf support. It installs into `public` and is not relocatable. It uses a native library without requiring shared preload or an external daemon.

### Build and read tiles

```sql
CREATE EXTENSION IF NOT EXISTS postgis;
CREATE EXTENSION geometry_pyramid;

CREATE TABLE gp_points (
    id integer PRIMARY KEY,
    name text,
    geom geometry(Point, 4326)
);
INSERT INTO gp_points VALUES
    (1, 'Central', ST_SetSRID(ST_MakePoint(116.4074, 39.9042), 4326));
CREATE INDEX ON gp_points USING gist (geom);

SELECT ST_BuildPyramid(
    'gp_points', 'geom', 'id',
    '{"name":"gp_points_pyr","minZoom":0,"maxZoom":5,"attributes":["name"],"parallel":1}'
);
SELECT octet_length(ST_Tile('gp_points_pyr', 0, 0, 0));
SELECT ST_Pyramid('gp_points_pyr');
SELECT * FROM ST_Pyramids();
```

Table, column and pyramid names must be bare ASCII identifiers of at most 45 characters; schema-qualified names are rejected. Keep a stable feature ID and a GiST geometry index. Output tiles use Web Mercator; set `sourceSRS` when the source coordinate system differs from the default 4326.

`ST_Tile` returns `bytea`, or `NULL` for an empty, missing or out-of-bounds tile. With `splitSize` above zero, sparse tiles are generated on demand. `dynamicBeyondMaxZoom` controls on-demand slicing above the materialized maximum. `attributes` and `buildRules` control tile properties and per-zoom filtering. `parallel` selects build concurrency; workers need available PostgreSQL worker slots and committed source data. `tileSizeBudgetBytes` cannot be combined with parallel builds. Start with a bounded zoom range: the default maximum zoom of 18 can imply substantial work.

### Refresh and remove

```sql
SELECT ST_UpdatePyramid('gp_points_pyr', 'id = 1');
SELECT ST_UpdatePyramidByExtent(
    'gp_points_pyr', 116.35, 39.85, 116.45, 39.95, 4326
);
SELECT ST_DropPyramid('gp_points_pyr');
```

Refreshes are explicit, rather than automatic triggers. For moved or deleted features, refresh an extent covering both old and new locations; a predicate cannot rediscover a deleted feature. Rebuilding an existing pyramid preserves its published storage relation. Drop the pyramid before dropping its source table.

### Privileges and upgrades

Functions run with the caller's privileges. Builds require source reads, a writable output schema, and appropriate access to `geometry_pyramid_catalog`, `geometry_pyramid_stats` and the tile tables. The installation SQL does not grant these table privileges to application roles. Cached tiles are stored data: permission to read a tile table exposes its materialized contents independently of later source-row policy changes. Grant access accordingly and refresh stale data. Replacing tile storage is rejected when an active RLS policy filters its rows.

`buildRules[].filter` and the predicate passed to `ST_UpdatePyramid` are trusted SQL expressions evaluated as the caller; never pass untrusted input directly. Worker execution preserves the effective role, search path and row-security context.

Stop extension work, install the matching library and reconnect sessions before upgrading. On Windows, stop PostgreSQL before replacing a loaded DLL. PostgreSQL 13–14 must be upgraded before using this release.

```sql
ALTER EXTENSION geometry_pyramid UPDATE TO '0.1.3';
SELECT geometry_pyramid_version();
```

Upgrades from 0.1.1 or 0.1.2 need no tile rebuild. From 0.1.0, perform the update and explicit `source_rel`/`storage_rel` bindings in one transaction using the official migration example; verify no registered pyramid has an unbound relation before committing. These bindings use `regclass` to prevent later search-path substitution.
