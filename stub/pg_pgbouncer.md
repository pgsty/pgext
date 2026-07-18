## Usage

Sources:

- [Official Rust package manifest](https://github.com/pgbouncer/pg_pgbouncer/blob/e71619b01d391316c98eb8456e5b0b37c76d96d9/Cargo.toml)

`pg_pgbouncer` — Runs and manages external PgBouncer processes from a PostgreSQL background worker through a pgrx extension.

The reviewed catalog snapshot records version `0.0.0`, kind `preload`, and implementation language `Rust`.
The curated compatibility set is `16`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_pgbouncer";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_pgbouncer';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
