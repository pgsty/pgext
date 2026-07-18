## Usage

Sources:

- [Official extension control file](https://github.com/DanielJDufour/safecast/blob/0ad39b4f99d373f309817ba803374b5b2c9229a3/safecast.control)
- [Official upstream documentation](https://github.com/DanielJDufour/safecast/blob/0ad39b4f99d373f309817ba803374b5b2c9229a3/README.md)

`safecast` — Pure-SQL helper functions that return NULL for some invalid text-to-number casts.

The reviewed catalog snapshot records version `0.0.1`, kind `puresql`, and implementation language `SQL`.

```sql
CREATE EXTENSION "safecast";
SELECT extversion
FROM pg_extension
WHERE extname = 'safecast';
```

The curated lifecycle is `abandoned`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
