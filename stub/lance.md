## Usage

Sources:

- [Official extension control file](https://github.com/lance-format/pglance/blob/6dece2f725624f3496816135d08b9b4eec63b219/lance.control)
- [Official upstream documentation](https://github.com/lance-format/pglance/blob/6dece2f725624f3496816135d08b9b4eec63b219/README.md)
- [Official Rust package manifest](https://github.com/lance-format/pglance/blob/6dece2f725624f3496816135d08b9b4eec63b219/Cargo.toml)

`lance` — PostgreSQL FDW extension for reading and querying Lance format tables

The reviewed catalog snapshot records version `0.0.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `13,14,15,16,17`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "lance";
SELECT extversion
FROM pg_extension
WHERE extname = 'lance';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
