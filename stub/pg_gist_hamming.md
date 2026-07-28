## Usage

Sources:

- [Official upstream README](https://github.com/fake-name/pg-spgist_hamming/blob/9fa70b08e0f0108de6a6673ce095c86a987d261d/README.md)
- [Official extension control file (pg_gist_hamming.control)](https://github.com/fake-name/pg-spgist_hamming/blob/9fa70b08e0f0108de6a6673ce095c86a987d261d/old/pg_gist_hamming.control)
- [Official extension SQL (pg_gist_hamming--1.0.sql)](https://github.com/fake-name/pg-spgist_hamming/blob/9fa70b08e0f0108de6a6673ce095c86a987d261d/old/pg_gist_hamming--1.0.sql)

`pg_gist_hamming` — support for indexing common datatypes in GiST. Use it for the corresponding vector, model, or retrieval workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_gist_hamming;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `gbt_decompress(internal)` is an extension function and returns `internal`.
- `gbt_int8_compress(internal)` is an extension function and returns `internal`.
- `gbt_int8_consistent(internal,int8,int2,oid,internal)` is an extension function.
- `gbt_int8_distance(internal,int8,int2,oid,internal)` is an extension function and returns `float8`.
- `gbt_int8_fetch(internal)` is an extension function and returns `internal`.
- `gbt_int8_hamming_distance(int8, int8)` is an extension function and returns `int8`.
- `gbt_int8_penalty(internal,internal,internal)` is an extension function and returns `internal`.
- `gbt_int8_picksplit(internal, internal)` is an extension function and returns `internal`.
- `gbt_int8_same(gbtreekey16, gbtreekey16, internal)` is an extension function and returns `internal`.
- `gbt_int8_union(internal, internal)` is an extension function and returns `gbtreekey16`.
- `gbtreekey16_in(cstring)` is an extension function and returns `gbtreekey16`.
- `gbtreekey16_out(gbtreekey16)` is an extension function and returns `cstring`.
- `gbtreekey32_in(cstring)` is an extension function and returns `gbtreekey32`.
- `gbtreekey32_out(gbtreekey32)` is an extension function and returns `cstring`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
