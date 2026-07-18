## Usage

Sources:

- [Official upstream documentation](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_ortools/README.md)
- [Official Rust package manifest](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_ortools/Cargo.toml)

`pg_ortools` — Constraint optimization extension using the HiGHS MIP/LP solver from PostgreSQL, with optional async execution.

The reviewed catalog snapshot records version `0.2.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_ortools";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_ortools';
```

The upstream project is associated with `Matroid`; verify its current support, license, packaging, and deployment boundary from the linked source.

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
