## Usage

Sources:

- [Official extension control file](https://github.com/alexandersamoylov/pg_dbwa/blob/3a8ea55aafb18ed4337be03a3f896e211b40a988/pg_dbwa.control)
- [Official upstream documentation](https://github.com/alexandersamoylov/pg_dbwa/blob/3a8ea55aafb18ed4337be03a3f896e211b40a988/README.md)

`pg_dbwa` — Analyzes local or remote PostgreSQL workloads using historical pg_stat_statements snapshots.

The reviewed catalog snapshot records version `0.3.1`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `dblink`, `pg_eyes`, `pg_prttn_tools`, `pg_stat_statements`, `plpgsql`.

```sql
CREATE EXTENSION "pg_dbwa";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_dbwa';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
