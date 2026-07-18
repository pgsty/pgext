## Usage

Sources:

- [pg_map README at the reviewed commit](https://github.com/semenikhind/pg_map/blob/f1a9c628625e20df148afa92a3df57bb8523bd40/README.md)
- [pg_map install SQL at the reviewed commit](https://github.com/semenikhind/pg_map/blob/f1a9c628625e20df148afa92a3df57bb8523bd40/pg_map--1.0.sql)
- [pg_map C implementation at the reviewed commit](https://github.com/semenikhind/pg_map/blob/f1a9c628625e20df148afa92a3df57bb8523bd40/pg_map.c)

`pg_map` applies a one-argument PostgreSQL function to every element of an array and returns an array of the function results. Its two overloads identify the mapped function by OID or by text; a text signature containing parentheses is resolved as a specific `regprocedure`, while a bare name is resolved as `regproc`.

### Map a Function over an Array

```sql
CREATE EXTENSION pg_map;

SELECT pg_map(
  'upper(text)',
  ARRAY['alpha', 'beta', 'gamma']::text[]
);
```

The example applies `upper(text)` to each element and returns an uppercase text array. Supplying a full signature avoids ambiguity when a function name is overloaded.

### Caveats

- Upstream only claims compatibility with PostgreSQL 9.6devel. It does not document support for modern PostgreSQL releases.
- The version 1.0 C implementation uses PostgreSQL internal array, catalog, and function-manager APIs. Build and regression-test it against the exact server source before use.
- The mapped function must accept one argument. The implementation attempts an element-type cast when the array element type differs from the function argument type and errors if no cast function is available.
- Upstream publishes no license or release compatibility matrix.
