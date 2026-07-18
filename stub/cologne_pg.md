## Usage

Sources:

- [Official extension control file](https://github.com/DrSloth/cologne_pg/blob/6871c07bb24c9a05e63b5e2b034bf34eda64426c/cologne_pg.control)
- [Official Rust package manifest](https://github.com/DrSloth/cologne_pg/blob/6871c07bb24c9a05e63b5e2b034bf34eda64426c/Cargo.toml)

`cologne_pg` — Cologne phonetics function for PostgreSQL built with pgrx.

The reviewed catalog snapshot records version `0.0.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `11,12,13,14,15,16`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "cologne_pg";
SELECT extversion
FROM pg_extension
WHERE extname = 'cologne_pg';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
