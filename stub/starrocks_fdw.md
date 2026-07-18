## Usage

Sources:

- [Official upstream documentation](https://cloud.tencent.com/document/product/409/111197)

`starrocks_fdw` — TencentDB for PostgreSQL foreign data wrapper for querying StarRocks tables through a MySQL-compatible interface.

The reviewed catalog snapshot records version `1.2`, kind `standard`, and implementation language `C`.
The curated compatibility set is `13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "starrocks_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'starrocks_fdw';
```

This is a provider-specific component for `Tencent Cloud`; availability, enablement, privileges, and upgrades follow that service rather than a portable community package.

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
