
## Usage

> Syntax:
>
> ```sql
> SELECT * FROM pg_store_plans ORDER BY total_time DESC;
> SELECT * FROM pg_store_plans_info;
> ```
>
> Sources: [Project page](http://ossc-db.github.io/pg_store_plans/), [Bundled docs](https://github.com/ossc-db/pg_store_plans/blob/master/docs/index.html)

`pg_store_plans` tracks execution plan statistics for SQL statements, similar in spirit to how `pg_stat_statements` tracks statements. It records plan text, plan hash, timing, row counts, and buffer statistics, and its docs note that `queryid` can be used to join with `pg_stat_statements`.

## Configuration

The upstream documentation requires:

```ini
shared_preload_libraries = 'pg_store_plans'
compute_query_id = 'on'
```

`pg_store_plans` needs shared memory, so adding or removing it requires a server restart. If `compute_query_id` is set to `no`, the module is silently disabled.

## Viewing Plan Statistics

The statistics are exposed through the `pg_store_plans` view:

```sql
SELECT queryid, planid, plan, calls, total_time, rows
FROM pg_store_plans
ORDER BY total_time DESC;

SELECT * FROM pg_store_plans_info;
```

Important columns documented upstream include:

- `queryid`, the core-generated query ID
- `planid`, a normalized plan hash
- `plan`, in the format chosen by `pg_store_plans.plan_format`
- `calls`, `total_time`, and `rows`
- buffer statistics such as `shared_blks_hit` and `shared_blks_read`
- timestamps such as `first_call` and `last_call`

## Helper Functions

```sql
SELECT pg_store_plans_reset();
SELECT pg_store_plans_textplan(plan);
SELECT pg_store_plans_jsonplan(plan);
SELECT pg_store_plans_xmlplan(plan);
SELECT pg_store_plans_yamlplan(plan);
SELECT pg_store_hash_query('SELECT 1');
```

These functions reset statistics, convert stored plans to different output formats, and compute query hashes.

## GUCs

The extension documentation describes settings such as:

- `pg_store_plans.max`
- `pg_store_plans.track`
- `pg_store_plans.plan_format`
- `pg_store_plans.min_duration`
- `pg_store_plans.log_analyze`
- `pg_store_plans.log_buffers`
- `pg_store_plans.log_timing`
- `pg_store_plans.plan_storage`
- `pg_store_plans.max_plan_length`
- `pg_store_plans.save`

Together these control collection scope, plan representation, persistence, and storage behavior.
