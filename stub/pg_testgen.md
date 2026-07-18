## Usage

Sources:

- [Official extension control file](https://github.com/yuesong-feng/pg_testgen/blob/5a6c127a1596eb8fa678fc0df6608ee9d83cd115/pg_testgen.control)
- [Official upstream documentation](https://github.com/yuesong-feng/pg_testgen/blob/5a6c127a1596eb8fa678fc0df6608ee9d83cd115/README.md)

`pg_testgen` — Generate random integers, strings, and test-data rows for PostgreSQL development

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_testgen";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_testgen';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
