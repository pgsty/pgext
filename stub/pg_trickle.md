## Usage

Sources:

- [sql/pg_trickle--0.108.0--0.108.1.sql](https://github.com/trickle-labs/pg-trickle/blob/v0.108.1/sql/pg_trickle--0.108.0--0.108.1.sql)
- [Version 0.108.1 README](https://github.com/trickle-labs/pg-trickle/blob/v0.108.1/README.md)
- [SQL reference](https://github.com/trickle-labs/pg-trickle/blob/v0.108.1/docs/SQL_REFERENCE.md)
- [Configuration](https://github.com/trickle-labs/pg-trickle/blob/v0.108.1/docs/CONFIGURATION.md)
- [GUC catalog](https://github.com/trickle-labs/pg-trickle/blob/v0.108.1/docs/GUC_CATALOG.md)
- [Upgrade guide](https://github.com/trickle-labs/pg-trickle/blob/v0.108.1/docs/UPGRADING.md)
- [Control file](https://github.com/trickle-labs/pg-trickle/blob/v0.108.1/pg_trickle.control)
- [0.107.0 to 0.108.0 migration](https://github.com/trickle-labs/pg-trickle/blob/v0.108.1/sql/pg_trickle--0.107.0--0.108.0.sql)

`pg_trickle` 0.108.1 maintains stream tables on PostgreSQL 18: ordinary queryable tables derived from a SQL query, refreshed incrementally when supported or recomputed in full. Tables may depend on other stream tables, forming a dependency graph. Same-transaction maintenance is also available.

### Enable the Extension

Append `pg_trickle` to `shared_preload_libraries` and restart PostgreSQL, then install as a superuser. Size the background-worker pool for the deployment; the upstream example uses eight workers.

```ini
shared_preload_libraries = 'pg_trickle'
max_worker_processes = 8
```

```sql
CREATE EXTENSION pg_trickle;
```

The default `pg_trickle.cdc_mode` is `trigger`: transactional change capture needs neither logical WAL nor replication slots. Opt-in `auto` starts with triggers and can move eligible sources to receipt-backed WAL capture. Explicit `wal` also falls back to triggers when admission or logical-decoding prerequisites fail. These choices have different write overhead and operational requirements.

### Create and Refresh a Stream Table

```sql
CREATE TABLE orders (id bigint PRIMARY KEY, region text, amount numeric);
SELECT pgtrickle.create_stream_table(
    name => 'regional_totals',
    query => 'SELECT region, SUM(amount) AS total, COUNT(*) AS cnt FROM orders GROUP BY region',
    schedule => '30s',
    refresh_mode => 'AUTO'
);
INSERT INTO orders VALUES (1, 'east', 10);
SELECT pgtrickle.refresh_stream_table('regional_totals');
SELECT * FROM regional_totals;
```

`initialize` defaults to true, so creation populates the result. `schedule` accepts durations, cron expressions such as `@hourly`, or the default `calculated` schedule inherited from downstream dependents. `AUTO` selects differential maintenance where possible and can fall back to full refresh. `DIFFERENTIAL` rejects queries it cannot maintain incrementally; `FULL` truncates and reloads the result.

`IMMEDIATE` uses statement-level triggers inside the base-table write transaction. It does not use WAL capture and rejects an effective explicit WAL request. Use the documented query-admission rules for joins, aggregates, subqueries, recursive queries, and other supported shapes; support is not a promise that every SQL expression is differentiable.

### Lifecycle and Monitoring

`pgtrickle.alter_stream_table` changes definitions or refresh policy, and `pgtrickle.drop_stream_table` removes managed tables. `pgtrickle.repair_stream_table` repairs missing capture infrastructure and resets maintenance state after events such as restore or operator DDL. Use lifecycle APIs instead of direct writes or foreign keys against managed stream tables.

```sql
SELECT * FROM pgtrickle.pgt_status();
SELECT * FROM pgtrickle.health_check();
SELECT * FROM pgtrickle.dependency_tree();
SELECT * FROM pgtrickle.explain_st('regional_totals');
```

Lifecycle functions require explicit execution grants and ownership checks; administrator-wide operations are restricted to the extension owner or superuser. Arbitrary-SQL helpers preserve caller privileges. Capture triggers add work to source writes, and refresh failures can accumulate change buffers, so monitor health and storage instead of treating a schedule as a hard freshness guarantee.

### External Coordination and Output Deltas

`orchestration_mode` selects `MANAGED` scheduling or `EXTERNAL` coordination. External coordination cannot be combined with immediate maintenance. `pgtrickle.integration_capabilities` advertises the available contracts; version 0.108.0 exposes Graph V1 1.2 and Delta V1 1.1.

`pgtrickle.output_delta_consumer_status` reports consumers, while `pgtrickle.validate_output_delta_consumer` checks whether a consumer can resume. `pgtrickle.request_output_delta_resnapshot`, `pgtrickle.begin_output_delta_resnapshot`, and `pgtrickle.ack_output_delta_resnapshot` manage rebuilding a baseline. A resnapshot is fenced by database-instance identity, output-contract digest, and row-identity version. Follow the exact SQL reference signatures and acknowledgement protocol before advancing external delivery.

### Upgrade to 0.108.1

Install the new library and extension files before applying the packaged migration:

```sql
ALTER EXTENSION pg_trickle UPDATE TO '0.108.1';
SELECT * FROM pgtrickle.output_delta_consumer_status();
```

The upgrade preserves consumers, cursors, batches, and typed payload while adding resnapshot fences. A 0.106.1 installation can traverse the packaged 0.107.0 migration. Validate every consumer before resuming delivery; `INVALIDATED` or `RESNAPSHOT_REQUIRED` requires a new acknowledged baseline. Older releases through 0.105.2 had documented differential-result bugs for certain query shapes; upgrading does not automatically repair previously materialized rows. Use the upgrade guide’s comparison and repair procedure where applicable.

### Version 0.108.1

This patch keeps downstream `IMMEDIATE` tables current after upstream FULL refresh or truncation, recovers missing change buffers, and fixes unintended suspension after source schema changes and scalar-subquery differential refresh. PostgreSQL 18.6 support is added. After installing matching files run `ALTER EXTENSION pg_trickle UPDATE TO '0.108.1'` and verify dependent stream-table results and capture health.
