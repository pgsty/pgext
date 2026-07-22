## Usage

Sources:

- [Alibaba Cloud usage documentation](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/use-the-pg-concurrency-control-plug-in)
- [Alibaba Cloud extension support matrix](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/extensions-supported-by-apsaradb-rds-for-postgresql)

`pg_concurrency_control` 1.0 is an Alibaba Cloud ApsaraDB RDS for PostgreSQL extension that queues SQL by workload class when a configured concurrency limit is reached. It is a provider-specific protection against resource saturation, not a portable community extension. The current official usage page limits this workflow to RDS PostgreSQL 10 and 11.

### Enable and Observe

Create the extension on a supported instance and inspect its queue counters:

```sql
CREATE EXTENSION pg_concurrency_control;

SELECT * FROM pg_concurrency_control_status();
```

The status function returns `autocommit_count`, `bigquery_count`, `query_count`, and `transaction_count`. A positive counter means statements are waiting in that class.

Configure concurrency parameters through the RDS parameter-management workflow. Each limit defaults to `0`, which disables that class; values from `1` through `1024` enable queuing.

### Limits and Timeouts

- `pg_concurrency_control.query_concurrency` limits SELECT statements.
- `pg_concurrency_control.bigquery_concurrency` limits statements explicitly marked as slow queries.
- `pg_concurrency_control.transaction_concurrency` limits transaction blocks.
- `pg_concurrency_control.autocommit_concurrency` limits autocommit DML.
- `pg_concurrency_control.control_timeout` and `pg_concurrency_control.bigsql_control_timeout` set queue timeouts.
- `pg_concurrency_control.timeout_action` and `pg_concurrency_control.bigsql_timeout_action` select `TCC_break`, `TCC_rollback`, or `TCC_wait` behavior.

Mark a statement for the big-query class with the provider hint:

```sql
/*+bigsql*/ SELECT pg_sleep(10);
```

### Operational Boundary

Concurrency limits change latency and failure behavior for every matching statement. Load-test queue depth, application timeouts, transaction retry logic, pool size, and each timeout action before production rollout. A queue can protect CPU or I/O but can also move overload into connection pools and client timeout storms. Monitor the four status counters together with active sessions, latency, errors, and instance resources.

Availability and editable parameter behavior depend on the exact RDS engine minor release. The support matrix currently lists 1.0 only for the older supported majors, so verify the target instance before a major-version upgrade. Configure and remove the extension through Alibaba Cloud's privilege and maintenance model; do not copy its binaries to self-managed PostgreSQL.
