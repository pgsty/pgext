

## Usage

> [weighted_statistics: weighted statistical functions for PostgreSQL](https://github.com/schmidni/pg_weighted_statistics)

High-performance C extension providing weighted statistical functions with automatic sparse data handling (when `sum(weights) < 1.0`).

```sql
CREATE EXTENSION weighted_statistics;
```

### Functions

| Function | Description |
|---|---|
| `weighted_mean(values[], weights[])` | Weighted mean |
| `weighted_variance(values[], weights[], ddof)` | Weighted variance (ddof: 0=population, 1=sample) |
| `weighted_std(values[], weights[], ddof)` | Weighted standard deviation |
| `weighted_quantile(values[], weights[], quantiles[])` | Empirical CDF quantiles |
| `wquantile(values[], weights[], quantiles[])` | Type 7 (Hyndman-Fan) quantiles |
| `whdquantile(values[], weights[], quantiles[])` | Harrell-Davis quantiles |
| `weighted_median(values[], weights[])` | 50th percentile shortcut (empirical CDF) |

### Examples

```sql
-- Weighted mean
SELECT weighted_mean(ARRAY[1.0, 2.0, 3.0], ARRAY[0.2, 0.3, 0.5]);
-- 2.3

-- Weighted quantiles
SELECT weighted_quantile(ARRAY[10.0, 20.0, 30.0], ARRAY[0.3, 0.4, 0.3], ARRAY[0.25, 0.5, 0.75]);
-- {15.0, 20.0, 25.0}

-- Sparse data (auto-adds implicit zeros when sum(weights) < 1.0)
SELECT weighted_mean(ARRAY[10, 20], ARRAY[0.2, 0.3]);
-- 8.0  (equivalent to weighted_mean(ARRAY[10, 20, 0, 0], ARRAY[0.2, 0.3, 0.25, 0.25]))

-- Multiple statistics
SELECT weighted_mean(vals, weights),
       weighted_std(vals, weights, 1),
       weighted_quantile(vals, weights, ARRAY[0.05, 0.95])
FROM (SELECT array_agg(val) AS vals, array_agg(weight) AS weights FROM data) t;
```
