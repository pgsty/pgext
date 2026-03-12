

## Usage

> [aggs_for_vecs: aggregate functions for arrays (vector/row-based)](https://github.com/pjungwir/aggs_for_vecs)

Provides aggregate functions that operate on arrays as vectors, computing position-wise statistics across multiple rows. Supports `SMALLINT`, `INTEGER`, `BIGINT`, `REAL`, and `DOUBLE PRECISION`.

```sql
CREATE EXTENSION aggs_for_vecs;
```

### Aggregate Functions

| Function | Description |
|---|---|
| `vec_to_count(anyarray)` | Count of non-nulls in each position |
| `vec_to_sum(anyarray)` | Sum in each position |
| `vec_to_min(anyarray)` | Minimum in each position |
| `vec_to_max(anyarray)` | Maximum in each position |
| `vec_to_mean(anyarray)` | Average in each position (returns `FLOAT[]`) |
| `vec_to_weighted_mean(values, weights)` | Weighted average in each position |
| `vec_to_var_samp(anyarray)` | Sample variance in each position |
| `vec_to_first(anyarray)` | First non-null in each position (use with ORDER BY) |
| `vec_to_last(anyarray)` | Last non-null in each position (use with ORDER BY) |
| `hist_2d(x, y, ...)` | 2-D histogram |
| `hist_md(vals, indexes, ...)` | N-dimensional histogram |

### Non-Aggregate Functions

| Function | Description |
|---|---|
| `vec_add(l, r)` | Element-wise addition |
| `vec_sub(l, r)` | Element-wise subtraction |
| `vec_mul(l, r)` | Element-wise multiplication |
| `vec_div(l, r)` | Element-wise division |
| `vec_elements(array, indexes)` | Select elements at given indexes |
| `pad_vec(array, length)` | Extend array to given length with NULLs |
| `vec_coalesce(array, default)` | Replace NULLs with default |
| `vec_trim_scale(numeric[])` | Trim trailing zeros from NUMERIC elements |
| `vec_without_outliers(vals, mins, maxs)` | Replace outliers with NULL |

### Examples

```sql
-- Position-wise minimum across rows
SELECT vec_to_min(vals) FROM (VALUES
  (ARRAY[1,2,3,4]),
  (ARRAY[5,0,-5,0]),
  (ARRAY[3,6,0,9])
) AS t(vals);
-- {1,0,-5,0}

-- Position-wise average
SELECT vec_to_mean(vals) FROM my_table;
```
