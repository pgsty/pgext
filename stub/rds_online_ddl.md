## Usage

Sources:

- [Official Alibaba Cloud RDS online DDL documentation](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/use-the-rds-online-ddl-extension-to-modify-column-data-type-online)
- [Official ApsaraDB RDS for PostgreSQL overview](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/what-is-apsaradb-rds-for-postgresql)

`rds_online_ddl` is an Alibaba Cloud ApsaraDB RDS for PostgreSQL service extension for changing a column's data type with a shorter final blocking window than a conventional table rewrite. It copies rows to a temporary table, builds indexes, applies concurrent changes through logical decoding, and then briefly takes an exclusive lock to swap the tables.

### Preconditions

The provider documentation requires an RDS PostgreSQL instance on PostgreSQL 12 or later with a qualifying minor engine version, `wal_level = logical`, a primary key or unique constraint on the target table, and a sufficiently privileged account. Verify the exact engine-version threshold in the current service console and documentation before scheduling the operation.

### Core Workflow

Create the provider extension using the supported RDS enablement path, estimate space and runtime, and submit only the `ALTER COLUMN ... TYPE ...` clause to `rds_online_ddl.alter_table`:

```sql
CREATE EXTENSION rds_online_ddl;

SELECT rds_online_ddl.alter_table(
  'public.orders',
  'ALTER COLUMN amount TYPE numeric(20,2)'
);
```

Monitor active work through the provider progress view:

```sql
SELECT *
FROM rds_online_ddl.pg_stat_progress_online_ddl;
```

After completion, verify column type, constraints, indexes, row counts, application reads/writes, and query plans rather than treating the function return as end-to-end validation.

### Limits and Operational Risk

- Alibaba documents support for changing column data types, not arbitrary online DDL. Do not pass unrelated `ALTER TABLE` actions.
- Budget roughly twice the table-plus-index footprint during the copy. Monitor instance storage, I/O, WAL generation, replication lag, long transactions, and logical-decoding retention.
- The provider excludes tables with foreign keys and partitioned tables. Check both inbound and outbound constraints and the actual target relation before starting.
- The final swap still requires a brief exclusive lock. Set a maintenance window, control lock wait behavior, and remove long-running transactions that could delay the cutover.
- Expressions using `USING` are subject to provider limitations and may not support every conversion. Rehearse the exact clause with representative data, including nulls, defaults, generated values, and conversion failures.
- The operation is provider-managed and rolls back atomically on failure according to Alibaba's documented workflow, but cancellation, failover, backups, replicas, and retry behavior should still be tested for the current RDS engine version.
