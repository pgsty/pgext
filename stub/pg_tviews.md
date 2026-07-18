## Usage

Sources:

- [Official extension control file](https://github.com/fraiseql/pg_tviews/blob/5f4288db5d1fdc565b01468ec3a5c9ca4124c447/pg_tviews.control)
- [Official upstream documentation](https://github.com/fraiseql/pg_tviews/blob/5f4288db5d1fdc565b01468ec3a5c9ca4124c447/README.md)
- [Official Rust package manifest](https://github.com/fraiseql/pg_tviews/blob/5f4288db5d1fdc565b01468ec3a5c9ca4124c447/Cargo.toml)

`pg_tviews` — Transactional materialized views with automatic incremental refresh for PostgreSQL

The reviewed catalog snapshot records version `0.1.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_tviews";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_tviews';
```

The curated lifecycle is `preview`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
