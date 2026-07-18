## Usage

Sources:

- [Official extension control file](https://github.com/uptimejp/pg_stat_statements2/blob/db2e94c49c9be23137f6a8925616535950f2ae89/pg_stat_statements2.control)
- [Official upstream documentation](https://github.com/uptimejp/pg_stat_statements2/blob/db2e94c49c9be23137f6a8925616535950f2ae89/README)

`pg_stat_statements2` — Track SQL execution statistics using a PostgreSQL 9.4 pg_stat_statements derivative integrated with sql_firewall

The reviewed catalog snapshot records version `1.2`, kind `preload`, and implementation language `C`.
Install and validate the declared extension dependencies first: `sql_firewall`.

```sql
CREATE EXTENSION "pg_stat_statements2";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_stat_statements2';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
