## Usage

Sources:

- [Enable DuckDB on an ApsaraDB RDS primary instance](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/enable-duckdb-for-the-master-instance)
- [AP acceleration engine (`rds_duckdb`)](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/how-to-use-duckdb-to-speed-up-queries/)
- [DuckDB analysis acceleration overview](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/duckdb-analytics-acceleration/)

`rds_duckdb` 1.3.2.3 is an Alibaba Cloud ApsaraDB RDS for PostgreSQL extension that maintains DuckDB column-store counterparts of PostgreSQL tables and offloads eligible analytical `SELECT` queries to DuckDB's vectorized engine. Use it for scans, joins, and aggregations on an RDS instance; it is not a portable DuckDB extension for community PostgreSQL.

### Core Workflow

Enable DuckDB acceleration for the instance in the RDS console. The managed setup configures required instance parameters, including `wal_level=logical` and `shared_preload_libraries` containing `rds_duckdb`, and causes a brief disconnection. Then connect with a privileged account:

```sql
CREATE EXTENSION IF NOT EXISTS rds_duckdb;

SELECT rds_duckdb.create_duckdb_table('public.orders');
-- Batch form:
SELECT rds_duckdb.create_duckdb_tables(
  '{public.customers,public.order_items}'
);
```

Enable offloading for a session or a single statement, and use `EXPLAIN` to prove that the plan contains `Custom Scan (DuckDBScan)`:

```sql
SET rds_duckdb.execution = on;
EXPLAIN SELECT customer_id, sum(total)
FROM orders
GROUP BY customer_id;

/*+ set(rds_duckdb.execution on) */
EXPLAIN SELECT count(*) FROM orders;
```

### Controls and Observability

- `rds_duckdb.execution` enables or disables query offloading; it is the only setting supported in a statement hint.
- `rds_duckdb.plan_cost_threshold` keeps low-cost queries on PostgreSQL, avoiding DuckDB forwarding overhead.
- `rds_duckdb.worker_threads` and `rds_duckdb.memory_limit` bound session-level DuckDB resources.
- `rds_duckdb.enable_fallback` controls whether unsupported or unavailable offloads fall back to PostgreSQL; disable it temporarily to expose the reason during troubleshooting.
- `rds_duckdb.wait_sync_timeout` bounds the wait for incremental synchronization.
- `rds_duckdb.duckdb_sync_stat` reports table synchronization state and errors and requires a privileged account.

### Requirements and Caveats

Current provider documentation requires a primary PostgreSQL 13–18 RDS instance at minor engine version 20260130 or later; read-only acceleration and newer features have narrower requirements. The cataloged extension version and the RDS engine minor version are separate identifiers.

Only eligible read-only `SELECT` statements are offloaded. DML, DDL, unsupported SQL, joins involving non-DuckDB tables, or tables that are not synchronized fall back to PostgreSQL unless fallback is disabled. A synchronized table needs a primary key or suitable replica identity. Columnar replicas consume additional storage and can lag PostgreSQL data; choose consistency and timeout settings deliberately. Availability, restart, recovery, replication, upgrade, and billing behavior remain Alibaba Cloud service concerns.
