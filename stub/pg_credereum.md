## Usage

Sources:

- [Official extension control file](https://github.com/postgrespro/pg_credereum/blob/master/pg_credereum.control)
- [Official upstream documentation](https://github.com/postgrespro/pg_credereum/blob/master/README.md)

`pg_credereum` — Cryptographically verifiable database auditing backed by blockchain-style transaction blocks.

The reviewed catalog snapshot records version `0.1`, kind `preload`, and implementation language `C`.
Install and validate the declared extension dependencies first: `plpgsql`.
The curated compatibility set is `10`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_credereum";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_credereum';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
