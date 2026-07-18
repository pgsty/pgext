## Usage

Sources:

- [Official extension control file](https://github.com/asotolongo/db_info/blob/cfd4110846932f74e403c987449f70011ccf61d4/db_info.control)
- [Official upstream documentation](https://github.com/asotolongo/db_info/blob/cfd4110846932f74e403c987449f70011ccf61d4/README.md)
- [Official PGXN distribution page](https://pgxn.org/dist/db_info/)

`db_info` — Functions and views reporting database size, ownership, tablespaces, extensions, and object counts.

The reviewed catalog snapshot records version `0.2.0`, kind `puresql`, and implementation language `SQL`.
The curated compatibility set is `11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "db_info";
SELECT extversion
FROM pg_extension
WHERE extname = 'db_info';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
