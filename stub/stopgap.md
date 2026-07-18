## Usage

Sources:

- [Official upstream documentation](https://github.com/Smertos/stopgap/blob/e076ae165ba35c6f2d484acd943451373376e03e/README.md)
- [Official Rust package manifest](https://github.com/Smertos/stopgap/blob/e076ae165ba35c6f2d484acd943451373376e03e/crates/stopgap/Cargo.toml)

`stopgap` — pgrx deployment layer for versioned TypeScript/JavaScript function bundles backed by PLTS.

The reviewed catalog snapshot records version `0.1.3`, kind `standard`, and implementation language `Rust`.
Install and validate the declared extension dependencies first: `plts`.
The curated compatibility set is `14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "stopgap";
SELECT extversion
FROM pg_extension
WHERE extname = 'stopgap';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
