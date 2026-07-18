## Usage

Sources:

- [Official extension control file](https://github.com/pgsql-io/hive_fdw/blob/600587892b0a811c129c81d19e01b3b4dd1c87b2/hive_fdw.control)
- [Official upstream documentation](https://github.com/pgsql-io/hive_fdw/blob/600587892b0a811c129c81d19e01b3b4dd1c87b2/README.md)

`hive_fdw` — Foreign data wrapper for querying Apache Hive from PostgreSQL.

The reviewed catalog snapshot records version `4.0`, kind `standard`, and implementation language `C`.
The curated compatibility set is `13`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "hive_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'hive_fdw';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
