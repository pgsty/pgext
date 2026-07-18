## Usage

Sources:

- [Official extension control file](https://github.com/ratbox/pg-digest/blob/44ebe696700423cdd64ccc8f1d8b9dc18728934a/pg_digest.control)
- [Official upstream documentation](https://github.com/ratbox/pg-digest/blob/44ebe696700423cdd64ccc8f1d8b9dc18728934a/README.md)

`pg_digest` — Data type for efficient storage of hash digests in PostgreSQL.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_digest";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_digest';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
