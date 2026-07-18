## Usage

Sources:

- [Official extension control file](https://github.com/cmu-db/pgextmgrext/blob/a412421903e0ddb5fb51a195ab92f70f698bb945/pgx_trace_hooks/pgx_trace_hooks.control)
- [Official Rust package manifest](https://github.com/cmu-db/pgextmgrext/blob/a412421903e0ddb5fb51a195ab92f70f698bb945/pgx_trace_hooks/Cargo.toml)
- [Official upstream README](https://github.com/cmu-db/pgextmgrext/blob/a412421903e0ddb5fb51a195ab92f70f698bb945/README.md)

`pgx_trace_hooks` — Trace PostgreSQL executor and utility-hook invocations through a preloaded pgrx library

The reviewed catalog snapshot records version `0.0.0`, kind `preload`, and implementation language `Rust`.
The curated compatibility set is `15`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pgx_trace_hooks";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgx_trace_hooks';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
