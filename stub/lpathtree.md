## Usage

Sources:

- [Official extension control file](https://github.com/mediait/lpathtree/blob/6b1fd0acd1d4a698d9f4b04d59eb4e144a1eda15/lpathtree.control)

`lpathtree` — Hierarchical path types using slash-separated labels

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "lpathtree";
SELECT extversion
FROM pg_extension
WHERE extname = 'lpathtree';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
