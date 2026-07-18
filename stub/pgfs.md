## Usage

Sources:

- [Official project or provider page](https://www.coredb.io)

`pgfs` — PostgreSQL extension for copying files on the server filesystem through SQL.

The reviewed catalog snapshot records version `0.0.1`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `11,12,13,14,15`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pgfs";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgfs';
```

The curated lifecycle is `abandoned`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
