## Usage

Sources:

- [Official upstream documentation](https://github.com/glynastill/table_log_pl/blob/564007f69efb810d0dc2123a8080037810283fff/README.md)

`table_log_pl` — PL/pgSQL drop-in replacement for the table_log trigger-based DML auditing extension.

The reviewed catalog snapshot records version `0.1`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plpgsql`.

```sql
CREATE EXTENSION "table_log_pl";
SELECT extversion
FROM pg_extension
WHERE extname = 'table_log_pl';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
