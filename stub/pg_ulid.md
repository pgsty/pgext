## Usage

Sources:

- [Official upstream documentation](https://github.com/edoceo/pg-ulid/blob/eda5ace98b0f1668297e95476cde5136ff06c97b/README.md)

`pg_ulid` — Generate ULID strings in PostgreSQL with a C function.

The reviewed catalog snapshot records version `0.0.4`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_ulid";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_ulid';
```

The curated lifecycle is `abandoned`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
