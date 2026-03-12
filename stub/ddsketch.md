

## Usage

> [ddsketch: DDSketch quantile estimation for PostgreSQL](https://github.com/tvondra/ddsketch)

Implements DDSketch, a fully-mergeable quantile sketch with relative-error guarantees. Much faster than `percentile_cont` and supports parallelism.

```sql
CREATE EXTENSION ddsketch;
```

### Direct Aggregation Functions

| Function | Description |
|---|---|
| `ddsketch_percentile(value, alpha, nbuckets, quantile)` | Estimate a single percentile |
| `ddsketch_percentile(value, alpha, nbuckets, quantiles[])` | Estimate multiple percentiles |
| `ddsketch_percentile_of(value, alpha, nbuckets, value)` | Estimate percentile rank of a value |
| `ddsketch_percentile_of(value, alpha, nbuckets, values[])` | Estimate percentile ranks of multiple values |

### Pre-aggregation Functions

| Function | Description |
|---|---|
| `ddsketch(value, alpha, nbuckets)` | Build a ddsketch from values |
| `ddsketch_percentile(sketch, quantile)` | Estimate percentile from a pre-built sketch |
| `ddsketch_percentile(sketch, quantiles[])` | Estimate multiple percentiles from a pre-built sketch |

### Utility Functions

| Function | Description |
|---|---|
| `ddsketch_count(sketch)` | Return the number of items in the sketch |
| `ddsketch_sum(sketch, low, high)` | Trimmed sum within a value range |
| `ddsketch_avg(sketch, low, high)` | Trimmed average within a value range |

### Parameters

- `alpha` -- controls accuracy and sketch size (lower = more accurate, larger)
- `nbuckets` -- maximum number of buckets (each 8 bytes)

### Examples

```sql
-- Instead of: SELECT percentile_cont(0.95) WITHIN GROUP (ORDER BY a) FROM t;
SELECT ddsketch_percentile(a, 0.05, 1024, 0.95) FROM t;

-- Multiple percentiles at once
SELECT ddsketch_percentile(a, 0.05, 1024, ARRAY[0.5, 0.95, 0.99]) FROM t;

-- Pre-aggregate for fast repeated queries
CREATE TABLE p AS SELECT a, b, ddsketch(c, 0.05, 1024) AS d FROM t GROUP BY a, b;

-- Query pre-aggregated data (~1.5ms vs ~7s for exact)
SELECT a, ddsketch_percentile(d, 0.95) FROM p GROUP BY a ORDER BY a;
```
