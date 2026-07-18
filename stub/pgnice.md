## Usage

Sources:

- [Official extension control file](https://github.com/mselee/pgnice/blob/3abc7e1d9449097c17eae5346021f0c2ea576fce/pgnice.control)
- [Official upstream documentation](https://github.com/mselee/pgnice/blob/3abc7e1d9449097c17eae5346021f0c2ea576fce/README.md)
- [Official Rust package manifest](https://github.com/mselee/pgnice/blob/3abc7e1d9449097c17eae5346021f0c2ea576fce/Cargo.toml)

`pgnice` — Manage PostgreSQL backend process priorities and resource limits.

The reviewed catalog snapshot records version `0.0.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `13,14,15,16,17`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pgnice";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgnice';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
