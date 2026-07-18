## Usage

Sources:

- [Official upstream documentation](https://support.huaweicloud.com/intl/en-us/usermanual-rds-pg/rds_09_0068.html)

`rds_hwdrs_privs` — Huawei Cloud RDS helper extension for granting selected replication and catalog privileges.

The reviewed catalog snapshot records version `1`, kind `standard`, and implementation language `SQL`.
The curated compatibility set is `10,11,12,13,14,15,16,17`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "rds_hwdrs_privs";
SELECT extversion
FROM pg_extension
WHERE extname = 'rds_hwdrs_privs';
```

This is a provider-specific component for `Huawei Cloud`; availability, enablement, privileges, and upgrades follow that service rather than a portable community package.

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
