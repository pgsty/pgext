## Usage

Sources:

- [Official extension control file](https://github.com/KaiserKarel/pg_bech32/blob/main/pg_bech32.control)
- [Official upstream documentation](https://github.com/KaiserKarel/pg_bech32/blob/main/README.md)
- [Official Rust package manifest](https://github.com/KaiserKarel/pg_bech32/blob/a4eb4aaae249540b9839cd8735cb450bad1ee8eb/Cargo.toml)

`pg_bech32` — Adds Bech32, Bech32m, and no-checksum Bech32 encoding/decoding support to PostgreSQL.

The reviewed catalog snapshot records version `0.0.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `11,12,13,14,15,16`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_bech32";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_bech32';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
