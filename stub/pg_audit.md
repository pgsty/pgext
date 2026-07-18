## Usage

Sources:

- [Official extension control file](https://github.com/disqus/pg_audit/blob/9c4c21c1a4bd94554fdcbc474043eaba68fb44cd/pg_audit.control)
- [Official upstream documentation](https://github.com/disqus/pg_audit/blob/9c4c21c1a4bd94554fdcbc474043eaba68fb44cd/README.md)

`pg_audit` — An extension for creating audit tables

The reviewed catalog snapshot records version `0.1.3`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plpgsql`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_audit";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_audit';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
