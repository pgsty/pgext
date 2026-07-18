## Usage

Sources:

- [Official PGXN distribution page](https://pgxn.org/dist/shared_ispell/)

`shared_ispell` — Shared-memory ispell dictionary template for reducing per-session initialization and memory.

The reviewed catalog snapshot records version `1.0.0`, kind `preload`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "shared_ispell";
SELECT extversion
FROM pg_extension
WHERE extname = 'shared_ispell';
```

The curated lifecycle is `abandoned`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
