## Usage

Sources:

- [Official upstream documentation](https://github.com/pgMemento/pgMemento/wiki/Home)
- [Official PGXN distribution page](https://pgxn.org/dist/pgmemento/)

`pgmemento` — Audit trail with schema versioning for PostgreSQL using transaction-based logging

The reviewed catalog snapshot records version `0.7.4`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plpgsql`.
The curated compatibility set is `13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pgmemento";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgmemento';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
