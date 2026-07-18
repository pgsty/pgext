## Usage

Sources:

- [Official extension control file](https://github.com/usagi-coffee/pg_chainsync/blob/c59b7a879dd427bb04f855c29a10979561f4fa08/pg_chainsync.control)
- [Official upstream documentation](https://github.com/usagi-coffee/pg_chainsync/blob/c59b7a879dd427bb04f855c29a10979561f4fa08/README.md)
- [Official Rust package manifest](https://github.com/usagi-coffee/pg_chainsync/blob/c59b7a879dd427bb04f855c29a10979561f4fa08/Cargo.toml)

`pg_chainsync` — Access blockchain blocks, events, and tasks directly inside PostgreSQL using background workers.

The reviewed catalog snapshot records version `0.0.0`, kind `preload`, and implementation language `Rust`.
The curated compatibility set is `17`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_chainsync";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_chainsync';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
