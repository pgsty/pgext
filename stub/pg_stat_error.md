## Usage

Sources:

- [TencentDB for PostgreSQL supported-extension matrix](https://cloud.tencent.com/document/product/409/75121)

`pg_stat_error` is a provider-managed TencentDB for PostgreSQL extension. The current official matrix lists extension version 1.0 for PostgreSQL 10 through 13 and does not list it for PostgreSQL 14 through 18.

```sql
CREATE EXTENSION pg_stat_error;
```

Install it only on a supported TencentDB instance and verify availability against the instance's engine version before creating it.

### Provider Boundary

Tencent Cloud does not publish a source repository, control file, SQL object reference, license, or portable installation package for this extension. The public support page does not document callable functions or views, so no error-statistics API should be assumed from the name alone. Use the objects exposed by the installed version and Tencent Cloud support guidance; it is not a self-hosted PostgreSQL component.
