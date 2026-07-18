## Usage

Sources:

- [Official extension control file](https://github.com/nate1001/chess_index/blob/4d6d473ea9c8238f06d8759f65862a382621f33e/chess_index.control)

`chess_index` — Data types for indexing chess positions.

The reviewed catalog snapshot records version `0.0.1`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "chess_index";
SELECT extversion
FROM pg_extension
WHERE extname = 'chess_index';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
