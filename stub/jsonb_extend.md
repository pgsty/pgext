## Usage

Sources:

- [Official upstream documentation](https://github.com/koctep/jsonb-extend/blob/fb704ad2b19a8d61d7cc0bcd50ecf507c629dba1/README.md)

`jsonb_extend` — Merge two or more jsonb values

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "jsonb_extend";
SELECT extversion
FROM pg_extension
WHERE extname = 'jsonb_extend';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
