

## Usage

> [pgstattuple: tuple-level statistics](https://www.postgresql.org/docs/current/pgstattuple.html)

pgstattuple provides functions to obtain tuple-level statistics for tables and indexes, including dead tuple counts and free space metrics.

### Table Statistics

```sql
SELECT * FROM pgstattuple('my_table');

 table_len          | 81920
 tuple_count        | 1000
 tuple_len          | 72000
 tuple_percent      | 87.89
 dead_tuple_count   | 50
 dead_tuple_len     | 3600
 dead_tuple_percent | 4.39
 free_space         | 6320
 free_percent       | 7.71
```

**Approximate (faster) version** using visibility map and FSM:

```sql
SELECT * FROM pgstattuple_approx('my_table');
-- Returns: table_len, scanned_percent, approx_tuple_count, approx_tuple_len,
--          dead_tuple_count, dead_tuple_len, approx_free_space, ...
```

### B-Tree Index Statistics

```sql
SELECT * FROM pgstatindex('my_index');

 version            | 4
 tree_level         | 2
 index_size         | 1327104
 root_block_no      | 3
 internal_pages     | 1
 leaf_pages         | 160
 empty_pages        | 0
 deleted_pages      | 0
 avg_leaf_density   | 89.27
 leaf_fragmentation | 0
```

### GIN Index Statistics

```sql
SELECT * FROM pgstatginindex('my_gin_index');
-- Returns: version, pending_pages, pending_tuples
```

### Hash Index Statistics

```sql
SELECT * FROM pgstathashindex('my_hash_index');
-- Returns: version, bucket_pages, overflow_pages, bitmap_pages,
--          unused_pages, live_items, dead_tuples, free_percent
```

### Page Count

```sql
SELECT pg_relpages('my_table');
```

### Access

Restricted to superusers and roles with `pg_stat_scan_tables` privileges.
