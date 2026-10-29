## Usage

Sources:

- [README.md](https://github.com/kiransabne04/pg_stat_index_queries/blob/6a079ffc709515574271e55f4c71b528f60e9f2d/README.md)
- [INSTALL.md](https://github.com/kiransabne04/pg_stat_index_queries/blob/6a079ffc709515574271e55f4c71b528f60e9f2d/docs/INSTALL.md)
- [GUIDE.md](https://github.com/kiransabne04/pg_stat_index_queries/blob/6a079ffc709515574271e55f4c71b528f60e9f2d/docs/GUIDE.md)
- [pg_stat_index_queries.control](https://github.com/kiransabne04/pg_stat_index_queries/blob/6a079ffc709515574271e55f4c71b528f60e9f2d/pg_stat_index_queries.control)
- [pg_stat_index_queries--0.10.sql](https://github.com/kiransabne04/pg_stat_index_queries/blob/6a079ffc709515574271e55f4c71b528f60e9f2d/pg_stat_index_queries--0.10.sql)
- [ci.yml](https://github.com/kiransabne04/pg_stat_index_queries/blob/6a079ffc709515574271e55f4c71b528f60e9f2d/.github/workflows/ci.yml)

`pg_stat_index_queries` records which queries plan and execute index scans, helping assess index use and drop risk. This source snapshot uses SQL version 0.10; upstream tests the native C extension on PostgreSQL 15–18. Observed workload is evidence of use, not proof that an unobserved index is unnecessary.

### Core Workflow

Add the library to the existing preload list and restart PostgreSQL:

```ini
shared_preload_libraries = 'pg_stat_index_queries'
compute_query_id = on
```

```sql
CREATE EXTENSION pg_stat_index_queries;
SELECT indexname, index_role, queries, verdict
FROM pg_index_usage ORDER BY verdict;
SELECT queryid, calls_executed, last_seen
FROM pg_index_dependents WHERE indexname = 'orders_customer_idx';
SELECT * FROM pg_stat_index_queries_info;
```

Collect representative workload on every replica as well as the primary. Native hooks count planned and executed use separately; bitmap scans count as executed because PostgreSQL constructs their scan state at startup.

### Important Objects

- `pg_index_usage` combines observed use, index role, blockers and hints; `pg_index_role` explains constraint, foreign-key, replica-identity and other roles.
- `pg_index_column_usage` and `pg_index_sole_path` help identify column use and queries whose observed path relied on one index.
- `pg_index_query_stats` enriches index/query pairs with `pg_stat_statements` when available. Statement timing and buffers are repeated for each index the query used; they are not per-index execution costs.
- `pg_index_table_stats`, `pg_index_column_stats` and `pg_index_hints` use catalog/statistics evidence, with its normal sampling and visibility limits.
- `pg_stat_index_queries_reset()` clears counters and is revoked from PUBLIC by default.

### Configuration and Retention

`pg_stat_index_queries.max` defaults to 10000 tracked pairs and needs restart to change. `pg_stat_index_queries.track` accepts top, all or none. `pg_stat_index_queries.save` defaults on; `pg_stat_index_queries.save_interval` defaults to 300 seconds and `pg_stat_index_queries.flush_interval` to 1 second. A crash can lose counts since the last save; other sessions see local counts after flushing.

Sampling through `pg_stat_index_queries.sample_rate` makes counts estimates and can miss rare users. Evictions and earlier counter resets also limit conclusions. Query-statistics resets are independent, so displayed percentages can exceed 100. Review constraints, month-end jobs, replicas and alternative plans before any index removal; the extension does not make a safe automatic DROP decision.

### Managed PostgreSQL Alternative

Custom C libraries are not generally installable on RDS/Aurora. The separate `pg_index_deps` SQL edition uses captured plans instead and has its own installation and data-coverage boundaries; it does not supply the native executed-plan counters.
