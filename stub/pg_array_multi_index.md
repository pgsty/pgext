## Usage

Sources:

- [Official README at the reviewed commit](https://github.com/rlichtenwalter/pg_array_multi_index/blob/d315364e110db6ec50e822e2eb5957d988d5a5a6/README.md)
- [Version 1.0 SQL declaration](https://github.com/rlichtenwalter/pg_array_multi_index/blob/d315364e110db6ec50e822e2eb5957d988d5a5a6/pg_array_multi_index--1.0.sql)
- [Array implementation](https://github.com/rlichtenwalter/pg_array_multi_index/blob/d315364e110db6ec50e822e2eb5957d988d5a5a6/pg_array_multi_index.c)
- [Extension control file](https://github.com/rlichtenwalter/pg_array_multi_index/blob/d315364e110db6ec50e822e2eb5957d988d5a5a6/pg_array_multi_index.control)

`pg_array_multi_index` provides `array_multi_index(anyarray, integer[])`, which selects many non-contiguous array elements in one C call and returns them as an array of the same element type. It is useful when an application already has a list of positions and wants to avoid repeatedly rescanning a variable-length array for separate subscripts.

### Core Workflow

```sql
CREATE EXTENSION pg_array_multi_index;

SELECT array_multi_index(
  ARRAY['zero', 'one', 'two', 'three', 'four'],
  ARRAY[2, 5, 2, NULL, 99]
);
-- {one,four,one,NULL,NULL}
```

Positions are interpreted as one-based offsets into the array's flattened storage order. The requested order is preserved, repeated positions repeat values, and a NULL or out-of-range position produces a NULL result element. An empty index array returns an empty array. Because the SQL function is strict, a NULL value array or NULL index array does not invoke the function.

### Shape and Compatibility Boundaries

The result is always a one-dimensional array with lower bound 1, even if the source has custom lower bounds or multiple dimensions. Source lower-bound metadata is ignored: position 1 means the first physically deconstructed element, not necessarily the source's declared subscript 1. Use built-in subscript logic when original dimensions or lower bounds are semantically important.

Version 1.0 uses the legacy `anyarray` polymorphic signature and PostgreSQL internal array APIs. The repository has no release tags, compatibility matrix, or recent maintenance record. Build and test it against the exact server major and exercise by-value, by-reference, NULL-containing, toasted, multidimensional, and custom-lower-bound arrays before depending on it.
