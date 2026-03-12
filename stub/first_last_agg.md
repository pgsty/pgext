

## Usage

> [first_last_agg: first and last aggregate functions for PostgreSQL](https://github.com/wulczer/first_last_agg)

Provides `first` and `last` aggregate functions that return the first or last value in a group, operating on any element type.

```sql
CREATE EXTENSION first_last_agg;
```

### Functions

| Function | Description |
|---|---|
| `first(anyelement)` | Return the first value in the group |
| `last(anyelement)` | Return the last value in the group |

### Examples

```sql
-- Get the first and last values (use ORDER BY for deterministic results)
SELECT id, first(val ORDER BY ts), last(val ORDER BY ts)
FROM events
GROUP BY id;

-- Without ORDER BY, the ordering within groups is undefined
SELECT department, first(salary ORDER BY hire_date)
FROM employees
GROUP BY department;
```
