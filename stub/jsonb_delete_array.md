## Usage

Sources:

- [Extension control file](https://github.com/mhagander/jsonb_delete_array/blob/5880579994603f4350431690594308c3d0c1d961/jsonb_delete_array.control)
- [Extension SQL](https://github.com/mhagander/jsonb_delete_array/blob/5880579994603f4350431690594308c3d0c1d961/jsonb_delete_array--1.0.sql)
- [C implementation](https://github.com/mhagander/jsonb_delete_array/blob/5880579994603f4350431690594308c3d0c1d961/jsonb_delete_array.c)

`jsonb_delete_array` adds the function `jsonb_delete_array(jsonb, text[])` and the operator `jsonb - text[]`. They remove several top-level object keys or top-level string array elements in one call:

```sql
CREATE EXTENSION jsonb_delete_array;

SELECT jsonb_delete_array(
  '{"a":1,"b":2,"c":3}'::jsonb,
  ARRAY['a', 'c']
);

SELECT '["a","b",1,true]'::jsonb - ARRAY['a', 'b']::text[];
```

### Scope

Version `1.0` operates only at the top level. For objects, matching keys are removed; for arrays, only matching string elements are removed, while numeric, boolean, object, and array elements remain. A scalar JSONB input raises an error, and an empty deletion list returns the input unchanged. The upstream repository provides source code but no README, so these narrow semantics come directly from the install SQL and C implementation.
