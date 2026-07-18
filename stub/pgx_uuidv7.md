## Usage

Sources:

- [Official extension control file](https://github.com/kaznak/pgx_uuidv7/blob/3b3828ae89d634c2e828cc872a1ac96394e6cf60/pgx_uuidv7.control)
- [Official upstream documentation](https://github.com/kaznak/pgx_uuidv7/blob/3b3828ae89d634c2e828cc872a1ac96394e6cf60/README.md)
- [Official Rust package manifest](https://github.com/kaznak/pgx_uuidv7/blob/3b3828ae89d634c2e828cc872a1ac96394e6cf60/Cargo.toml)

`pgx_uuidv7` — A PostgreSQL extension to generate and cast v7 UUIDs

The reviewed catalog snapshot records version `0.1.7`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pgx_uuidv7";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgx_uuidv7';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
