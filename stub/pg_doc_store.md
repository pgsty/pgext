## Usage

Sources:

- [Official PGXN distribution page](https://pgxn.org/dist/pg_doc_store/)

`pg_doc_store` — Provides a JSONB document-store API with full-text search.

The reviewed catalog snapshot records version `0.2.1`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `pgcrypto`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_doc_store";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_doc_store';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
