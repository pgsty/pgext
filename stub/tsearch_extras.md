## Usage

Sources:

- [Official extension control file](https://github.com/zulip/tsearch_extras/blob/f566e9606bac00a22c4b62e9511b022c861db05c/tsearch_extras.control)
- [Official upstream documentation](https://github.com/zulip/tsearch_extras/blob/f566e9606bac00a22c4b62e9511b022c861db05c/README.md)

`tsearch_extras` — Full-text-search helpers for returning match locations and extracting tsvector lexemes.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "tsearch_extras";
SELECT extversion
FROM pg_extension
WHERE extname = 'tsearch_extras';
```

The curated lifecycle is `archived`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
