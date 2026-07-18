## Usage

Sources:

- [Official extension control file](https://github.com/lij55/pgrand/blob/607b698255007094ca5c44ab4e2bda1ad8d03558/pgrand.control)
- [Official Rust package manifest](https://github.com/lij55/pgrand/blob/607b698255007094ca5c44ab4e2bda1ad8d03558/Cargo.toml)

`pgrand` — Generates random test data through a PostgreSQL FDW, with an experimental table access method.

The reviewed catalog snapshot records version `0.1.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `15,16`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pgrand";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgrand';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
