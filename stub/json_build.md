## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/pgexperts/json_build/blob/9a9ce7a38674499528e91dd46e280b95b94c7114/README.md)
- [Version 1.1.0 SQL objects](https://github.com/pgexperts/json_build/blob/9a9ce7a38674499528e91dd46e280b95b94c7114/json_build--1.1.0.sql)
- [Extension control file](https://github.com/pgexperts/json_build/blob/9a9ce7a38674499528e91dd46e280b95b94c7114/json_build.control)

`json_build` is a historical helper for constructing nested JSON. It provides `build_json_object(VARIADIC "any")`, `build_json_array(VARIADIC "any")`, and `json_object_agg("any", "any")`; all three return `json`.

```sql
CREATE EXTENSION json_build;
SELECT build_json_object(
  'user', build_json_object('id', 7, 'active', true),
  'roles', build_json_array('reader', 'writer')
);
```

Object keys must be non-null scalar values, and `build_json_object(VARIADIC "any")` requires an even argument count. Arrays become JSON arrays, records become objects, and existing JSON values pass through.

Modern PostgreSQL includes core JSON constructors and an aggregate with overlapping names and behavior. In particular, creating this extension can conflict with an existing `json_object_agg("any", "any")` signature. Prefer core `json_build_object`, `json_build_array`, and supported JSON aggregates for new code; install this old extension only after testing object-name conflicts and application semantics in an isolated database.
