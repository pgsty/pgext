## Usage

Sources:

- [Official upstream documentation](https://github.com/ab1smo/pg_normalize_email/blob/709ccbe4cc2733adeb629bc4873d3f3afda868c4/README.md)
- [Official PGXN distribution page](https://pgxn.org/dist/pg_normalize_email/)

`pg_normalize_email` — Normalize email addresses in PostgreSQL

The reviewed catalog snapshot records version `1.0.9`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plpgsql`.
The curated compatibility set is `12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_normalize_email";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_normalize_email';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
