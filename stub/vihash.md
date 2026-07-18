## Usage

Sources:

- [Official upstream documentation](https://github.com/petere/pgvihash/blob/70e637321ab07746ecaebf6c47385b06c84c050c/README)
- [Official PGXN distribution page](https://pgxn.org/dist/pgvihash/)

`vihash` — version-independent hash functions

The reviewed catalog snapshot records version `1`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "vihash";
SELECT extversion
FROM pg_extension
WHERE extname = 'vihash';
```

The curated lifecycle is `archived`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
