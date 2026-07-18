## Usage

Sources:

- [Official extension control file](https://github.com/postgrespro/ptrack/blob/62bd42d7308ec7ed786abf2ddc1af24fc3dd72dd/ptrack.control)
- [Official upstream documentation](https://github.com/postgrespro/ptrack/blob/62bd42d7308ec7ed786abf2ddc1af24fc3dd72dd/README.md)

`ptrack` — Tracks changed PostgreSQL blocks for incremental backup using a shared-memory map and core storage hooks.

The reviewed catalog snapshot records version `2.4`, kind `preload`, and implementation language `C`.
The curated compatibility set is `11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "ptrack";
SELECT extversion
FROM pg_extension
WHERE extname = 'ptrack';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
