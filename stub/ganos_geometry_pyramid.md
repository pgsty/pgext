## Usage

Sources:

- [Alibaba Cloud Ganos vector pyramid documentation](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/vector-pyramid)
- [Extensions supported by ApsaraDB RDS for PostgreSQL](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/extensions-supported-by-apsaradb-rds-for-postgresql)

`ganos_geometry_pyramid` is an Alibaba Cloud GanosBase extension that turns a geometry table into a multilevel vector-tile pyramid and returns Mapbox Vector Tile (MVT/PBF) payloads. It is a managed-service capability, not a community extension that can be installed on an arbitrary PostgreSQL server.

### Core Workflow

Use a GanosBase-enabled ApsaraDB RDS for PostgreSQL instance, prepare a geometry table, and create a spatial index before building the pyramid:

```sql
CREATE EXTENSION ganos_geometry_pyramid WITH SCHEMA public CASCADE;

CREATE TABLE test (
    id bigint PRIMARY KEY,
    name text,
    geom geometry(Point, 4326)
);
CREATE INDEX test_geom_gist ON test USING gist (geom);

SELECT ST_BuildPyramid('test', 'geom', 'id', '');
SELECT ST_Tile('test', '0_0_0');
SELECT ST_DeletePyramid('test');
```

`ST_BuildPyramid(table_name, geom_col, id_col, options)` creates the pyramid. `ST_Tile(table_name, tile_id)` retrieves a tile whose identifier is `z_x_y` in Web Mercator (`EPSG:3857`). `ST_DeletePyramid(table_name)` removes generated pyramid data.

### Build Options

Pass JSON text in `options` when defaults are insufficient. Important keys include `name`, `parallel`, `tileSize`, `tileExtend`, `maxLevel`, and `buildRules`. `tileSize` must be a multiple of 256 in the documented 0–4096 range; `tileExtend` is 0–256; `buildRules` can assign per-level filters, selected attribute fields, and geometry merging. Levels above `maxLevel` are served in real time rather than prebuilt.

Parallel construction uses prepared transactions. Alibaba Cloud documents setting `max_prepared_transactions` to at least 100 and restarting the managed database before enabling it; the service limits parallelism to at most four times the CPU core count. Coordinate this with the service configuration rather than changing SQL alone.

### Service and Version Boundaries

Availability varies by ApsaraDB edition, PostgreSQL major version, and minor engine version. The provider matrix lists `ganos_geometry_pyramid` 7.0 for PostgreSQL 17 and 6.9 for several older majors at the time cited, so do not assume the catalog version is available everywhere. Confirm with the instance's extension-management page or `pg_available_extensions`. Building and refreshing a pyramid consume storage and compute, and source-table changes may require maintenance of the generated representation.
