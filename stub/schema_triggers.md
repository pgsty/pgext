## Usage

Sources:

- [Official extension control file](https://github.com/CartoDB/pg_schema_triggers/blob/07469164cd50460ff4f1d76b36c0a03f26d91e92/schema_triggers.control)
- [Official upstream documentation](https://github.com/CartoDB/pg_schema_triggers/blob/07469164cd50460ff4f1d76b36c0a03f26d91e92/README.md)

`schema_triggers` — Archived C hook extension adding relation, column, and trigger schema-change events to PostgreSQL event triggers.

The reviewed catalog snapshot records version `0.1`, kind `preload`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "schema_triggers";
SELECT extversion
FROM pg_extension
WHERE extname = 'schema_triggers';
```

The curated lifecycle is `archived`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
