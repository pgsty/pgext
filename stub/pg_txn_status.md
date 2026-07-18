## Usage

Sources:

- [Official extension control file](https://github.com/zilder/pg_txn_status/blob/c1d49147b72a8c3230349d99a819b72a66881cde/pg_txn_status.control)
- [Official upstream documentation](https://github.com/zilder/pg_txn_status/blob/c1d49147b72a8c3230349d99a819b72a66881cde/README.md)

`pg_txn_status` — One-byte txn_status data type for storing transaction state values.

The reviewed catalog snapshot records version `0.1`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_txn_status";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_txn_status';
```

The curated lifecycle is `abandoned`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
