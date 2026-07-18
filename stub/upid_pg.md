## Usage

Sources:

- [Official extension control file](https://github.com/carderne/upid/blob/d820a74227adb803d378d89da9eacba78c1a5e09/upid_pg/upid_pg.control)
- [Official upstream documentation](https://github.com/carderne/upid/blob/d820a74227adb803d378d89da9eacba78c1a5e09/README.md)
- [Official Rust package manifest](https://github.com/carderne/upid/blob/d820a74227adb803d378d89da9eacba78c1a5e09/upid_pg/Cargo.toml)

`upid_pg` — Prefixed, sortable 128-bit identifier type and generator

The reviewed catalog snapshot records version `0.0.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "upid_pg";
SELECT extversion
FROM pg_extension
WHERE extname = 'upid_pg';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
