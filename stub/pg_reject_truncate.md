## Usage

Sources:

- [Official extension control file](https://github.com/katsuragi-y/pg_reject_truncate/blob/4ed44aa2bde9ff52e7569dce90c3cc0352062fe5/pg_reject_truncate.control)
- [Official upstream README](https://github.com/katsuragi-y/pg_reject_truncate/blob/4ed44aa2bde9ff52e7569dce90c3cc0352062fe5/README.md)

`pg_reject_truncate` — Rejects or prevents TRUNCATE operations in PostgreSQL.

The reviewed catalog snapshot records version `1.0`, kind `preload`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_reject_truncate";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_reject_truncate';
```

The curated lifecycle is `abandoned`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
