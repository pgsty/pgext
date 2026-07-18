## Usage

Sources:

- [Official extension control file](https://github.com/kotsachin/pg_shardman/blob/17820427be98e931aab1ef0f2813d526dcb4ac58/pg_shardman.control)
- [Official upstream documentation](https://github.com/kotsachin/pg_shardman/blob/17820427be98e931aab1ef0f2813d526dcb4ac58/readme.md)

`pg_shardman` — Manages PostgreSQL 10 hash-sharded tables across nodes using pg_pathman, postgres_fdw, and logical replication.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.
Install and validate the declared extension dependencies first: `pg_pathman`, `plpgsql`, `postgres_fdw`.
The curated compatibility set is `10`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_shardman";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_shardman';
```
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
