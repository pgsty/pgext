

## Usage

> [financial: financial calculation functions for PostgreSQL](https://github.com/intgr/pg_financial)

Provides the XIRR (Irregular Internal Rate of Return) aggregate function, similar to XIRR in spreadsheet programs.

```sql
CREATE EXTENSION financial;
```

### Functions

| Function | Description |
|---|---|
| `xirr(amount float8, date timestamptz [, guess float8] ORDER BY date)` | Compute Irregular Internal Rate of Return |

### Examples

```sql
-- Basic XIRR calculation
SELECT xirr(amount, time ORDER BY time) FROM transaction;
-- 0.0176201237088334

-- With explicit initial guess (Excel-compatible default is 0.1)
SELECT xirr(amount, time, 0.1 ORDER BY time) FROM transaction;

-- Per-portfolio XIRR
SELECT portfolio, xirr(amount, time ORDER BY time)
FROM transaction
GROUP BY portfolio;

-- As a window function
SELECT xirr(amount, time) OVER (ORDER BY time)
FROM transaction;
```

Returns NULL when Newton's method fails to converge. Providing a better guess may help in some cases.
