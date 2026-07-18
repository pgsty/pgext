## Usage

Sources:

- [Official upstream documentation](https://github.com/norayr-m/dagdb-engine/tree/main/dagdb/pg_dagdb)

`pg_dagdb` — PostgreSQL SQL client extension for the DagDB graph engine daemon.

The reviewed catalog snapshot records version `0.0.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `13,14,15,16,17`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_dagdb";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_dagdb';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
