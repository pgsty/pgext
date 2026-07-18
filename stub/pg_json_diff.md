## Usage

Sources:

- [Official upstream documentation](https://github.com/KhaledSMQ/pg-jsondiff/blob/95cd668636ce9c88d2f6ac57d92ac5301a81ed3c/README.md)

`pg_json_diff` — JSONB diff, JSON Patch, and merge-patch functions for PostgreSQL

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C++`.
The curated compatibility set is `12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_json_diff";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_json_diff';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
