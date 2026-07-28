## Usage

Sources:

- [Official upstream README](https://github.com/massinissadjellouli/ctfstuff/blob/ddf59b4bb6dab358f8cda26a5b09adef8c2d62d3/README.md)
- [Official extension control file (lshvector.control)](https://github.com/massinissadjellouli/ctfstuff/blob/ddf59b4bb6dab358f8cda26a5b09adef8c2d62d3/tools/ghidra_11.2.1_PUBLIC/Ghidra/Features/BSim/src/lshvector/lshvector.control)
- [Official extension SQL (lshvector--1.0.sql)](https://github.com/massinissadjellouli/ctfstuff/blob/ddf59b4bb6dab358f8cda26a5b09adef8c2d62d3/tools/ghidra_11.2.1_PUBLIC/Ghidra/Features/BSim/src/lshvector/lshvector--1.0.sql)

`lshvector` — a feature vector type and a locality sensitive hashing index. Use it for the corresponding vector, model, or retrieval workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION lshvector;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `lsh_getweight(lshvector)` is an extension function and returns `float8`.
- `lsh_load()` is an extension function and returns `int4`.
- `lsh_reload()` is an extension function and returns `int4`.
- `lshvector_compare(lshvector,lshvector)` is an extension function and returns `lshvector_comptype`.
- `lshvector_gin_consistent(internal, int2, lshvector, int4, internal, internal, internal, internal)` is an extension function.
- `lshvector_gin_extract_query(lshvector,internal,int2,internal,internal,internal,internal)` is an extension function and returns `internal`.
- `lshvector_gin_extract_value(lshvector,internal)` is an extension function and returns `internal`.
- `lshvector_hash(lshvector)` is an extension function and returns `int8`.
- `lshvector_in(cstring)` is an extension function and returns `lshvector`.
- `lshvector_out(lshvector)` is an extension function and returns `cstring`.
- `lshvector_overlap(lshvector,lshvector)` is an extension function.
- `lshvector_recv(internal)` is an extension function and returns `lshvector`.
- `lshvector_send(lshvector)` is an extension function and returns `bytea`.
- `lshvector` is an extension-defined type.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
