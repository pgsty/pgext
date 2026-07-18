## Usage

Sources:

- [Official extension control file](https://github.com/frost242/ssi_failures/blob/8b6b38b1f688632c3520d1b6aaac24feb7734c84/pg_stat_ssi_failures.control)
- [Official upstream documentation](https://github.com/frost242/ssi_failures/blob/8b6b38b1f688632c3520d1b6aaac24feb7734c84/README.md)

`pg_stat_ssi_failures` — Cluster-wide counter for PostgreSQL SSI serialization failures.

The reviewed catalog snapshot records version `0.1`, kind `preload`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_stat_ssi_failures";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_stat_ssi_failures';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
