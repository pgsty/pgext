

## Usage

> [arraymath: element-by-element array operations for PostgreSQL](https://github.com/pramsey/pgsql-arraymath)

Provides element-by-element operators and utility functions for integer, float, and numeric arrays.

```sql
CREATE EXTENSION arraymath;
```

### Operators

All operators are prefixed with `@` to indicate element-by-element behavior. Works with array-vs-array (same length or cycling) and array-vs-scalar.

| Operator | Description | Returns |
|---|---|---|
| `@=` | Element-by-element equality | `boolean[]` |
| `@<` | Element-by-element less than | `boolean[]` |
| `@>` | Element-by-element greater than | `boolean[]` |
| `@<=` | Element-by-element less than or equals | `boolean[]` |
| `@>=` | Element-by-element greater than or equals | `boolean[]` |
| `@+` | Element-by-element addition | same type |
| `@-` | Element-by-element subtraction | same type |
| `@*` | Element-by-element multiplication | same type |
| `@/` | Element-by-element division | same type |

### Functions

| Function | Description |
|---|---|
| `array_sum(anyarray)` | Sum of all elements |
| `array_avg(anyarray)` | Average of all elements |
| `array_min(anyarray)` | Minimum element |
| `array_max(anyarray)` | Maximum element |
| `array_median(anyarray)` | Median element |
| `array_sort(anyarray)` | Sort ascending |
| `array_rsort(anyarray)` | Sort descending |

### Examples

```sql
-- Array vs scalar
SELECT ARRAY[1,2,3,4] @< 4;             -- {t,t,t,f}
SELECT ARRAY[3.4,5.6,7.6] @* 8.1;       -- {27.54,45.36,61.56}

-- Array vs array (cycling shorter array)
SELECT ARRAY[1,2,3,4,5,6] @* ARRAY[1,2]; -- {1,4,3,8,5,12}
SELECT ARRAY[1,2,3] @= ARRAY[3,2,1];     -- {f,t,f}

-- Utility functions
SELECT array_sort(ARRAY[9,1,8,2,7]);     -- {1,2,7,8,9}
SELECT array_sum(ARRAY[1,2,3,4,5]);      -- 15
SELECT array_median(ARRAY[1,2,3,4,5]);   -- 3
```
