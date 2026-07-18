## Usage

Sources:

- [Official extension control file](https://github.com/matheus-consoli/boom-index-pgrx/blob/7025b8186d770eee48ecff55f683c50355d98ff4/bloom_fulon.control)
- [Official Rust package manifest](https://github.com/matheus-consoli/boom-index-pgrx/blob/7025b8186d770eee48ecff55f683c50355d98ff4/Cargo.toml)

`bloom_fulon` — Experimental pgrx index access method for PostgreSQL, created for learning purposes.

The reviewed catalog snapshot records version `0.0.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `11,12,13,14,15,16`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "bloom_fulon";
SELECT extversion
FROM pg_extension
WHERE extname = 'bloom_fulon';
```

The curated lifecycle is `preview`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
