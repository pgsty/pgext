## Usage

Sources:

- [Official extension control file](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_ledger/pg_ledger.control)
- [Official upstream documentation](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/README.md)
- [Official Rust package manifest](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_ledger/Cargo.toml)

`pg_ledger` — Provides a double-entry accounting engine with immutable journals, fiscal periods, currencies and posting rules inside PostgreSQL.

The reviewed catalog snapshot records version `0.2.0`, kind `standard`, and implementation language `Rust`.
Install and validate the declared extension dependencies first: `plpgsql`.
The curated compatibility set is `14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_ledger";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_ledger';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
