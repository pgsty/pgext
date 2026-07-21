


## Usage

Sources:

- [tdigest v1.4.4 README](https://github.com/tvondra/tdigest/blob/v1.4.4/README.md)
- [v1.4.4 release](https://github.com/tvondra/tdigest/releases/tag/v1.4.4)
- [Extension control file](https://github.com/tvondra/tdigest/blob/v1.4.4/tdigest.control)

`tdigest` implements an approximate, mergeable t-digest for online rank statistics such as quantiles, percentile ranks, and trimmed means. It supports parallel aggregation and storing pre-aggregated digests for later rollups.

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

### Caveats

- Results are estimates. Validate the chosen compression against exact `percentile_cont` results on representative data before setting accuracy targets.
- Higher compression usually improves tail accuracy but increases state size and CPU cost.
- Stored digests can be merged across groups and time windows. Version `1.4.4` fixes combining digests created with different parameters, so use that patch level when heterogeneous states may meet.
- Version `1.4.4` also strengthens text-input parsing and validation and adds PostgreSQL 19 build/test coverage; malformed serialized digests that older builds accepted may now be rejected.
