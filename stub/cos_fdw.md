## Usage

Sources:

- [Official upstream documentation](https://cloud.tencent.com/document/product/409/75121)
- [Official project or provider page](https://cloud.tencent.com/product/postgresql)

`cos_fdw` — TencentDB PostgreSQL foreign data wrapper for reading CSV data from Tencent Cloud Object Storage (COS).

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.
Install and validate the declared extension dependencies first: `pgcrypto`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "cos_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'cos_fdw';
```

This is a provider-specific component for `Tencent Cloud`; availability, enablement, privileges, and upgrades follow that service rather than a portable community package.

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
