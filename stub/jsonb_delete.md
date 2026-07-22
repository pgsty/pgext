## Usage

Sources:

- [Official README](https://github.com/glynastill/jsonb_delete/blob/f87d296d725ae1ba207a3b65335967c4444c8d6d/README.md)
- [Version 1.0 SQL API](https://github.com/glynastill/jsonb_delete/blob/f87d296d725ae1ba207a3b65335967c4444c8d6d/jsonb_delete--1.0.sql)
- [C deletion implementation](https://github.com/glynastill/jsonb_delete/blob/f87d296d725ae1ba207a3b65335967c4444c8d6d/jsonb_delete.c)
- [Official regression examples](https://github.com/glynastill/jsonb_delete/blob/f87d296d725ae1ba207a3b65335967c4444c8d6d/sql/jsonb_delete.sql)

`jsonb_delete` 1.0 adds `jsonb_delete(jsonb, jsonb)` and the binary `jsonb - jsonb` operator. For objects, it removes a top-level entry only when both its key and complete value match the right-hand object. For arrays, it removes top-level elements equal to elements in the right-hand array.

### Object Difference

```sql
CREATE EXTENSION jsonb_delete;

SELECT '{"a":1,"b":2,"nested":{"x":3}}'::jsonb
       - '{"a":9,"b":2,"nested":{"x":4}}'::jsonb;
-- {"a": 1, "nested": {"x": 3}}

SELECT jsonb_delete(
  '{"a":1,"b":2}'::jsonb,
  '{"b":2}'::jsonb
);
```

Changing only the right-hand value is not enough to delete a key. This differs from PostgreSQL's core key-deletion forms such as `jsonb - text`.

### Array Difference

```sql
SELECT '["a",2,{"x":1}]'::jsonb
       - '["a",{"x":1}]'::jsonb;
-- [2]
```

### Boundaries

Deletion is not recursive. A nested object or array is removed only when the entire top-level value is equal; matching descendants inside a retained value are not edited. The function is `STRICT`, so a SQL NULL argument produces SQL NULL, while JSON `null` participates in JSON equality normally.

This API was created for PostgreSQL 9.4/9.5 before most core JSONB modification operators existed and uses PostgreSQL's internal JSONB C structures. Build and regression-test it on each target major. Because it introduces the `jsonb - jsonb` overload, audit application operator resolution when other JSONB extensions define related overloads.
