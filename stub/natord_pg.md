## Usage

Sources:

- [Official extension control file](https://github.com/mpajkowski/natord_pg/blob/a1f2923934a1ae4179e559ea651641f1b975bef0/natord_pg.control)
- [Official Rust package manifest](https://github.com/mpajkowski/natord_pg/blob/a1f2923934a1ae4179e559ea651641f1b975bef0/Cargo.toml)

`natord_pg` — Natural-order comparison operators and btree operator class for PostgreSQL text values.

The reviewed catalog snapshot records version `0.0.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `13,14,15,16,17`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "natord_pg";
SELECT extversion
FROM pg_extension
WHERE extname = 'natord_pg';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
