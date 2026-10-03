## Usage

Sources:

- [v3.0.2 README](https://github.com/tvondra/count_distinct/blob/v3.0.2/README.md)
- [v3.0.2 aggregate definitions](https://github.com/tvondra/count_distinct/blob/v3.0.2/sql/count_distinct--3.0.2.sql)
- [Control file](https://github.com/tvondra/count_distinct/blob/v3.0.2/count_distinct.control)

Provides an alternative to `COUNT(DISTINCT ...)` that avoids sorting and supports parallel aggregation.

```sql
CREATE EXTENSION count_distinct;
```

### Functions

| Function | Description |
|---|---|
| `count_distinct(value anyelement)` | Count distinct values (alternative to `COUNT(DISTINCT ...)`) |
| `array_agg_distinct(value anyelement)` | Aggregate distinct values into an array |
| `count_distinct_elements(value anyarray)` | Count distinct elements within input arrays |
| `array_agg_distinct_elements(value anyarray)` | Aggregate distinct elements from input arrays |

### Examples

```sql
CREATE TABLE test_table (id INT, val INT);
INSERT INTO test_table
SELECT mod(i, 1000), (1000 * random())::int
FROM generate_series(1, 10000) s(i);

-- Instead of:  SELECT id, COUNT(DISTINCT val) FROM test_table GROUP BY 1;
-- Use:
SELECT id, count_distinct(val) FROM test_table GROUP BY 1;

-- Aggregate distinct values into an array
SELECT id, array_agg_distinct(val) FROM test_table GROUP BY 1;

-- Count distinct elements across arrays
SELECT count_distinct_elements(ARRAY[1, 2, 2, 3]);
```

### Types, Memory and Version

Version 3.0.2 accepts fixed-length values passed by value, such as integers, and arrays of those values; it is not a general replacement for distinct text aggregation. Hashing a larger value first makes collisions possible and changes the result into an estimate. These aggregates keep their state in RAM and cannot reliably enforce `work_mem`; high cardinality or many concurrent groups can exhaust memory. Compare correctness, plans and memory use with the built-in aggregate on the actual workload rather than assuming a speedup. The control file is relocatable and has no preload requirement; C-extension installation normally needs a superuser.
