## Usage

Sources:

- [Official upstream README](https://github.com/maludb/maludb-core/blob/b4d6d521cb94cb0c05cfbadd2a9958c2e5ddfc4f/README.md)
- [Official extension control file (maludb_core.control)](https://github.com/maludb/maludb-core/blob/b4d6d521cb94cb0c05cfbadd2a9958c2e5ddfc4f/maludb_core.control)
- [Official implementation source](https://github.com/maludb/maludb-core/blob/b4d6d521cb94cb0c05cfbadd2a9958c2e5ddfc4f/src/maludb_core.c)

`maludb_core` — MaluDB is a memory DBMS for long-term institutional memory, human-AI knowledge sharing, and contextual recall. Built in **C** as PostgreSQL extensions on **Ubuntu 24.04 LTS**, with **PostgreSQL 17** (PGDG) as the foundation. Use it for the corresponding vector, model, or retrieval workflow. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION maludb_core;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.104.0`.
- Install the confirmed extension dependencies first: `vector`, `btree_gist`, `pg_trgm`, `pgcrypto`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
