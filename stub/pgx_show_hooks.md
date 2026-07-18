## Usage

Sources:

- [Official extension control file](https://github.com/cmu-db/pgextmgrext/blob/a412421903e0ddb5fb51a195ab92f70f698bb945/pgx_show_hooks/pgx_show_hooks.control)
- [Official upstream documentation](https://github.com/cmu-db/pgextmgrext/blob/a412421903e0ddb5fb51a195ab92f70f698bb945/README.md)
- [Official Rust package manifest](https://github.com/cmu-db/pgextmgrext/blob/a412421903e0ddb5fb51a195ab92f70f698bb945/pgx_show_hooks/Cargo.toml)

`pgx_show_hooks` — Lists the in-process addresses of PostgreSQL hook functions and PL/pgSQL hook callbacks.

The reviewed catalog snapshot records version `0.0.3`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `15`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pgx_show_hooks";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgx_show_hooks';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
