## Usage

Sources:

- [Official PGXN distribution page](https://pgxn.org/dist/pg_wkhtmltopdf/)

`pg_wkhtmltopdf` — Convert HTML pages to PDF bytea values inside PostgreSQL using wkhtmltopdf/QtWebKit.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_wkhtmltopdf";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_wkhtmltopdf';
```

The curated lifecycle is `abandoned`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
