## Usage

Sources:

- [Official upstream documentation](https://github.com/veridit/sql_saga/blob/9e39fcc7fc862fcb793f59d996eb01268cf9e822/README.md)

`sql_saga` — C and PL/pgSQL extension for valid-time temporal tables, temporal foreign keys, history-preserving views, and bulk temporal merge operations.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.
Install and validate the declared extension dependencies first: `btree_gist`, `plpgsql`.
The curated compatibility set is `18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "sql_saga";
SELECT extversion
FROM pg_extension
WHERE extname = 'sql_saga';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
