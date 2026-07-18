## Usage

Sources:

- [Official extension control file](https://github.com/RekGRpth/pg_mupdf/blob/a15db6ca07c5e95b8f2d78fe92774c659a3f3c54/pg_mupdf.control)
- [Official PGXN distribution page](https://pgxn.org/dist/pg_mupdf/)

`pg_mupdf` — Convert HTML to PDF inside PostgreSQL using MuPDF

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_mupdf";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_mupdf';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
