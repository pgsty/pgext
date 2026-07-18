## Usage

Sources:

- [Official PGXN distribution page](https://pgxn.org/dist/query_recorder/)

`query_recorder` — query_recorder: record executed SQL queries to rotating files for later analysis or replay

The reviewed catalog snapshot records version `1.0.1`, kind `preload`, and implementation language `C`.

```sql
CREATE EXTENSION "query_recorder";
SELECT extversion
FROM pg_extension
WHERE extname = 'query_recorder';
```

The curated lifecycle is `preview`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
