## Usage

Sources:

- [Official upstream documentation](https://github.com/pandrewhk/pgcurl/blob/97d1e7e7447adae345fd90695916031237e1b06b/README.md)

`pgcurl` — Calls curl/libcurl from PostgreSQL through a SQL function.

The reviewed catalog snapshot records version `0.0.1`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pgcurl";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgcurl';
```

The curated lifecycle is `abandoned`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
