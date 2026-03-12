

## Usage

> [pg_logicalinspect: inspect logical decoding snapshot files](https://www.postgresql.org/docs/current/pglogicalinspect.html)

pg_logicalinspect provides SQL functions to examine serialized logical decoding snapshots stored in `pg_logical/snapshots/`, useful for debugging and educational purposes.

### Functions

**`pg_get_logical_snapshot_meta(filename text)`** -- snapshot file metadata:

```sql
SELECT * FROM pg_get_logical_snapshot_meta('0-40796E18.snap');

 magic      | 1369563137
 checksum   | 1028045905
 version    | 6
```

**`pg_get_logical_snapshot_info(filename text)`** -- detailed snapshot information:

```sql
SELECT * FROM pg_get_logical_snapshot_info('0-40796E18.snap');

 state                    | consistent
 xmin                     | 751
 xmax                     | 751
 start_decoding_at        | 0/40796AF8
 two_phase_at             | 0/40796AF8
 initial_xmin_horizon     | 0
 building_full_snapshot   | f
 in_slot_creation         | f
 last_serialized_snapshot | 0/0
 committed_count          | 0
 committed_xip            |
 catchange_count          | 2
 catchange_xip            | {751,752}
```

### Listing Available Snapshots

Combine with `pg_ls_logicalsnapdir()` to discover and inspect all snapshots:

```sql
SELECT ss.name, meta.*
FROM pg_ls_logicalsnapdir() AS ss,
     pg_get_logical_snapshot_meta(ss.name) AS meta;

SELECT ss.name, info.*
FROM pg_ls_logicalsnapdir() AS ss,
     pg_get_logical_snapshot_info(ss.name) AS info;
```

### Access Control

Restricted to superusers and members of `pg_read_server_files` by default.
