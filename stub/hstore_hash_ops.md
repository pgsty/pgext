## Usage

Sources:

- [Official upstream README](https://github.com/obartunov/hstore_ops/blob/8839ab7447e913f9109e94cffd38b78aa97d5505/README.md)
- [Official extension control file (hstore_hash_ops.control)](https://github.com/obartunov/hstore_ops/blob/8839ab7447e913f9109e94cffd38b78aa97d5505/hstore_hash_ops.control)
- [Official extension SQL (hstore_hash_ops--1.0.sql)](https://github.com/obartunov/hstore_ops/blob/8839ab7447e913f9109e94cffd38b78aa97d5505/hstore_hash_ops--1.0.sql)

`hstore_hash_ops` — Revived non-default GIN operator classes for PostgreSQL hstore, ported to current PostgreSQL master (20devel):. Use it when an application needs this specific database capability. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION hstore_hash_ops;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `gin_compare_hstore_hash(int8, int8)` is an extension function and returns `int4`.
- `gin_compare_hstore_pair(bytea, bytea)` is an extension function and returns `int4`.
- `gin_compare_partial_hstore_hash(int8, int8, int2, internal)` is an extension function and returns `int4`.
- `gin_consistent_hstore_hash(internal, int2, internal, int4, internal, internal)` is an extension function.
- `gin_consistent_hstore_pair(internal, int2, internal, int4, internal, internal)` is an extension function.
- `gin_extract_hstore_hash(internal, internal)` is an extension function and returns `internal`.
- `gin_extract_hstore_pair(internal, internal)` is an extension function and returns `internal`.
- `gin_extract_hstore_query_hash(internal, internal, int2, internal, internal, internal, internal)` is an extension function and returns `internal`.
- `gin_extract_hstore_query_pair(internal, internal, int2, internal, internal, internal, internal)` is an extension function and returns `internal`.
- `gin_hstore_hash_ops` is an extension-defined operator class.
- `gin_hstore_pair_ops` is an extension-defined operator class.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- Install the confirmed extension dependencies first: `hstore`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
