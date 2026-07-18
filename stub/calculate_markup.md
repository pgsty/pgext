## Usage

Sources:

- [Official extension control file](https://github.com/chobostar/calculate_markup/blob/c70a994c8a215e11d3fd444f0d1412289bcadf3b/calculate_markup.control)

`calculate_markup` — C function that calculates a price markup percentage by linear interpolation over a numeric price/markup array.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "calculate_markup";
SELECT extversion
FROM pg_extension
WHERE extname = 'calculate_markup';
```

The curated lifecycle is `archived`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
