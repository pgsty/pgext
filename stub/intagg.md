

## Usage

> [intagg: integer aggregator and enumerator (obsolete)](https://www.postgresql.org/docs/current/intagg.html)

Provides an integer aggregator and enumerator. These are now wrappers around the built-in `array_agg()` and `unnest()` functions.

```sql
CREATE EXTENSION intagg;
```

### Functions

| Function | Description |
|---|---|
| `int_array_aggregate(integer)` | Aggregate integers into an array (wrapper for `array_agg()`) |
| `int_array_enum(integer[])` | Expand array into rows (wrapper for `unnest()`) |

### Examples

```sql
-- Aggregate integers into an array
SELECT id_left, int_array_aggregate(id_right) AS rights
FROM many_to_many
GROUP BY id_left;

-- Expand an integer array into rows
SELECT int_array_enum(ARRAY[1, 2, 3, 4]);
-- Returns: 1, 2, 3, 4 (as separate rows)
```

Note: This module is obsolete. Use the built-in `array_agg()` and `unnest()` functions instead for new code.
