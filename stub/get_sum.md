## Usage

Sources:

- [Official extension control file](https://github.com/idrawone/get_sum/blob/43f8ed114da79321c4125c190044ab4e06e1322e/get_sum.control)

`get_sum` — simple sum of two integers for postgres externsion using c

The reviewed catalog snapshot records version `0.0.1`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "get_sum";
SELECT extversion
FROM pg_extension
WHERE extname = 'get_sum';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
