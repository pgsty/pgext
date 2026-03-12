

## Usage

> [tdigest: t-digest percentile estimation for PostgreSQL](https://github.com/tvondra/tdigest)

Implements t-digest for on-line accumulation of rank-based statistics such as quantiles and trimmed means. Much faster than `percentile_cont`, supports parallelism, and allows pre-aggregation.

```sql
CREATE EXTENSION tdigest;
```

### Direct Aggregation Functions

| Function | Description |
|---|---|
| `tdigest_percentile(value, compression, quantile)` | Estimate a single percentile |
| `tdigest_percentile(value, compression, quantiles[])` | Estimate multiple percentiles |
| `tdigest_percentile_of(value, compression, value)` | Estimate percentile rank of a value |
| `tdigest_percentile_of(value, compression, values[])` | Estimate percentile ranks of multiple values |

### Pre-aggregation Functions

| Function | Description |
|---|---|
| `tdigest(value, compression)` | Build a t-digest from values |
| `tdigest_percentile(digest, quantile)` | Estimate percentile from a pre-built digest |
| `tdigest_percentile(digest, quantiles[])` | Estimate multiple percentiles from a pre-built digest |

### Incremental Update Functions

| Function | Description |
|---|---|
| `tdigest_add(digest, value)` | Add a single value to an existing digest |
| `tdigest_add(digest, values[])` | Add an array of values to an existing digest |
| `tdigest_union(digest, digest)` | Merge two digests |

### Utility Functions

| Function | Description |
|---|---|
| `tdigest_count(digest)` | Return the number of items in the digest |
| `tdigest_sum(digest, low, high)` | Trimmed sum within a value range |
| `tdigest_avg(digest, low, high)` | Trimmed average within a value range |

### Parameters

- `compression` -- controls accuracy (higher = more accurate, larger digest). Error is roughly `1/compression`.

### Examples

```sql
-- Instead of: SELECT percentile_cont(0.95) WITHIN GROUP (ORDER BY a) FROM t;
SELECT tdigest_percentile(a, 100, 0.95) FROM t;

-- Multiple percentiles
SELECT tdigest_percentile(a, 100, ARRAY[0.5, 0.95, 0.99]) FROM t;

-- Pre-aggregate for fast repeated queries
CREATE TABLE p AS SELECT a, b, tdigest(c, 100) AS d FROM t GROUP BY a, b;

-- Query pre-aggregated data (~1.5ms vs ~7s for exact)
SELECT a, tdigest_percentile(d, 0.95) FROM p GROUP BY a ORDER BY a;
```
