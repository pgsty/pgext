## Usage

Sources:

- [Official extension control file](https://github.com/shurizzle/pg_url/blob/7fe5192e25b68e2c22b096b74daf84c26bf5e384/url.control)
- [Official upstream documentation](https://github.com/shurizzle/pg_url/blob/7fe5192e25b68e2c22b096b74daf84c26bf5e384/README.md)

`url` — SQL and PL/pgSQL types and helper functions for URL parsing, encoding, decoding, and query-string key/value handling.

The reviewed catalog snapshot records version `1.0`, kind `puresql`, and implementation language `SQL`.

```sql
CREATE EXTENSION "url";
SELECT extversion
FROM pg_extension
WHERE extname = 'url';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
