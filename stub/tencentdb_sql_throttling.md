## Usage

Sources:

- [Official project or provider page](https://cloud.tencent.com/product/postgres)

`tencentdb_sql_throttling` — TencentDB proprietary preloaded extension for limiting SQL concurrency by normalized SQL text or query ID across primary and read-only nodes.

The reviewed catalog snapshot records version `1.0`, kind `preload`, and implementation language `C`.
The curated compatibility set is `14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "tencentdb_sql_throttling";
SELECT extversion
FROM pg_extension
WHERE extname = 'tencentdb_sql_throttling';
```

This is a provider-specific component for `Tencent Cloud`; availability, enablement, privileges, and upgrades follow that service rather than a portable community package.

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
