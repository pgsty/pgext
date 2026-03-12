

## Usage

> [count_distinct: alternative to COUNT(DISTINCT ...) with better performance](https://github.com/tvondra/count_distinct)

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
FROM generate_series(1, 10000000) s(i);

-- Instead of:  SELECT id, COUNT(DISTINCT val) FROM test_table GROUP BY 1;
-- Use:
SELECT id, count_distinct(val) FROM test_table GROUP BY 1;

-- Aggregate distinct values into an array
SELECT id, array_agg_distinct(val) FROM test_table GROUP BY 1;

-- Count distinct elements across arrays
SELECT count_distinct_elements(ARRAY[1, 2, 2, 3]);
```
