## Usage

Sources:

- [Official source](https://github.com/diviyank/pg_reflex/blob/c1e787fe390ad2398bcdd438e51e4c38bdb8978c/README.md)
- [Official source](https://github.com/diviyank/pg_reflex/blob/c1e787fe390ad2398bcdd438e51e4c38bdb8978c/Cargo.toml)
- [Official source](https://github.com/diviyank/pg_reflex/blob/c1e787fe390ad2398bcdd438e51e4c38bdb8978c/pg_reflex.control)
- [Official source](https://github.com/diviyank/pg_reflex/blob/c1e787fe390ad2398bcdd438e51e4c38bdb8978c/docs/api/create_reflex_ivm.md)
- [Official source](https://github.com/diviyank/pg_reflex/blob/c1e787fe390ad2398bcdd438e51e4c38bdb8978c/docs/getting-started/upgrading.md)
- [Official source](https://github.com/diviyank/pg_reflex/blob/c1e787fe390ad2398bcdd438e51e4c38bdb8978c/sql/pg_reflex--1.12.4--1.12.5.sql)
- [Official source](https://github.com/diviyank/pg_reflex/releases/tag/1.12.5)

`pg_reflex` 1.12.5 maintains query results through source-table triggers on PostgreSQL 15–18. Installation requires a superuser; it is not trusted or relocatable. The documented setup uses `CREATE EXTENSION` without shared preload.

### Create and inspect a maintained result

```sql
CREATE EXTENSION pg_reflex;
CREATE TABLE sales (id bigint PRIMARY KEY, region text, amount numeric);
SELECT create_reflex_ivm(
    'sales_by_region',
    'SELECT region, SUM(amount) AS total FROM sales GROUP BY region',
    NULL, 'LOGGED', 'IMMEDIATE'
);
INSERT INTO sales VALUES (1, 'EU', 150);
SELECT * FROM sales_by_region;
SELECT * FROM reflex_ivm_status();
```

The result is backed by ordinary tables and generated maintenance objects. Query definitions can include supported aggregates, joins and decomposed CTEs, but not every SELECT is maintainable. The current create-time checks reject unsupported shapes; some accepted shapes recompute the full result on writes. Do not assume constant-time updates.

The example deliberately uses `LOGGED`. The default is `UNLOGGED`, whose data is not crash durable. The default maintenance mode is `IMMEDIATE`; `DEFERRED` coalesces work until COMMIT. Passthrough deletion needs a unique key, which can be inferred from a suitable source primary key. `ignore_sources` deliberately permits stale results until the documented reconcile or heal operation.

### Administration and access

Use `reflex_reconcile`, `reflex_rebuild_imv` and `drop_reflex_ivm` for their documented repair, rebuild and removal paths. Generated triggers and internal registry tables are part of the maintenance contract; do not modify them directly. Run creation with the required source-table and target-schema rights, then grant access to the result deliberately. Installation includes event triggers and a narrowly scoped security-definer partition helper; installation permission does not establish a general application security boundary.

### Upgrade boundary

After replacing the library, run `ALTER EXTENSION pg_reflex UPDATE TO '1.12.5'` promptly in each database. This migration adds helper functions and rewrites source delta triggers; inspect its warnings about bodies it could not match. Upgrading from before 1.12.0 also requires the documented SQL update and recreation of every IMV after validating its definition; copying a library alone leaves old generated maintenance code.

1.12.5 fixes operator/flush lock ordering, maintenance when both join sources change in one statement, and filled-partition removal. Reconcile previously affected join results as recommended by the upgrade guide. Source DDL, partition changes, deferred failures and crash recovery require operational checks, not just a successful SELECT.
