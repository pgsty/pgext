## Usage

Sources:

- [Official extension control file](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_prob/pg_prob.control)
- [Official Rust package manifest](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_prob/Cargo.toml)

`pg_prob` — Probabilistic data types for PostgreSQL - Monte Carlo simulation in SQL

The reviewed catalog snapshot records version `0.2.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_prob";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_prob';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
