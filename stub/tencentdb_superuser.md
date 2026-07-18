## Usage

Sources:

- [TencentDB PostgreSQL privilege documentation](https://cloud.tencent.com/document/product/409/61551)
- [TencentDB PostgreSQL extension support matrix](https://cloud.tencent.com/document/product/409/75121)

`tencentdb_superuser` is a TencentDB for PostgreSQL provider extension associated with the managed `pg_tencentdb_superuser` role. That role substitutes for unavailable native superuser access and grants broad administration capabilities, including database/role creation, row-security bypass, replication, object access, and controlled extension management.

```sql
CREATE EXTENSION tencentdb_superuser;
SELECT rolname, rolcreatedb, rolcreaterole, rolreplication, rolbypassrls
FROM pg_roles
WHERE rolname = 'pg_tencentdb_superuser';
```

This is a Tencent Cloud service component, not a portable open-source extension. Capabilities and version availability depend on the TencentDB engine release and service policy. Grant the managed role only to trusted administrators, audit membership changes, avoid application ownership, and verify every required operation against the current instance documentation.
