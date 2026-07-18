## Usage

Sources:

- [Official extension control file](https://github.com/chet0xhenry/pgcompress/blob/8264107a26c07b6415f75c4a8eb84f9d27e77ea3/pgcompress.control)
- [Official upstream documentation](https://github.com/chet0xhenry/pgcompress/blob/8264107a26c07b6415f75c4a8eb84f9d27e77ea3/README.md)

`pgcompress` — Provides deflate, inflate, gzip, and Brotli compression functions for bytea, text, json, and jsonb.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pgcompress";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgcompress';
```

The curated lifecycle is `abandoned`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
