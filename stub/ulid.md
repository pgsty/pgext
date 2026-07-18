## Usage

Sources:

- [Official upstream documentation](https://github.com/Creatif/creatif-backend/blob/7b0466a2b0f505c824d98c2276107531ff1af2c3/pgx_ulid/README.md)
- [Official Rust package manifest](https://github.com/Creatif/creatif-backend/blob/7b0466a2b0f505c824d98c2276107531ff1af2c3/pgx_ulid/Cargo.toml)
- [Official project or provider page](https://github.com/Creatif/creatif-backend/tree/master/pgx_ulid)

`ulid` — Embedded pgrx ULID type, generators, and casts

The reviewed catalog snapshot records version `0.1.3`, kind `preload`, and implementation language `Rust`.
The curated compatibility set is `14,15,16`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "ulid";
SELECT extversion
FROM pg_extension
WHERE extname = 'ulid';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
