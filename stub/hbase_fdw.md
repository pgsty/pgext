## Usage

Sources:

- [Official extension control file](https://github.com/troels/hbase_fdw/blob/a4c0c5d981113a0f37186cb85374601f9d69eb60/hbase_fdw.control)

`hbase_fdw` — Foreign data wrapper for querying HBase from PostgreSQL.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "hbase_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'hbase_fdw';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
