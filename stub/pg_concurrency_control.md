## Usage

Sources:

- [Official Alibaba Cloud pg_concurrency_control guide](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/use-the-pg-concurrency-control-plug-in)
- [Official ApsaraDB RDS PostgreSQL supported-extension list](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/extensions-supported-by-apsaradb-rds-for-postgresql)

`pg_concurrency_control` is an Alibaba Cloud ApsaraDB RDS PostgreSQL extension that caps concurrently executing SQL by category and queues excess work. It has separate limits for SELECT queries, autocommit DML, transaction blocks, and explicitly marked slow queries through `query_concurrency`, `autocommit_concurrency`, `transaction_concurrency`, and `bigquery_concurrency`.

### Enable a Query Limit

```sql
CREATE EXTENSION pg_concurrency_control;

SET pg_concurrency_control.query_concurrency = 10;

SELECT * FROM pg_concurrency_control_status();
```

Each concurrency limit defaults to zero, which disables that category's control. `pg_concurrency_control_status()` reports the current counts in the autocommit, big-query, query, and transaction queues.

### Timeouts and Scope

Queue wait limits are controlled by `control_timeout` and `bigsql_control_timeout`. The documented timeout actions are `TCC_break`, `TCC_rollback`, and `TCC_wait`; choose them according to whether a timed-out statement should be skipped, abort its transaction, or continue waiting.

### Caveats

- This is a provider-only extension for ApsaraDB RDS PostgreSQL, not a generally published extension for self-managed PostgreSQL.
- The dedicated usage guide currently requires an ApsaraDB RDS instance running PostgreSQL 10 or 11. Check the supported-extension table and upgrade to the latest minor engine version before enabling it.
- Version 1.0 is listed for supported engine combinations. The provider does not publish its source, control file, software license, or portable installation artifacts.
- The official guide installs it directly with `CREATE EXTENSION`; it does not document a preload step for this extension.
