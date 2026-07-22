## Usage

Sources:

- [Official README](https://github.com/scottjustin5000/pg-sum-array/blob/4845e1e605773ea51ca51f8044dbd25da2698a0b/README.md)
- [Version 0.0.1 SQL declaration](https://github.com/scottjustin5000/pg-sum-array/blob/4845e1e605773ea51ca51f8044dbd25da2698a0b/sum_array--0.0.1.sql)
- [C implementation](https://github.com/scottjustin5000/pg-sum-array/blob/4845e1e605773ea51ca51f8044dbd25da2698a0b/sum_array.c)

`sum_array` adds a C function that totals a PostgreSQL numeric array and always returns `double precision`. It accepts arrays whose element type is `smallint`, `integer`, `bigint`, `real`, or `double precision`; it is not an aggregate over rows.

### Core Workflow

```sql
CREATE EXTENSION sum_array;

SELECT sum_array(ARRAY[1, 2, 3, 4, 5]::integer[]);
SELECT sum_array(ARRAY[1.1, 2.2, 3.3]::double precision[]);
```

The SQL function is declared `IMMUTABLE` and `STRICT`. A NULL array therefore returns NULL, while a non-NULL supported array is converted element by element to `float8` and accumulated from zero.

### Numerical and NULL Caveats

The implementation does not inspect the per-element NULL flags returned by PostgreSQL. Do not pass arrays containing NULL elements or rely on their current accidental result; remove or replace them explicitly first. Unsupported element types, including `numeric`, raise an error.

All integer inputs are converted to `double precision`, so sufficiently large `bigint` values lose exactness, and floating-point totals retain ordinary rounding, overflow, infinity, and NaN behavior. The project is abandoned at `0.0.1`; PostgreSQL's built-in `unnest()` with `sum()` is better maintained and lets the caller choose explicit NULL and result-type semantics.
