## Usage

Sources:

- [Official extension control file](https://github.com/clowder/pg_logfmt/blob/aaa2da2b71c6864264ad71b8d1d4a198d3fb5d9c/pg_logfmt.control)

`pg_logfmt` — Parse logfmt strings into JSONB and key sets

The reviewed catalog snapshot records version `0.0.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `14,15`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_logfmt";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_logfmt';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
