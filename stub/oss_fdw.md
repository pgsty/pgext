## Usage

Sources:

- [Official upstream documentation](https://www.alibabacloud.com/help/en/polardb/polardb-for-postgresql/read-and-write-external-data-files-by-using-oss-fdw)

`oss_fdw` — Alibaba Cloud OSS foreign data wrapper for reading and writing OSS objects from PostgreSQL-compatible managed services.

The reviewed catalog snapshot records version `unknown`, kind `standard`, and implementation language `C`.
The curated compatibility set is `11,14,16`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "oss_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'oss_fdw';
```

This is a provider-specific component for `Alibaba Cloud`; availability, enablement, privileges, and upgrades follow that service rather than a portable community package.

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
