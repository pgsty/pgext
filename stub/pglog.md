## Usage

Sources:

- [Official extension control file](https://github.com/2ndQuadrant/pglog/blob/5d57809e4e73913a06b4aaa6553f53f43d010a39/pglog.control)

`pglog` — PostgreSQL log via SQL

The reviewed catalog snapshot records version `1.0`, kind `preload`, and implementation language `C`.

```sql
CREATE EXTENSION "pglog";
SELECT extversion
FROM pg_extension
WHERE extname = 'pglog';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
