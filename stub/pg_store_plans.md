## Usage

Sources: [official docs](https://ossc-db.github.io/pg_store_plans/), [repo](https://github.com/ossc-db/pg_store_plans), [1.10 release notes](https://github.com/ossc-db/pg_store_plans/releases/tag/1.10)

`pg_store_plans` tracks execution plan statistics for SQL statements, similar to how `pg_stat_statements` tracks statement statistics. The upstream `1.10` release note says this version adds PostgreSQL 18 support; the documented SQL surface is otherwise the same as current docs.

### Required server settings

```ini
shared_preload_libraries = 'pg_store_plans'
compute_query_id = 'on'
```

`pg_store_plans` requires shared memory, so adding or removing it needs a server restart. The docs say it is silently disabled if `compute_query_id` is `no`.

### Inspect stored plans

```sql
SELECT queryid, planid, plan, calls, total_time, rows
FROM pg_store_plans
ORDER BY total_time DESC;

SELECT * FROM pg_store_plans_info;
```

The docs describe `queryid` as the join key for `pg_stat_statements`, and `pg_store_plans_info` as a one-row view that exposes module-level stats such as `dealloc` and `stats_reset`.

### Helper functions

```sql
SELECT pg_store_plans_reset();
SELECT pg_store_hash_query('SELECT 1');
SELECT pg_store_plans_textplan(plan);
SELECT pg_store_plans_jsonplan(plan);
SELECT pg_store_plans_xmlplan(plan);
SELECT pg_store_plans_yamlplan(plan);
```

`pg_store_plans_*plan()` is useful when `pg_store_plans.plan_format = 'raw'`.

### Key GUCs

- `pg_store_plans.max`
- `pg_store_plans.track`
- `pg_store_plans.max_plan_length`
- `pg_store_plans.plan_storage`
- `pg_store_plans.plan_format`
- `pg_store_plans.min_duration`
- `pg_store_plans.log_analyze`
- `pg_store_plans.log_buffers`
- `pg_store_plans.log_timing`
- `pg_store_plans.save`

The docs describe `plan_storage` as `file` or `shmem`, and `plan_format` as `text`, `json`, `xml`, `yaml`, or `raw`.

### Caveats

- Non-superusers cannot see `plan`, `queryid`, or `planid` for statements executed by other users.
- `pg_store_plans` and `pg_stat_statements` maintain entries independently, so low-frequency rows may not always have a matching peer.
