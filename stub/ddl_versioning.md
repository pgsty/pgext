## Usage

Sources:

- [Official extension control file](https://github.com/filiprem/pg-tools/blob/3ad4b21874bed6b629f984d046650de354e95d92/ddl_versioning/ddl_versioning.control)
- [Official upstream documentation](https://github.com/filiprem/pg-tools/blob/3ad4b21874bed6b629f984d046650de354e95d92/ddl_versioning/README.ddl_versioning)
- [Official upstream README](https://github.com/filiprem/pg-tools/blob/3ad4b21874bed6b629f984d046650de354e95d92/README.md)

`ddl_versioning` — DDL object versioning with event triggers for tables, indexes, functions, and views.

The reviewed catalog snapshot records version `1.0`, kind `puresql`, and implementation language `SQL`.
The curated compatibility set is `13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "ddl_versioning";
SELECT extversion
FROM pg_extension
WHERE extname = 'ddl_versioning';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
