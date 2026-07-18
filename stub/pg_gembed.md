## Usage

Sources:

- [Official PGXN distribution page](https://pgxn.org/dist/pg_gembed/)
- [Official primary documentation](https://pgxn.org/dist/pg_gembed/1.0.0/)

`pg_gembed` — Generate embeddings inside PostgreSQL

The reviewed catalog snapshot records version `1.0.0`, kind `standard`, and implementation language `C`.
Install and validate the declared extension dependencies first: `vector`.
The curated compatibility set is `18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_gembed";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_gembed';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
