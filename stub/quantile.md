

## Usage

> [quantile: quantile aggregate functions for PostgreSQL](https://github.com/tvondra/quantile)

Provides aggregate functions to compute quantiles. Overloaded for `int`, `bigint`, `double precision`, and `numeric`.

```sql
CREATE EXTENSION quantile;
```

### Functions

| Function | Description |
|---|---|
| `quantile(value, quantile float)` | Compute a single quantile (0 to 1) |
| `quantile(value, quantiles float[])` | Compute multiple quantiles at once, returns array |

### Examples

```sql
-- Compute the median (0.5 quantile)
SELECT quantile(i, 0.5) FROM generate_series(1, 1000) s(i);
-- 500

-- Compute the 95th percentile
SELECT quantile(i, 0.95) FROM generate_series(1, 1000) s(i);

-- Compute all quartiles at once (more efficient than separate calls)
SELECT quantile(i, ARRAY[0.25, 0.5, 0.75])
FROM generate_series(1, 1000) s(i);
-- {250, 500, 750}
```

Note: Since PostgreSQL 9.4, built-in `percentile_cont` and `percentile_disc` functions are available. Consider using those first and only use this extension if it provides measurably better performance for your workload.
