## Usage

Sources:

- [Official extension control file](https://github.com/shestero/pglas/blob/efc9f4e89585ac773550c2093782d65d1049a466/pglas.control)

`pglas` — Reads LAS 2.0 well logging files from PostgreSQL functions

The reviewed catalog snapshot records version `0.1.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `11,12,13,14,15`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pglas";
SELECT extversion
FROM pg_extension
WHERE extname = 'pglas';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
