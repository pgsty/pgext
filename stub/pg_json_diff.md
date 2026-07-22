## Usage

Sources:

- [Official README](https://github.com/KhaledSMQ/pg-jsondiff/blob/95cd668636ce9c88d2f6ac57d92ac5301a81ed3c/README.md)
- [Extension SQL](https://github.com/KhaledSMQ/pg-jsondiff/blob/95cd668636ce9c88d2f6ac57d92ac5301a81ed3c/pg_json_diff--1.0.sql)
- [C++ implementation](https://github.com/KhaledSMQ/pg-jsondiff/blob/95cd668636ce9c88d2f6ac57d92ac5301a81ed3c/src/json_diff_impl.cpp)

`pg_json_diff` 1.0 provides three immutable JSONB transformation functions backed by nlohmann/json: generate an RFC 6902 JSON Patch, apply such a patch, or apply an RFC 7386 merge patch. Use it for explicit document transformations, not as automatic conflict resolution or a stable audit-diff format.

### Core Workflow

Generate a patch, store or inspect it, and verify that applying it produces the expected target:

```sql
CREATE EXTENSION pg_json_diff;

WITH docs AS (
  SELECT '{"name":"Ada","age":30}'::jsonb AS old_doc,
         '{"name":"Ada","age":31,"city":"London"}'::jsonb AS new_doc
), patch AS (
  SELECT old_doc, new_doc, generate_json_diff(old_doc, new_doc) AS operations
  FROM docs
)
SELECT operations, apply_patch(old_doc, operations) = new_doc AS verified
FROM patch;
```

Validate externally supplied patches before applying them and run the transformation in a transaction with application-level authorization.

### Function Index

- `generate_json_diff(source, target)` returns an RFC 6902 operation array.
- `apply_patch(document, patch)` supports add, remove, replace, move, copy, and test operations.
- `merge_patch(document, patch)` recursively merges objects; a null member removes a key and arrays are replaced rather than merged.

All three functions accept and return `jsonb`.

### Operational Notes

Each call serializes PostgreSQL JSONB to text, parses it into a C++ JSON value, performs the operation, serializes again, and reparses the result as JSONB. Large or deeply nested documents therefore incur copying and parser cost. nlohmann/json and PostgreSQL JSONB have different numeric models, so test very large integers and high-precision decimals for fidelity. Patch errors are surfaced as general internal errors, and generated operation order or paths should not be treated as a canonical human audit record.
