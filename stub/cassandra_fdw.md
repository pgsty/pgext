## Usage

Sources:

- [Official upstream README](https://github.com/pgsql-io/cassandra_fdw/blob/2ee3d7950954f99464c6f3249d622136e558cfc8/README.md)

`cassandra_fdw` — foreign-data wrapper for querying Cassandra 3+

The reviewed catalog snapshot records version `3.1`, kind `standard`, and implementation language `C`.
The curated compatibility set is `11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "cassandra_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'cassandra_fdw';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
