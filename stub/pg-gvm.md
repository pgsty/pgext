## Usage

Sources:

- [Official extension control file](https://github.com/greenbone/pg-gvm/blob/2427faa6e26ebb02d134b65efad8c9bf936dfa81/control.in)

`pg-gvm` — Functions for GVMd

The reviewed catalog snapshot records version `22.6`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg-gvm";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg-gvm';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
