

## Usage

> [extra_window_functions: additional window functions for PostgreSQL](https://github.com/xocolatl/extra_window_functions)

Provides window functions that simulate SQL Standard features not available in PostgreSQL syntax, plus novel functions like `flip_flop`.

```sql
CREATE EXTENSION extra_window_functions;
```

### Functions Simulating SQL Standard

| Function | Description |
|---|---|
| `lag_ignore_nulls(expr [, offset [, default]])` | LAG that skips NULL values |
| `lead_ignore_nulls(expr [, offset [, default]])` | LEAD that skips NULL values |
| `first_value_ignore_nulls(expr)` | FIRST_VALUE skipping NULLs |
| `last_value_ignore_nulls(expr)` | LAST_VALUE skipping NULLs |
| `nth_value_from_last(expr, offset)` | NTH_VALUE counting from end of frame |
| `nth_value_ignore_nulls(expr, offset)` | NTH_VALUE skipping NULLs |
| `nth_value_from_last_ignore_nulls(expr, offset)` | NTH_VALUE from last, skipping NULLs |

### Functions Extending SQL Standard (with default values)

| Function | Description |
|---|---|
| `first_value_ignore_nulls(expr, default)` | FIRST_VALUE with default when out of frame |
| `last_value_ignore_nulls(expr, default)` | LAST_VALUE with default when out of frame |
| `nth_value_from_last(expr, offset, default)` | NTH_VALUE from last with default |
| `nth_value_ignore_nulls(expr, offset, default)` | NTH_VALUE with default, skipping NULLs |
| `nth_value_from_last_ignore_nulls(expr, offset, default)` | Combined from-last, ignore-nulls, with default |

### Non-Standard Functions

| Function | Description |
|---|---|
| `flip_flop(expr [, expr])` | Flip-flop operator: returns false until first expr is true, then true until second expr matches |

### Examples

```sql
-- Equivalent to SQL Standard: NTH_VALUE(x, 3) FROM LAST IGNORE NULLS OVER w
SELECT nth_value_from_last_ignore_nulls(x, 3) OVER w FROM t WINDOW w AS (ORDER BY id);

-- Fill forward: carry last non-null value
SELECT lead_ignore_nulls(val, 1) OVER (ORDER BY ts) FROM measurements;
```
