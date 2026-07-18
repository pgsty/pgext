## Usage

Sources:

- [Official extension control file](https://github.com/VitoVan/pg_opencc/blob/c919ba0cff11949c682c583d80f72dfb15eca87a/pg_opencc.control)
- [Official upstream documentation](https://github.com/VitoVan/pg_opencc/blob/c919ba0cff11949c682c583d80f72dfb15eca87a/README.md)

`pg_opencc` — Converts text among Simplified Chinese, Traditional Chinese regional variants and Japanese Shinjitai using OpenCC.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.
The curated compatibility set is `13`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_opencc";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_opencc';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
