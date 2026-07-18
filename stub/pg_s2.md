## Usage

Sources:

- [Upstream README and API examples](https://github.com/yuiseki/pg_s2/blob/0cb0f78877223dcb4c2429a29d4d9edb3e106c83/README.md)
- [Extension control file](https://github.com/yuiseki/pg_s2/blob/0cb0f78877223dcb4c2429a29d4d9edb3e106c83/pg_s2.control)
- [Cargo package metadata](https://github.com/yuiseki/pg_s2/blob/0cb0f78877223dcb4c2429a29d4d9edb3e106c83/Cargo.toml)
- [Rust SQL implementation](https://github.com/yuiseki/pg_s2/blob/0cb0f78877223dcb4c2429a29d4d9edb3e106c83/src/lib.rs)
- [API specification](https://github.com/yuiseki/pg_s2/blob/0cb0f78877223dcb4c2429a29d4d9edb3e106c83/SPEC.md)

`pg_s2` exposes a compact S2 CellID API using PostgreSQL built-in spatial types, without requiring PostGIS. Its order-preserving `s2cellid` type supports B-tree indexes, token and bigint conversion, hierarchy and neighbour traversal, cap/rectangle coverings, boundaries, and great-circle distance.

### Create and inspect cells

PostgreSQL `point` values use longitude as X and latitude as Y:

```sql
CREATE EXTENSION pg_s2;

SELECT s2_lat_lng_to_cell(point(139.767, 35.681), 14);
SELECT s2_cell_to_token(
  s2_lat_lng_to_cell(point(139.767, 35.681), 14)
);
SELECT s2_cell_to_parent(
  s2_lat_lng_to_cell(point(139.767, 35.681), 14),
  12
);
SELECT *
FROM s2_cover_cap_ranges(point(139.767, 35.681), 2000.0, 12, 16);
```

For indexed spatial searches, store the cell, use `s2_cover_cap_ranges()` or `s2_cover_rect_ranges()` as a coarse B-tree prefilter, then apply an exact predicate such as `s2_great_circle_distance()`. A covering is deliberately an over-approximation and can include false positives. User-settable defaults include `pg_s2.default_level`, `pg_s2.default_cover_level`, and `pg_s2.earth_radius_m`.

Version `0.0.6` is described upstream as an early MVP/testing release. The README states PostgreSQL 14 through 17 support, while Cargo also exposes build features for 13 and 18; those extra features are not the same as a support guarantee. Invalid tokens, coordinates, or levels raise errors, and longitude/latitude order must be kept consistent throughout the application.
