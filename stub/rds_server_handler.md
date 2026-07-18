## Usage

Sources:

- [Official upstream documentation](https://cloud.tencent.com/document/product/409/67017)
- [Official project or provider page](https://cloud.tencent.com/product/postgres)
- [Official primary documentation](https://cloud.tencent.com/document/product/409/75121)

`rds_server_handler` — TencentDB PostgreSQL server maintenance handler extension

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "rds_server_handler";
SELECT extversion
FROM pg_extension
WHERE extname = 'rds_server_handler';
```

This is a provider-specific component for `Tencent Cloud`; availability, enablement, privileges, and upgrades follow that service rather than a portable community package.

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
