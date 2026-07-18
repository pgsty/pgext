## Usage

Sources:

- [Official upstream documentation](https://github.com/ibotty/pg-row-hashes/blob/a654208eba84d767f8d25656ab01d154a86d594c/README.md)
- [Official Rust package manifest](https://github.com/ibotty/pg-row-hashes/blob/a654208eba84d767f8d25656ab01d154a86d594c/Cargo.toml)

`pg_row_hashes` — Computes deterministic row fingerprints, checksums, and XOR aggregates with SeaHash and FarmHash variants.

The reviewed catalog snapshot records version `0.3.2`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_row_hashes";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_row_hashes';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
