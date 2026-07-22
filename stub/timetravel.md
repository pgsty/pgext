## Usage

Sources:

- [TencentDB for PostgreSQL supported-extension matrix](https://cloud.tencent.com/document/product/409/75121)
- [TencentDB for PostgreSQL product page](https://cloud.tencent.com/product/postgres)
- [Historical PostgreSQL 11 timetravel documentation](https://www.postgresql.org/docs/11/contrib-spi.html)
- [PostgreSQL 12 release note removing timetravel](https://www.postgresql.org/docs/12/release-12.html)

`timetravel` is available as a provider-managed extension on TencentDB for PostgreSQL. Tencent Cloud's current matrix lists version 1.0 for PostgreSQL 10 and 11 only, with no support shown for PostgreSQL 12 or later. Confirm the instance engine and current matrix before relying on it.

### Provider Workflow

TencentDB documents self-service extension creation for entries in its supported matrix. On an eligible test database, create and immediately inventory the installed objects before designing application logic:

```sql
CREATE EXTENSION timetravel;

SELECT extname, extversion
FROM pg_extension
WHERE extname = 'timetravel';

SELECT n.nspname, p.proname, pg_get_function_identity_arguments(p.oid)
FROM pg_proc AS p
JOIN pg_namespace AS n ON n.oid = p.pronamespace
JOIN pg_depend AS d ON d.objid = p.oid
JOIN pg_extension AS e ON e.oid = d.refobjid
WHERE e.extname = 'timetravel'
ORDER BY 1, 2, 3;
```

Use the TencentDB account authorized to manage extensions and test on the same kernel minor version as production. Provider availability, upgrade behavior, backups, replicas, and failover remain TencentDB service boundaries.

### API and Compatibility Boundary

PostgreSQL historically shipped a version 1.0 module with the same name that implemented row-history behavior through triggers and session-local control functions. PostgreSQL removed that module in version 12 because it depended on obsolete types and old code. These historical documents explain the likely lineage, but Tencent Cloud's matrix does not publish source, control files, object definitions, privileges, or a statement that its build is identical.

Do not assume historical functions, trigger arguments, data types, or semantics solely from the shared name. Compare the installed object inventory with provider documentation or open a Tencent Cloud support case before creating tables or triggers. Plan migration away from `timetravel` before a PostgreSQL major upgrade because the current matrix does not list it beyond version 11.
