## Usage

Sources:

- [Official PGXN distribution page](https://pgxn.org/dist/pg_cld2/)

`pg_cld2` — Detects likely human language and script in PostgreSQL using CLD2.

The reviewed catalog snapshot records version `1.0.0`, kind `standard`, and implementation language `C++`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_cld2";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_cld2';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
