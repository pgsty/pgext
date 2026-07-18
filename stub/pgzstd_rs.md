## Usage

Sources:

- [Official extension control file](https://github.com/stella3d/pgzstd_rs/blob/d9b56279f41f1375e9db6a3d397c89346dade5f8/pgzstd_rs.control)
- [Official upstream documentation](https://github.com/stella3d/pgzstd_rs/blob/d9b56279f41f1375e9db6a3d397c89346dade5f8/README.md)
- [Official Rust package manifest](https://github.com/stella3d/pgzstd_rs/blob/d9b56279f41f1375e9db6a3d397c89346dade5f8/Cargo.toml)

`pgzstd_rs` — Compress and decompress bytea values with Zstandard from PostgreSQL, including a parallel batch compressor

The reviewed catalog snapshot records version `0.0.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `11,12,13,14,15,16`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pgzstd_rs";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgzstd_rs';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
