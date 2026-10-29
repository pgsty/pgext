## Usage

Sources:

- [README.md](https://github.com/kiransabne04/pg_stat_index_queries/blob/6a079ffc709515574271e55f4c71b528f60e9f2d/README.md)
- [INSTALL.md](https://github.com/kiransabne04/pg_stat_index_queries/blob/6a079ffc709515574271e55f4c71b528f60e9f2d/docs/INSTALL.md)
- [pg_index_deps.control](https://github.com/kiransabne04/pg_stat_index_queries/blob/6a079ffc709515574271e55f4c71b528f60e9f2d/sql_edition/pg_index_deps.control)
- [pg_index_deps--0.7.sql](https://github.com/kiransabne04/pg_stat_index_queries/blob/6a079ffc709515574271e55f4c71b528f60e9f2d/sql_edition/pg_index_deps--0.7.sql)
- [Makefile](https://github.com/kiransabne04/pg_stat_index_queries/blob/6a079ffc709515574271e55f4c71b528f60e9f2d/sql_edition/Makefile)
- [ci.yml](https://github.com/kiransabne04/pg_stat_index_queries/blob/6a079ffc709515574271e55f4c71b528f60e9f2d/.github/workflows/ci.yml)

`pg_index_deps` is the pure-SQL edition of `pg_stat_index_queries`, with SQL version 0.7 and fixed schema `index_deps`. It derives index dependencies from captured plan text rather than observing native execution hooks. Upstream tests PostgreSQL 16–18; its Aurora instructions additionally describe 14.10/15.5 or newer, with the drop-impact function requiring 16 or newer.

### Core Workflow

On a self-hosted server with the extension files and plan-capture dependencies installed:

```sql
CREATE EXTENSION pg_stat_statements;
CREATE EXTENSION pg_stat_plans;
CREATE EXTENSION hypopg;
CREATE EXTENSION pg_index_deps;
SELECT indexname, index_role, queries, verdict
FROM index_deps.index_usage;
SELECT queryid, calls, scan_types, index_role
FROM index_deps.index_dependents WHERE indexname = 'orders_customer_idx';
SELECT * FROM index_deps.index_drop_impact('public.orders_customer_idx');
```

The native `pg_stat_index_queries` extension is optional. On managed servers, use the official plain-SQL installer or its pg_tle installer instead of assuming arbitrary server-side control files can be installed. The pg_tle route has its own preload requirement. Aurora plan capture requires `aurora_compute_plan_id`; follow the provider's parameter-group activation procedure.

### Data Sources and Objects

`index_deps.config` selects auto, aurora_stat_plans, pg_stat_plans or none. `index_deps.plan_index_usage` parses index references; `index_deps.index_dependents` aggregates them by query and index. `index_deps.index_usage`, `index_deps.index_role`, `index_deps.index_table_stats`, `index_deps.index_column_stats` and `index_deps.index_hints` add purpose and statistics context. `index_deps.unresolved_plan_indexes` identifies captured names that no longer resolve.

A missing plan source, evicted query, renamed/dropped index or uncaptured workload limits the result. In particular, an empty dependency set with no plan source does not prove that an index is unused. Query timing remains per statement, not per index.

### Drop-Impact Boundary

`index_deps.index_drop_impact` uses HypoPG to hide the index in the current session and compares `EXPLAIN (GENERIC_PLAN)` costs. It does not execute the statements or drop the real index. Cost ratios are planner estimates, not latency measurements; stale statistics and partial indexes with parameterized predicates can make the answer unknown. Review blockers, real parameter values and representative workload before acting.

### Upgrade

After installing the matching 0.7 files, an extension-managed installation can use:

```sql
ALTER EXTENSION pg_index_deps UPDATE TO '0.7';
```

For pg_tle, first register the newer installer and update paths, then update the extension. A plain-SQL installation is maintained through its installer and is not automatically tracked by `pg_extension`.
