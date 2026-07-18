## Usage

Sources:

- [Official extension control file](https://github.com/bmuratshin/zcurve/blob/76b88d80794a4ccde95c01111453ad8d7fbdbf36/zcurve.control)

`zcurve` — Two- and three-dimensional Z-order bit interleaving plus direct range lookup over numeric B-tree indexes.

The reviewed catalog snapshot records version `1.3`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "zcurve";
SELECT extversion
FROM pg_extension
WHERE extname = 'zcurve';
```

The curated lifecycle is `abandoned`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
