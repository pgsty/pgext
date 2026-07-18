## Usage

Sources:

- [Official PGXN distribution page](https://pgxn.org/dist/pg_htmldoc/)

`pg_htmldoc` — Converts HTML, Markdown, files or web pages to PDF, PostScript and other HTMLDOC output formats from PostgreSQL.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_htmldoc";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_htmldoc';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
