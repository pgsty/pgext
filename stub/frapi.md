## Usage

Sources:

- [Official extension control file](https://github.com/adauhr/pg_frapi/blob/f6ac6decaf54558ed34beda4f35b9f027cb55835/frapi.control)
- [Official upstream documentation](https://github.com/adauhr/pg_frapi/blob/f6ac6decaf54558ed34beda4f35b9f027cb55835/README.md)

`frapi` — In-database wrappers for French public-administration data APIs.

The reviewed catalog snapshot records version `0.1.2`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plsh`, `postgis`.

```sql
CREATE EXTENSION "frapi";
SELECT extversion
FROM pg_extension
WHERE extname = 'frapi';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
