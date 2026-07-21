## Usage

Sources:

- [Official pg_map README](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_map/README.md)
- [Version 3.4 control file](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_map/pg_map.control)
- [Base SQL definition](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_map/pg_map--1.2.sql)
- [Official extension tests and examples](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_map/tests/pytests/extension_test.py)

`pg_map` generates strongly typed key/value map domains from PostgreSQL types. A generated map is an array of composite key/value pairs, with type-specific extraction, cardinality, entry, and operator functions. It is used by pg_lake for nested data but can also be used directly.

### Create and Use a Map Type

```sql
CREATE EXTENSION pg_map;

-- Requires a privileged role; PUBLIC execution is revoked.
SELECT map_type.create('text', 'integer');
-- map_type.key_text_val_int
```

Construct a value and read it through generated functions or the `->` operator:

```sql
SELECT map_type.extract(
    '{"(me,1)","(myself,2)","(i,3)"}'::map_type.key_text_val_int,
    'i'
);
-- 3

SELECT
    '{"(me,1)","(myself,2)","(i,3)"}'::map_type.key_text_val_int
    -> 'myself';
-- 2

SELECT key, value
FROM map_type.entries(
    '{"(me,1)","(myself,2)","(i,3)"}'::map_type.key_text_val_int
);
```

### Generated API

- `map_type.create(keytype regtype, valtype regtype, typname text default null)`: idempotently creates or returns a map type for a key/value pair.
- `map_type.extract(map, key)` and generated `map -> key`: return the value for a key.
- `map_type.cardinality(map)`: returns the number of entries.
- `map_type.entries(map)`: expands a map to `key, value` rows.
- Generated names normally follow `map_type.key_<keytype>_val_<valuetype>`; supply `typname` when a controlled name is required.

### Type and Lifecycle Caveats

- Array types cannot be used as map keys. Array values and nested generated map types are supported.
- A call to `map_type.create` creates PostgreSQL types, functions, and operators. Treat it as schema DDL and run it in migrations rather than per-request code.
- Generated objects are registered as dependencies of `pg_map`; dropping the extension can remove them and columns that depend on them when `CASCADE` is used.
- Map values use PostgreSQL composite-array syntax. Duplicate-key and ordering semantics should be tested for the application's chosen construction path rather than assumed from JSON objects.
- Version `3.4` changes no map SQL API relative to `3.3`.
