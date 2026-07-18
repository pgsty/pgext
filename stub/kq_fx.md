## Usage

Sources:

- [Official extension control file](https://github.com/ketteq-neon/kq_fx/blob/d62ad40bf4dc61d568801f6ed768c77128256e4b/kq_fx.control)
- [Official upstream documentation](https://github.com/ketteq-neon/kq_fx/blob/d62ad40bf4dc61d568801f6ed768c77128256e4b/README.md)
- [Official Rust package manifest](https://github.com/ketteq-neon/kq_fx/blob/d62ad40bf4dc61d568801f6ed768c77128256e4b/Cargo.toml)

`kq_fx` — ketteQ, Inc.

The reviewed catalog snapshot records version `1.0.1`, kind `preload`, and implementation language `Rust`.
The curated compatibility set is `13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "kq_fx";
SELECT extversion
FROM pg_extension
WHERE extname = 'kq_fx';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
