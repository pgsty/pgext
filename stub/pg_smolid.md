## Usage

Sources:

- [Official Rust package manifest](https://github.com/dotvezz/pg_smolid/blob/b9840b6764545c4e59353506cfa4bcf2cfb2bf4f/Cargo.toml)

`pg_smolid` — A Postgres pgrx extension adding smolid type and helper functions.

The reviewed catalog snapshot records version `0.1.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_smolid";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_smolid';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
