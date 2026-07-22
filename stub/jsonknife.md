## Usage

Sources:

- [Official README and path model](https://github.com/niquola/jsonknife/blob/ce6e65ad39a0344feab49754636daf8bd06c5c92/README.md)
- [Extension SQL API](https://github.com/niquola/jsonknife/blob/ce6e65ad39a0344feab49754636daf8bd06c5c92/jsonknife--1.0.sql)
- [Official regression tests](https://github.com/niquola/jsonknife/blob/ce6e65ad39a0344feab49754636daf8bd06c5c92/sql/test.sql)

`jsonknife` 1.0 provides C functions for extracting and reducing values along declarative paths through `jsonb`. A path element that is a string selects an object key, a number selects an array index, and an object filters array elements by a matching pattern. Use it for compact, read-only traversal; it does not provide an index access method or JSON mutation API.

### Core Workflow

The second argument is a JSON array of paths. Each individual path is itself a JSON array:

```sql
CREATE EXTENSION jsonknife;

SELECT knife_extract_text(
  '{"friends":[
     {"best":true,"name":"Niquola"},
     {"best":false,"name":"Avraam"},
     {"best":true,"name":"Ivan"}
   ]}'::jsonb,
  '[["friends", {"best":true}, "name"]]'::jsonb
);
```

The result is a PostgreSQL array containing the matched names. Multiple path arrays may be supplied in the outer array, and traversal can flatten matches found in JSON arrays.

### Function Index

- `knife_extract(jsonb, jsonb) RETURNS jsonb[]` preserves JSON values.
- `knife_extract_text`, `knife_extract_numeric`, and `knife_extract_timestamptz` return typed PostgreSQL arrays.
- `knife_extract_min_numeric`, `knife_extract_max_numeric`, `knife_extract_min_timestamptz`, and `knife_extract_max_timestamptz` reduce matched values to one bound.
- `knife_date_bound(text, text) RETURNS timestamptz` interprets a partial date using either `min` or `max` as the missing-field boundary.

All functions are declared `STRICT` and `IMMUTABLE`; a NULL input produces NULL, and type conversion errors abort the statement. The implementation directly uses PostgreSQL's internal `jsonb` C structures and dates from early `jsonb` releases. Confirm binary/source compatibility with the target PostgreSQL major version before deployment, and benchmark large documents because extraction walks the supplied JSON rather than using a database index.
