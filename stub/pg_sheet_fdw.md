## Usage

Sources:

- [Official PGXN distribution page](https://pgxn.org/dist/pg_sheet_fdw/)

`pg_sheet_fdw` — Read local Excel XLSX worksheets as PostgreSQL foreign tables

The reviewed catalog snapshot records version `0.1.2`, kind `standard`, and implementation language `C`.
The curated compatibility set is `13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_sheet_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_sheet_fdw';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
