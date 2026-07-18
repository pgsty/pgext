## Usage

Sources:

- [Official Rust package manifest](https://github.com/zvectorlabs/pgauthz/blob/a24a3bae33d1bac320c233e99ceb9153a4d0667d/crates/pgauthz/Cargo.toml)

`pgauthz` — Provide Zanzibar-style relationship-based authorization inside PostgreSQL

The reviewed catalog snapshot records version `0.3.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pgauthz";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgauthz';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
