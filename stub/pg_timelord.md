## Usage

Sources:

- [Official extension control file](https://github.com/rjuju/pg_timelord/blob/309155c3e6b5b5ee72117aa035236b791c47fe4d/pg_timelord.control)
- [Official upstream documentation](https://github.com/rjuju/pg_timelord/blob/309155c3e6b5b5ee72117aa035236b791c47fe4d/README.md)

`pg_timelord` — Proof-of-concept time-travel queries using commit timestamps and a custom MVCC snapshot.

The reviewed catalog snapshot records version `0.0.1`, kind `preload`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_timelord";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_timelord';
```

The curated lifecycle is `abandoned`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
