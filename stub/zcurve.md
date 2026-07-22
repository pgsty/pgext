## Usage

Sources:

- [Version 1.3 SQL API](https://github.com/bmuratshin/zcurve/blob/76b88d80794a4ccde95c01111453ad8d7fbdbf36/zcurve--1.3.sql)
- [Lookup implementation](https://github.com/bmuratshin/zcurve/blob/76b88d80794a4ccde95c01111453ad8d7fbdbf36/zcurve.c)
- [B-tree traversal implementation](https://github.com/bmuratshin/zcurve/blob/76b88d80794a4ccde95c01111453ad8d7fbdbf36/sp_tree.c)

`zcurve` interleaves two or three integer coordinates into a Z-order key and performs rectangular lookup by traversing an existing numeric B-tree index. Version 1.3 exposes `zcurve_val_from_xy`, `zcurve_num_from_xy`, `zcurve_num_from_xyz`, and 2-D/3-D lookup functions returning heap TIDs and decoded coordinates.

```sql
CREATE EXTENSION zcurve;
CREATE TABLE points (x integer, y integer, zkey numeric);
CREATE INDEX points_zkey_idx ON points (zkey);
SELECT * FROM zcurve_2d_lookup('points_zkey_idx', 0, 100, 0, 100);
```

The lookup argument is an index relation name, not a table. The implementation opens PostgreSQL B-tree internals directly, assumes its first key contains the expected numeric Z-order encoding, and returns physical TIDs; join results back to the heap immediately and recheck coordinate predicates because TIDs change after updates or table rewrites.

This is abandoned server-internal C code with copied numeric/B-tree structures, debug logging, no user README, and no stated modern PostgreSQL compatibility. A mismatched server major, index definition, coordinate range, or numeric representation can crash or return wrong rows. Use only for historical research on disposable data; prefer maintained spatial indexing such as GiST/SP-GiST/PostGIS for production.
