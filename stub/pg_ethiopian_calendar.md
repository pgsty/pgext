## Usage

Sources:

- [Official PGXN distribution page](https://pgxn.org/dist/pg_ethiopian_calendar/)

`pg_ethiopian_calendar` — Convert between Gregorian and Ethiopian calendar dates

The reviewed catalog snapshot records version `1.1`, kind `standard`, and implementation language `C`.
The curated compatibility set is `11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_ethiopian_calendar";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_ethiopian_calendar';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
