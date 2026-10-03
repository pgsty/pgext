## Usage

Sources:

- [v2.0 README](https://github.com/xocolatl/extra_window_functions/blob/v2.0/README.md)
- [Version 2.0 SQL](https://github.com/xocolatl/extra_window_functions/blob/v2.0/extra_window_functions--2.0.sql)
- [1.0-to-2.0 upgrade SQL](https://github.com/xocolatl/extra_window_functions/blob/v2.0/extra_window_functions--1.0--2.0.sql)

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
SELECT lag_ignore_nulls(val, 1) OVER (ORDER BY ts) FROM measurements;
```

### Version 2.0 and Window Frames

Version 2.0 adds `count_ties()`, `avg_rank()`, `avg_percent_rank()`, `group_number(boolean)`, `run_length(anyelement)`, `run_position(anyelement)`, `ema(double precision, double precision)`, `interpolate(double precision)` and `most_common(anyelement)`. These functions cover peer groups, runs, smoothing and interpolation. Their frame and ordering rules differ: choose ORDER BY and the window frame deliberately, and inspect the official function reference rather than treating every function as an aggregate.

The lag example reads the previous non-null value; use `last_value_ignore_nulls` over a frame ending at the current row when the current non-null value must also be retained. Upgrade existing installations with `ALTER EXTENSION extra_window_functions UPDATE TO '2.0'`. The relocatable C extension does not require preload; installation normally needs a superuser. Upstream supports PostgreSQL 9.6–19. PostgreSQL 19 provides standard IGNORE NULLS for several functions, while FROM LAST and the extra 2.0 functions have their own uses.
