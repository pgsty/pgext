## Usage

Sources:

- [Official extension control file](https://github.com/Preet-Govind/pg_https/blob/cb3d887433b7118bd42f3a216b16db4ae2788033/pg_https.control)
- [Official upstream documentation](https://github.com/Preet-Govind/pg_https/blob/cb3d887433b7118bd42f3a216b16db4ae2788033/readme.md)

`pg_https` — HTTP/HTTPS client functions for PostgreSQL using libcurl

The reviewed catalog snapshot records version `1.1`, kind `preload`, and implementation language `C`.
The curated compatibility set is `17`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_https";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_https';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
