## Usage

Sources:

- [Official extension control file](https://github.com/asotolongo/pgrgraphic/blob/ff2a6ddcf42d4d73be4ec3c694b3e330ffcb3c0c/pgrgraphic.control)
- [Official upstream documentation](https://github.com/asotolongo/pgrgraphic/blob/ff2a6ddcf42d4d73be4ec3c694b3e330ffcb3c0c/README.md)

`pgrgraphic` — PL/R-based PostgreSQL extension that renders database query results as PNG or JPEG chart files.

The reviewed catalog snapshot records version `1.0.0`, kind `puresql`, and implementation language `R`.
Install and validate the declared extension dependencies first: `plr`.

```sql
CREATE EXTENSION "pgrgraphic";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgrgraphic';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
