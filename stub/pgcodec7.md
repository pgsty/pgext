## Usage

Sources:

- [Official upstream documentation](https://github.com/AbdulYadi/pgcodec7/blob/b54cdcbadb03042396a65b8db7f9913fb0686fc1/README.md)

`pgcodec7` — C extension that encodes and decodes bytea values using compact 7-bit or 6-bit text grouping.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.
The curated compatibility set is `11`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pgcodec7";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgcodec7';
```

The curated lifecycle is `abandoned`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
