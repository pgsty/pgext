## Usage

Sources:

- [Official upstream documentation](https://github.com/hailelagi/pg-sweep-stats/blob/12bc23a94f15db0e55298ccb51364637ac959326/README.md)
- [Official Rust package manifest](https://github.com/hailelagi/pg-sweep-stats/blob/12bc23a94f15db0e55298ccb51364637ac959326/Cargo.toml)

`pg_sweep` — Collect PostgreSQL database and table statistics as JSON for monitoring.

The reviewed catalog snapshot records version `0.0.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `14,15,16`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_sweep";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_sweep';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
