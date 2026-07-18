## Usage

Sources:

- [Official extension control file](https://github.com/le0pard/pg_web/blob/6e936f451090482bdab8d6f749ccccd4c1cbbc60/pg_web.control)
- [Official upstream documentation](https://github.com/le0pard/pg_web/blob/6e936f451090482bdab8d6f749ccccd4c1cbbc60/README.md)

`pg_web` — Run an embedded HTTP server from a PostgreSQL background worker

The reviewed catalog snapshot records version `0.1.0`, kind `preload`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_web";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_web';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
