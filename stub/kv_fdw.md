## Usage

Sources:

- [Official extension control file](https://github.com/vidardb/pgrocks-fdw/blob/8f89fcad7e9c471e8dacfd167457716b7cf3d6e5/kv_fdw.control)

`kv_fdw` — Foreign data wrapper for key-value storage backed by RocksDB/VidarDB.

The reviewed catalog snapshot records version `0.0.1`, kind `preload`, and implementation language `C++`.
The curated compatibility set is `13`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "kv_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'kv_fdw';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
