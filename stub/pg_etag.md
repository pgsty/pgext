## Usage

Sources:

- [Official extension control file](https://github.com/ergo70/pg_etag/blob/444c21d9df4475179d9ee40ec07fcf51a48e5501/pg_etag.control)
- [Official upstream documentation](https://github.com/ergo70/pg_etag/blob/444c21d9df4475179d9ee40ec07fcf51a48e5501/README.md)

`pg_etag` — Fast ETag generation for rows and result sets using BLAKE2 hashes, with single-input functions and aggregates.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_etag";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_etag';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
