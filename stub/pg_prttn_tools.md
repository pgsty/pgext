## Usage

Sources:

- [Official extension control file](https://github.com/alexandersamoylov/pg_prttn_tools/blob/c7a034b56cd8c2327779dbc0d8060fd2220eb010/pg_prttn_tools.control)

`pg_prttn_tools` — Provides PL/pgSQL tools for managing inheritance-based partition tables.

The reviewed catalog snapshot records version `0.5`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plpgsql`.

```sql
CREATE EXTENSION "pg_prttn_tools";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_prttn_tools';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
