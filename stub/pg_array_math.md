## Usage

Sources:

- [Official source](https://github.com/idexsts/pg_array_math/blob/3e9877014f567fb23b643db9a5c5db259419ead6/README.md)
- [Official source](https://github.com/idexsts/pg_array_math/blob/3e9877014f567fb23b643db9a5c5db259419ead6/pg_array_math.control)
- [Official source](https://github.com/idexsts/pg_array_math/blob/3e9877014f567fb23b643db9a5c5db259419ead6/pg_array_math--1.0.sql)
- [Official source](https://github.com/idexsts/pg_array_math/blob/3e9877014f567fb23b643db9a5c5db259419ead6/pg_array_math.c)

pg_array_math provides immutable, strict mathematical functions over one-dimensional arrays of smallint, integer, bigint, real, double precision or numeric. The reviewed source does not declare a license or a supported PostgreSQL-major matrix.

### Array operations

```sql
CREATE EXTENSION pg_array_math;
SELECT array_add(ARRAY[1,2,3], ARRAY[4,5,6]);
SELECT array_mul_scalar(ARRAY[1,2,3], 2);
SELECT array_sqrt(ARRAY[4.0,9.0]::numeric[]);
SELECT array_log(ARRAY[1.0,2.0]::double precision[]);
```

The API is array_add, array_sub, array_mul, array_div, array_add_scalar, array_mul_scalar, array_power, array_sqrt, array_abs, array_log and array_exp. Two-array operations require the same element type and length. The result preserves the first array's lower bound; NULL elements propagate, and a NULL input returns NULL.

### Limits

Multidimensional arrays, unsupported types, division by zero, invalid domains and numeric overflow raise errors. array_log and array_exp accept only floating-point or numeric arrays. Results retain the input type: integer division truncates, and integer power/square-root results are not floating-point approximations. Cast deliberately before computation. The relocatable C extension requires no preload; installation uses the ordinary untrusted-extension administrator boundary.
