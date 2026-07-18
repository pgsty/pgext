## Usage

Sources:

- [Official upstream documentation](https://github.com/jnidzwetzki/pg_debug_scan#readme)

`pg_debug_scan` — Debug PostgreSQL table scans with custom MVCC snapshot definitions.

The reviewed catalog snapshot records version `0.1.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `11,12,13,14,15,16`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_debug_scan";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_debug_scan';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
