## Usage

Sources:

- [Official extension control file](https://github.com/postgrespro/tsexact/blob/b08a6ce7ed40b5b62cccdb12d7e138c653d2efd0/tsexact.control)
- [Official upstream documentation](https://github.com/postgrespro/tsexact/blob/b08a6ce7ed40b5b62cccdb12d7e138c653d2efd0/README.md)

`tsexact` — Full-text-search helpers for exact phrase matching on PostgreSQL 9.5 and earlier.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "tsexact";
SELECT extversion
FROM pg_extension
WHERE extname = 'tsexact';
```

The curated lifecycle is `deprecated`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
