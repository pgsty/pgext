## Usage

Sources:

- [Official extension control file](https://github.com/scottjustin5000/pg-sum-array/blob/4845e1e605773ea51ca51f8044dbd25da2698a0b/sum_array.control)
- [Official upstream documentation](https://github.com/scottjustin5000/pg-sum-array/blob/4845e1e605773ea51ca51f8044dbd25da2698a0b/README.md)

`sum_array` — C function that sums PostgreSQL integer and floating-point arrays into float8.

The reviewed catalog snapshot records version `0.0.1`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "sum_array";
SELECT extversion
FROM pg_extension
WHERE extname = 'sum_array';
```

The curated lifecycle is `abandoned`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
