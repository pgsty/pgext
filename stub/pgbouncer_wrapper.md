## Usage

Sources:

- [Official extension control file](https://github.com/davidfetter/pgbouncer_wrapper/blob/07be9547479c765588098843671abdcadbf45e33/pgbouncer_wrapper.control)
- [Official upstream documentation](https://github.com/davidfetter/pgbouncer_wrapper/blob/07be9547479c765588098843671abdcadbf45e33/README.md)

`pgbouncer_wrapper` — Monitor and control the PgBouncer administration console through SQL views and functions backed by dblink_fdw

The reviewed catalog snapshot records version `1.2.2`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `dblink`.

```sql
CREATE EXTENSION "pgbouncer_wrapper";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgbouncer_wrapper';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
