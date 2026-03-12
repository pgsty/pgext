

## Usage

> [pg_store_plans: execution plan storage and statistics](https://github.com/ossc-db/pg_store_plans)

pg_store_plans tracks execution plan statistics for all SQL statements, complementing `pg_stat_statements` with plan-level detail. Joinable via `queryid` on PostgreSQL 14+.

### Viewing Plan Statistics

```sql
-- View tracked plans with statistics
SELECT queryid, planid, plan, calls, total_time, rows
FROM pg_store_plans
ORDER BY total_time DESC;

-- Check module status
SELECT * FROM pg_store_plans_info;
```

### Key View Columns

| Column | Type | Description |
|--------|------|-------------|
| `queryid` | bigint | Query ID (joinable with pg_stat_statements) |
| `planid` | bigint | Plan hash code |
| `plan` | text | Representative plan text |
| `calls` | bigint | Execution count |
| `total_time` | double precision | Total execution time (ms) |
| `rows` | bigint | Total rows retrieved/affected |
| `shared_blks_hit` | bigint | Shared buffer hits |
| `shared_blks_read` | bigint | Shared blocks read |
| `first_call` | timestamptz | First execution time |
| `last_call` | timestamptz | Last execution time |

### Functions

```sql
-- Reset all statistics (superuser only)
SELECT pg_store_plans_reset();

-- Convert plan formats
SELECT pg_store_plans_textplan(plan);   -- to text
SELECT pg_store_plans_jsonplan(plan);   -- to JSON
SELECT pg_store_plans_xmlplan(plan);    -- to XML
SELECT pg_store_plans_yamlplan(plan);   -- to YAML

-- Calculate query hash
SELECT pg_store_hash_query('SELECT 1');
```

### Configuration

| Parameter | Default | Description |
|-----------|---------|-------------|
| `pg_store_plans.max` | 1000 | Maximum tracked plans (server start only) |
| `pg_store_plans.track` | `top` | `top`, `all`, `verbose`, `none` |
| `pg_store_plans.plan_format` | `text` | `text`, `json`, `xml`, `yaml`, `raw` |
| `pg_store_plans.min_duration` | 0 | Minimum execution time to track (ms) |
| `pg_store_plans.log_analyze` | `off` | Include EXPLAIN ANALYZE output |
| `pg_store_plans.log_buffers` | `off` | Include buffer statistics |
| `pg_store_plans.log_timing` | `true` | Record actual timings |
| `pg_store_plans.plan_storage` | `file` | Storage: `file` or `shmem` |
| `pg_store_plans.max_plan_length` | 5000 | Max bytes for plan text |
| `pg_store_plans.save` | `on` | Persist stats across restarts |
