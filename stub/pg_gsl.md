## Usage

Sources:

- [Official PGXN distribution page](https://pgxn.org/dist/pg_gsl/)

`pg_gsl` — Wraps selected GNU Scientific Library routines, including real FFT transforms and helper functions, for PostgreSQL.

The reviewed catalog snapshot records version `0.0.2`, kind `standard`, and implementation language `C`.
Install and validate the declared extension dependencies first: `plpgsql`.

```sql
CREATE EXTENSION "pg_gsl";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_gsl';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
