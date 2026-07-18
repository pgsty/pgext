## Usage

Sources:

- [Official upstream README](https://github.com/RekGRpth/pg_wthtmltopdf/blob/dc68969a4821945d7743cd4ab4b93258dd365251/README.md)

`pg_wthtmltopdf` — Render HTML into PDF from a PostgreSQL function.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C++`.

```sql
CREATE EXTENSION "pg_wthtmltopdf";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_wthtmltopdf';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
