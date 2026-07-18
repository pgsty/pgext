## Usage

Sources:

- [Official extension control file](https://github.com/cmu-db/pgextmgrext/blob/a412421903e0ddb5fb51a195ab92f70f698bb945/pgextmgr/pgextmgr.control)
- [Official upstream documentation](https://github.com/cmu-db/pgextmgrext/blob/a412421903e0ddb5fb51a195ab92f70f698bb945/README.md)
- [Official Rust package manifest](https://github.com/cmu-db/pgextmgrext/blob/a412421903e0ddb5fb51a195ab92f70f698bb945/pgextmgr/Cargo.toml)

`pgextmgr` — pgrx extension manager and hook framework for registering, ordering, enabling, and disabling PostgreSQL extension plugins.

The reviewed catalog snapshot records version `0.0.0`, kind `preload`, and implementation language `Rust`.
The curated compatibility set is `15`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pgextmgr";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgextmgr';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
