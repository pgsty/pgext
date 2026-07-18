## Usage

Sources:

- [jsonb_delete README at the reviewed commit](https://github.com/glynastill/jsonb_delete/blob/f87d296d725ae1ba207a3b65335967c4444c8d6d/README.md)
- [jsonb_delete 1.0 install SQL at the reviewed commit](https://github.com/glynastill/jsonb_delete/blob/f87d296d725ae1ba207a3b65335967c4444c8d6d/jsonb_delete--1.0.sql)
- [jsonb_delete regression SQL at the reviewed commit](https://github.com/glynastill/jsonb_delete/blob/f87d296d725ae1ba207a3b65335967c4444c8d6d/sql/jsonb_delete.sql)

`jsonb_delete` version 1.0 adds `jsonb_delete(jsonb, jsonb)` and the corresponding binary `-` operator. It removes only top-level entries from the left value whose key and complete value match an entry in the right value.

```sql
CREATE EXTENSION jsonb_delete;

SELECT '{"a": 1, "b": 2, "nested": {"x": 3}}'::jsonb
       - '{"a": 9, "b": 2, "nested": {"x": 4}}'::jsonb;

SELECT jsonb_delete(
  '{"a": 1, "b": 2}'::jsonb,
  '{"b": 2}'::jsonb
);
```

The first result retains `a` and `nested` because their values do not match, and removes `b` because both key and value match.

### Caveats

- Deletion is not recursive. A nested object is removed only when the entire nested value matches at the top level.
- The API predates most core JSONB modification operators and was written around PostgreSQL 9.4/9.5. Test compilation and behavior on the target major version.
- The operator compares JSONB values, rather than deleting by key name alone; do not substitute it for core key-deletion expressions without checking semantics.
