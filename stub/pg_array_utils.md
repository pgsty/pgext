## Usage

Sources:

- [Official source](https://github.com/UuAcC/pgpro_practice_repos/blob/0ec613aeb5e6cfeb6bba202a5592940bf5f86cab/pg_array_utils/README.md)
- [Official source](https://github.com/UuAcC/pgpro_practice_repos/blob/0ec613aeb5e6cfeb6bba202a5592940bf5f86cab/pg_array_utils/pg_array_utils.control)
- [Official source](https://github.com/UuAcC/pgpro_practice_repos/blob/0ec613aeb5e6cfeb6bba202a5592940bf5f86cab/pg_array_utils/pg_array_utils--1.0.sql)
- [Official source](https://github.com/UuAcC/pgpro_practice_repos/blob/0ec613aeb5e6cfeb6bba202a5592940bf5f86cab/pg_array_utils/design_doc.md)
- [Official source](https://github.com/UuAcC/pgpro_practice_repos/blob/0ec613aeb5e6cfeb6bba202a5592940bf5f86cab/pg_array_utils/sql/init__.sql)

`pg_array_utils` 1.0 is a C extension for array transformations from a PostgreSQL training repository. The documented environment is PostgreSQL 18.4; broader major-version support and a distribution license are not established.

### Array transformations

The extension is relocatable and requires no preload. Its C functions use ordinary extension installation privileges.

```sql
CREATE EXTENSION pg_array_utils;
SELECT array_flatten(ARRAY[[1,2],[3,4]]);
SELECT array_chunk(ARRAY[1,2,3,4], 2);
SELECT array_compact(ARRAY[1,NULL,2]);
SELECT array_slice(ARRAY[1,2,3,4], 1, 4, 2);
```

### Function index and limits

`array_flatten` flattens dimensions; `array_chunk` partitions a one-dimensional array; `array_zip` returns JSONB tuples from variadic arrays; `array_compact` removes selected empty values; `array_pad` pads to a target size; `array_slice` selects a stepped range; `array_repeat` repeats a scalar or array. The padding and repeat overloads distinguish scalar and array arguments. Functions are declared `IMMUTABLE STRICT`, so SQL NULL arguments short-circuit to NULL.

`array_chunk` and `array_compact` reject multidimensional inputs. Follow the documented index, dimension and type rules rather than assuming Python slicing semantics. Large results remain limited by PostgreSQL allocation limits; this is not a streaming array processor.
