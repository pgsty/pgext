## Usage

Sources:

- [Official extension control file](https://github.com/choplin/pg-jsonpath/blob/968eff51fcb958b2641fd81cb5739b3c033a7979/jsonpath.control)
- [Official upstream README](https://github.com/choplin/pg-jsonpath/blob/968eff51fcb958b2641fd81cb5739b3c033a7979/README.md)

`jsonpath` — Run JSONPath queries through a bundled JavaScript implementation in PLV8

The reviewed catalog snapshot records version `0.1`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plv8`.

```sql
CREATE EXTENSION "jsonpath";
SELECT extversion
FROM pg_extension
WHERE extname = 'jsonpath';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
