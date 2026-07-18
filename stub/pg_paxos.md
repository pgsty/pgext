## Usage

Sources:

- [Official extension control file](https://github.com/microsoft/pg_paxos/blob/7d4f656b1cfa0943626b658fbc23a375eab41c3a/pg_paxos.control)
- [Official upstream documentation](https://github.com/microsoft/pg_paxos/blob/7d4f656b1cfa0943626b658fbc23a375eab41c3a/README.md)

`pg_paxos` — Paxos distributed consensus functions and table replication for PostgreSQL

The reviewed catalog snapshot records version `0.2`, kind `preload`, and implementation language `C`.
Install and validate the declared extension dependencies first: `dblink`, `plpgsql`.

```sql
CREATE EXTENSION "pg_paxos";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_paxos';
```

The curated lifecycle is `archived`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
