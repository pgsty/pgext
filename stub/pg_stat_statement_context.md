## Usage

Sources:

- [README.md](https://github.com/benchub/pg_stat_statement_context/blob/63c012c63e9a8ec1ebbb5271ba8c44f0b53e0fc1/README.md)
- [sql/pg_stat_statement_context--1.0.sql](https://github.com/benchub/pg_stat_statement_context/blob/63c012c63e9a8ec1ebbb5271ba8c44f0b53e0fc1/sql/pg_stat_statement_context--1.0.sql)
- [docs/sql-interface.md](https://github.com/benchub/pg_stat_statement_context/blob/63c012c63e9a8ec1ebbb5271ba8c44f0b53e0fc1/docs/sql-interface.md)
- [docs/limitations.md](https://github.com/benchub/pg_stat_statement_context/blob/63c012c63e9a8ec1ebbb5271ba8c44f0b53e0fc1/docs/limitations.md)
- [docs/configuration.md](https://github.com/benchub/pg_stat_statement_context/blob/63c012c63e9a8ec1ebbb5271ba8c44f0b53e0fc1/docs/configuration.md)
- [pg_stat_statement_context.control](https://github.com/benchub/pg_stat_statement_context/blob/63c012c63e9a8ec1ebbb5271ba8c44f0b53e0fc1/pg_stat_statement_context.control)

`pg_stat_statement_context` records calls and execution time by query identity, user, database and application tags in rolling time buckets. It supports PostgreSQL 14–18 and complements `pg_stat_statements`.

### Core Workflow

```ini
shared_preload_libraries = 'pg_stat_statements, pg_stat_statement_context'
```

```sql
CREATE EXTENSION pg_stat_statement_context;
SELECT 1 /*controller:users,action:show*/;
SELECT queryid, tags, calls, total_exec_time
FROM pg_stat_statement_context_totals ORDER BY calls DESC;
```

### Operational Boundaries

Merge the configuration below with existing preload libraries and restart. If `pg_stat_statements` is enabled, it must precede this library so utility query IDs are captured. Collection is cluster-wide after preload; SQL objects are installed per database. `pg_stat_statement_context_totals` summarizes buckets; `pg_stat_statement_context_activity` exposes current tags, while `pg_stat_statement_context_info` reports capacity and skipped-work counters. Reset and extraction helpers are superuser-only by default; other users' tags require the privileges of `pg_read_all_stats`. Keep tags low-cardinality and exclude personal data. Prepared statements retain comments from preparation; use `pg_stat_statement_context.tags_override` for execution-time request context. Statistics are bounded and evictable, and failed statements or statements without a query ID are not a complete audit trail.
