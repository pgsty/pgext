## Usage

Sources:

- [Official TencentDB for PostgreSQL extension matrix](https://cloud.tencent.com/document/product/409/75121)

`tencentdb_stat` is a TencentDB for PostgreSQL provider extension. The official support matrix lists version 1.0 for PostgreSQL 10 through 14 and does not list it for PostgreSQL 15 through 18. It is not a portable community package, and the public provider page does not document its SQL objects or function signatures.

### Enable and Verify

On a TencentDB instance and engine version where the extension appears as supported, create it using the service's extension-management privileges and verify the installed version:

```sql
CREATE EXTENSION tencentdb_stat;

SELECT extversion
FROM pg_extension
WHERE extname = 'tencentdb_stat';
```

If creation is rejected or the extension is absent, verify the exact TencentDB PostgreSQL major and kernel minor release in the console and consult the current provider support channel. Do not copy binaries from another service or self-managed PostgreSQL host into TencentDB.

### Provider Boundary

Availability, privileges, upgrades, and removal are controlled by Tencent Cloud. Recheck the official matrix before an engine-major upgrade: a database using `tencentdb_stat` on PostgreSQL 10 through 14 may not have the same extension available on a PostgreSQL 15 or later target.

Because the public official page only confirms name, version, and supported engine majors, do not invent QPS views, functions, GUCs, or permissions from catalog descriptions. After creation, inspect the objects installed on a nonproduction instance and obtain the applicable Tencent Cloud documentation or support guidance before granting access or integrating monitoring. Test upgrade, failover, read-only nodes, performance overhead, and extension removal under the provider's actual service behavior.
