## Usage

Sources:

- [Official extension control file](https://github.com/pierreforstmann/pg_logqueryid/blob/306da69ba0dc73f49ab08feaf61b83e5f521cdd1/pg_logqueryid.control)
- [Official PGXN distribution page](https://pgxn.org/dist/pg_logqueryid/)

`pg_logqueryid` — Logs pg_stat_statements query IDs alongside auto_explain output.

The reviewed catalog snapshot records version `0.0.1`, kind `preload`, and implementation language `C`.
Install and validate the declared extension dependencies first: `auto_explain`, `pg_stat_statements`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_logqueryid";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_logqueryid';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
