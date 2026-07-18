## Usage

Sources:

- [Official PGXN distribution page](https://pgxn.org/dist/pg_icu_parser/)

`pg_icu_parser` — Full text search parser using ICU boundary analysis

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_icu_parser";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_icu_parser';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
