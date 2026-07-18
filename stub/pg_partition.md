## Usage

Sources:

- [Official extension control file](https://github.com/disqus/pg_partition/blob/6ace26cfc6ffb2fdb6bb2163762a847b553d7691/pg_partition.control)

`pg_partition` — Generate and manage timestamp-based PostgreSQL table partitions

The reviewed catalog snapshot records version `0.2.0`, kind `puresql`, and implementation language `SQL`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_partition";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_partition';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
