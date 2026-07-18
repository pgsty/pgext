## Usage

Sources:

- [Official extension control file](https://github.com/vishvish/pgjwt/blob/ddaab1e4d47b78c91812dd757f3b878d43e0bad4/pgjwt_rs.control)
- [Official upstream documentation](https://github.com/vishvish/pgjwt/blob/ddaab1e4d47b78c91812dd757f3b878d43e0bad4/README.md)
- [Official Rust package manifest](https://github.com/vishvish/pgjwt/blob/ddaab1e4d47b78c91812dd757f3b878d43e0bad4/Cargo.toml)

`pgjwt_rs` — JWT verification for RS256 and Ed25519 tokens inside PostgreSQL

The reviewed catalog snapshot records version `0.1.2`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pgjwt_rs";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgjwt_rs';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
