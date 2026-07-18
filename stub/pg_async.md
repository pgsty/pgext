## Usage

Sources:

- [Official extension control file](https://github.com/RekGRpth/pg_async/blob/20c2984087d89a6d9d30ea727325c808051147bb/pg_async.control)
- [Official upstream documentation](https://github.com/RekGRpth/pg_async/blob/20c2984087d89a6d9d30ea727325c808051147bb/README.md)

`pg_async` — C extension that provides async LISTEN/NOTIFY queue support on PostgreSQL replicas in hot standby.

The reviewed catalog snapshot records version `1.0`, kind `preload`, and implementation language `C`.
The curated compatibility set is `13,14`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_async";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_async';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
