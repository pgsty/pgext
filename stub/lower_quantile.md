

## Usage

> [lower_quantile: lower quantile aggregate for PostgreSQL](https://github.com/tvondra/lower_quantile)

Implements the "lower quantile" aggregate, which differs slightly from `percentile_disc` by returning the value whose rank in the sorted multiset is `floor(1 + q*(n-1))`.

```sql
CREATE EXTENSION lower_quantile;
```

### Functions

| Function | Description |
|---|---|
| `lower_quantile(value, quantile float)` | Compute the lower quantile for the given quantile value (0 to 1) |

### Examples

```sql
-- Compute the lower-quantile median
SELECT lower_quantile(i, 0.5)
FROM generate_series(1, 1000) s(i);

-- Compute the 95th lower quantile
SELECT lower_quantile(i, 0.95)
FROM generate_series(1, 1000) s(i);
```

This definition is used by some papers (e.g., the DDSketch paper) to formulate accuracy guarantees.
