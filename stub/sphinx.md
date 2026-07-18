## Usage

Sources:

- [Official extension control file](https://github.com/andy128k/pg-sphinx/blob/6a63ade3a0f7119474435d92b570d580638b4d90/sphinx.control)
- [Official upstream documentation](https://github.com/andy128k/pg-sphinx/blob/6a63ade3a0f7119474435d92b570d580638b4d90/README.org)

`sphinx` — C connector for querying and updating an external SphinxSearch server from PostgreSQL.

The reviewed catalog snapshot records version `0.3`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "sphinx";
SELECT extversion
FROM pg_extension
WHERE extname = 'sphinx';
```

The curated lifecycle is `archived`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
