## Usage

Sources:

- [Alibaba Cloud RDS for PostgreSQL extension matrix](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/extensions-supported-by-apsaradb-rds-for-postgresql)

`encdb` is an Alibaba Cloud RDS for PostgreSQL provider extension described as supplying confidential-database capabilities. It is not a portable community package: availability, enablement, upgrade timing, privileges, and supported operations are controlled by the managed service.

### Check Availability First

The current Standard Edition matrix lists `encdb` `1.1.14` for PostgreSQL 11 through 17, `1.1.13` for PostgreSQL 10, and no version for PostgreSQL 18. Other editions, regions, engine minor versions, or service updates can differ, so query the target instance rather than relying on a static matrix.

```sql
SELECT name, default_version, installed_version
FROM pg_available_extensions
WHERE name = 'encdb';
```

Only after the provider reports the extension as available and the service's enablement requirements have been satisfied should an authorized administrator install it:

```sql
CREATE EXTENSION encdb;

SELECT extversion
FROM pg_extension
WHERE extname = 'encdb';
```

### Managed-Service Boundary

The public extension matrix does not document a portable SQL API, preload setting, client requirement, key-management procedure, or privilege model for `encdb`. Obtain the instance-specific confidential-database documentation and support confirmation before use; do not infer APIs from similarly named projects. Test backup/restore, replica and failover behavior, client compatibility, encryption-key lifecycle, version upgrades, and what metadata or query results remain visible to service administrators. Follow Alibaba Cloud's restrictions for extension creation and removal.
