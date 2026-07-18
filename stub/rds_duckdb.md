## Usage

Sources:

- [Official upstream documentation](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/how-to-use-duckdb-to-speed-up-queries/)
- [Official project or provider page](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/duckdb-analytics-acceleration/)

`rds_duckdb` — Alibaba Cloud RDS PostgreSQL DuckDB analytical acceleration extension

The reviewed catalog snapshot records version `1.3.2.3`, kind `preload`, and implementation language `C`.
The curated compatibility set is `13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "rds_duckdb";
SELECT extversion
FROM pg_extension
WHERE extname = 'rds_duckdb';
```

This is a provider-specific component for `Alibaba Cloud`; availability, enablement, privileges, and upgrades follow that service rather than a portable community package.

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
