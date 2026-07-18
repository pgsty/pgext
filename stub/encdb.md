## Usage

Sources:

- [Official upstream documentation](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/extensions-supported-by-apsaradb-rds-for-postgresql)

`encdb` — Alibaba Cloud RDS extension that provides Always Confidential Database features for PostgreSQL.

The reviewed catalog snapshot records version `1.1.14`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "encdb";
SELECT extversion
FROM pg_extension
WHERE extname = 'encdb';
```

This is a provider-specific component for `Alibaba Cloud`; availability, enablement, privileges, and upgrades follow that service rather than a portable community package.

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
