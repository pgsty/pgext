

## Usage

> [floatvec: element-by-element arithmetic on PostgreSQL arrays](https://github.com/pjungwir/floatvec)

Provides basic arithmetic functions for operating on arrays treated as vectors. Supports `SMALLINT`, `INTEGER`, `BIGINT`, `REAL`, and `DOUBLE PRECISION`.

```sql
CREATE EXTENSION floatvec;
```

### Functions

Each function accepts two arrays of the same length, or an array and a scalar, or a scalar and an array. Both arguments must be of the same type.

| Function | Description |
|---|---|
| `vec_add(a, b)` | Element-by-element addition |
| `vec_sub(a, b)` | Element-by-element subtraction |
| `vec_mul(a, b)` | Element-by-element multiplication |
| `vec_div(a, b)` | Element-by-element division |
| `vec_pow(a, b)` | Element-by-element exponentiation |

### Examples

```sql
-- Array + Array
SELECT vec_add(ARRAY[1,2,3], ARRAY[10,20,30]);  -- {11,22,33}

-- Array * Scalar
SELECT vec_mul(ARRAY[1.0, 2.0, 3.0], 2.0);     -- {2.0, 4.0, 6.0}

-- Scalar - Array
SELECT vec_sub(10, ARRAY[1,2,3]);               -- {9,8,7}
```
