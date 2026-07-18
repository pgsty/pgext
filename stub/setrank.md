## Usage

Sources:

- [Official extension control file](https://github.com/obartunov/setrank/blob/b8492a5968c33aa1c2d87d6aecf4a35477d75e10/setrank.control)
- [Official upstream documentation](https://github.com/obartunov/setrank/blob/b8492a5968c33aa1c2d87d6aecf4a35477d75e10/README)

`setrank` — TF-IDF and cover-density ranking functions for tsvector/tsquery using document-frequency statistics from a table.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "setrank";
SELECT extversion
FROM pg_extension
WHERE extname = 'setrank';
```

The curated lifecycle is `abandoned`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
