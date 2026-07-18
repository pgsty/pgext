## Usage

Sources:

- [Official extension control file](https://github.com/okbob/orafce_sql/blob/76f311c75074131f4d076bf28981e5f97bc8bf6a/dbms_sql.control)
- [Official upstream documentation](https://github.com/okbob/orafce_sql/blob/76f311c75074131f4d076bf28981e5f97bc8bf6a/README.md)

`dbms_sql` — Implements a subset of Oracle's DBMS_SQL package API to ease application migration to PostgreSQL.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.
The curated compatibility set is `11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "dbms_sql";
SELECT extversion
FROM pg_extension
WHERE extname = 'dbms_sql';
```

The curated lifecycle is `archived`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
