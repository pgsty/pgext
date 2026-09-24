## Usage

Sources:

- [README](https://github.com/mikkelam/compact-uuid/blob/f81ba87320fb42fcfa6b8f1c07ee905bb2e2f8d0/README.md)
- [Control file](https://github.com/mikkelam/compact-uuid/blob/f81ba87320fb42fcfa6b8f1c07ee905bb2e2f8d0/compact_uuid.control)
- [Cargo.toml](https://github.com/mikkelam/compact-uuid/blob/f81ba87320fb42fcfa6b8f1c07ee905bb2e2f8d0/Cargo.toml)
- [src/lib.rs](https://github.com/mikkelam/compact-uuid/blob/f81ba87320fb42fcfa6b8f1c07ee905bb2e2f8d0/src/lib.rs)

`compact_uuid` stores UUIDs in 16 bytes while rendering them as canonical 22-character Base64url text. The reviewed 0.1.1 extension supports PostgreSQL 18.

### Core Workflow

```sql
CREATE EXTENSION compact_uuid;
SELECT '550e8400-e29b-41d4-a716-446655440000'::compact_uuid;
SELECT 'VQ6EAOKbQdSnFkRmVUQAAA'::compact_uuid::uuid;
CREATE TABLE entities (id compact_uuid PRIMARY KEY, label text);
INSERT INTO entities VALUES
  ('550e8400-e29b-41d4-a716-446655440000', 'example');
SELECT * FROM entities
WHERE id = '550e8400-e29b-41d4-a716-446655440000'::uuid;
```

### Type Compatibility

The type accepts ordinary UUID text and canonical compact text, but always emits the compact representation. Its comparison and hash behavior interoperates with `uuid` through mixed-type operators and operator families. It supports indexes, primary keys, foreign keys, arrays, and binary COPY.

The shorter display does not reduce UUID storage below 16 bytes. Explicit casts are still required in contexts that need one common SQL type; bidirectional implicit casts are deliberately absent, so mixed-type `JOIN USING` can require rewriting as an explicit join condition.

### Operation

A superuser installs the relocatable extension. No preload is required. PostgreSQL 18 is the current compatibility boundary; validate client codecs and backup/restore behavior before changing an application’s UUID column type.
