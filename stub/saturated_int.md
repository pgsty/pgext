## Usage

Sources:

- [Official extension control file](https://github.com/adjust/pg-saturated_int/blob/a1228ed87d0b27a365eb6d0b241a27f654f49156/saturated_int.control)
- [Official upstream documentation](https://github.com/adjust/pg-saturated_int/blob/a1228ed87d0b27a365eb6d0b241a27f654f49156/README.md)

`saturated_int` — Integer data type with saturation arithmetic that clamps overflow to int4 bounds.

The reviewed catalog snapshot records version `0.0.1`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "saturated_int";
SELECT extversion
FROM pg_extension
WHERE extname = 'saturated_int';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
