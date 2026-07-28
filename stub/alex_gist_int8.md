## Usage

Sources:

- [Official upstream README](https://github.com/baofuhann/postgres/blob/f0cf77cfdf5c548de599237c47b7f63e72217024/contrib/README)
- [Official extension control file (alex_gist_int8.control)](https://github.com/baofuhann/postgres/blob/f0cf77cfdf5c548de599237c47b7f63e72217024/contrib/alex_gist_int8/alex_gist_int8.control)
- [Official extension SQL (alex_gist_int8--1.0.sql)](https://github.com/baofuhann/postgres/blob/f0cf77cfdf5c548de599237c47b7f63e72217024/contrib/alex_gist_int8/alex_gist_int8--1.0.sql)

`alex_gist_int8` — GiST index support for common PostgreSQL data types. Use it when an application needs this specific database capability. Upstream describes this capability as experimental.

### Core Workflow

```sql
CREATE EXTENSION alex_gist_int8;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `gbt_decompress(internal)` is an extension function and returns `internal`.
- `gbt_int8_compress(internal)` is an extension function and returns `internal`.
- `gbt_int8_consistent(internal,int8,int2,oid,internal)` is an extension function.
- `gbt_int8_distance(internal,int8,int2,oid,internal)` is an extension function and returns `float8`.
- `gbt_int8_fetch(internal)` is an extension function and returns `internal`.
- `gbt_int8_penalty(internal,internal,internal)` is an extension function and returns `internal`.
- `gbt_int8_picksplit(internal, internal)` is an extension function and returns `internal`.
- `gbt_int8_same(gbtreekey16, gbtreekey16, internal)` is an extension function and returns `internal`.
- `gbt_int8_train(oid)` is an extension function.
- `gbt_int8_union(internal, internal)` is an extension function and returns `gbtreekey16`.
- `gbtreekey16_in(cstring)` is an extension function and returns `gbtreekey16`.
- `gbtreekey16_out(gbtreekey16)` is an extension function and returns `cstring`.
- `gbtreekey8_in(cstring)` is an extension function and returns `gbtreekey8`.
- `gbtreekey8_out(gbtreekey8)` is an extension function and returns `cstring`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Upstream labels part or all of the project experimental.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
