## Usage

Sources:

- [Official upstream README](https://github.com/pjungwir/bigintarray/blob/c0c6ba77e9b296009debe6c7acdbaded7c6a64a7/README.md)
- [Official extension control file (bigintarray.control)](https://github.com/pjungwir/bigintarray/blob/c0c6ba77e9b296009debe6c7acdbaded7c6a64a7/bigintarray.control)
- [Official extension SQL (bigintarray--1.0.sql)](https://github.com/pjungwir/bigintarray/blob/c0c6ba77e9b296009debe6c7acdbaded7c6a64a7/bigintarray--1.0.sql)

`bigintarray` — The bigintarray extension provides functions, operators, and index support for one-dimensional arrays of bigints (bigint[]), following the same behavior as PostgreSQL's built-in intarray extension but for bigint (8-byte integer) arrays instead of integer (4-byte integer) arrays. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION bigintarray;

-- Does the array contain both 1 and 2?
SELECT '{1,2,3}'::bigint[] @@ '1&2'::query_bigint;  -- true

-- Does the array contain 1 or 3?
SELECT '{1,2,3}'::bigint[] @@ '1|3'::query_bigint;  -- true

-- Does the array contain 1 and not 5?
SELECT '{1,2,3}'::bigint[] @@ '1&!5'::query_bigint;  -- true
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `bigintarray_del_elem(_int8, int8)` is an extension function and returns `_int8`.
- `bigintarray_push_array(_int8, _int8)` is an extension function and returns `_int8`.
- `bigintarray_push_elem(_int8, int8)` is an extension function and returns `_int8`.
- `bigintset(int8)` is an extension function and returns `_int8`.
- `bigintset_subtract(_int8, _int8)` is an extension function and returns `_int8`.
- `bigintset_union_elem(_int8, int8)` is an extension function and returns `_int8`.
- `boolop(_int8, query_bigint)` is an extension function.
- `bqarr_in(cstring)` is an extension function and returns `query_bigint`.
- `bqarr_out(query_bigint)` is an extension function and returns `cstring`.
- `g_bigint_compress(internal)` is an extension function and returns `internal`.
- `g_bigint_consistent(internal,_int8,smallint,oid,internal)` is an extension function.
- `g_bigint_decompress(internal)` is an extension function and returns `internal`.
- `g_bigint_options(internal)` is an extension function and returns `void`.
- `g_bigint_penalty(internal,internal,internal)` is an extension function and returns `internal`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- The control file marks the extension as trusted.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
