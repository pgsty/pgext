## Usage

Sources:

- [TencentDB for PostgreSQL supported-extension matrix](https://cloud.tencent.com/document/product/409/75121)

`timetravel` is a provider-managed TencentDB for PostgreSQL extension. The current official matrix lists extension version 1.0 only for PostgreSQL 10 and 11.

```sql
CREATE EXTENSION timetravel;
```

Create it only on a TencentDB instance whose engine version lists the extension, and check the current matrix before provisioning or upgrading an instance.

### Provider Boundary

Tencent Cloud publishes no public source, control file, SQL object reference, license, dependency list, or portable package for this component. In particular, the support matrix does not establish that it implements the historical PostgreSQL `timetravel` contrib module or any similarly named third-party API. Inspect the installed objects and use provider support documentation instead of assuming functions from the name.
