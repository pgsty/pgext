## Usage

Sources:

- [Official README](https://github.com/postgrespro/sr_plan/blob/19a6a425fc91bdfc5482650db9747bead67efd40/README.md)
- [Extension SQL for 1.2](https://github.com/postgrespro/sr_plan/blob/19a6a425fc91bdfc5482650db9747bead67efd40/init.sql)
- [Configuration implementation](https://github.com/postgrespro/sr_plan/blob/19a6a425fc91bdfc5482650db9747bead67efd40/sr_plan.c)

`sr_plan` 1.2 captures PostgreSQL execution plans and can force later matching queries to reuse an enabled saved plan. It resembles an outline or plan-baseline facility and changes planner behavior globally when preloaded; use it only with controlled queries and a rollback plan.

### Core Workflow

Preload the module, restart PostgreSQL, and create the extension in the target database.

```conf
shared_preload_libraries = 'sr_plan'
```

```sql
CREATE EXTENSION sr_plan;

SET sr_plan.write_mode = true;
SELECT * FROM orders WHERE customer_id = 42;
SET sr_plan.write_mode = false;

SELECT query_hash, plan_hash, enable, query
FROM sr_plans
ORDER BY query_hash;

UPDATE sr_plans SET enable = true WHERE query_hash = 12345;
```

Capture mode stores plans for all subsequent queries, including duplicates. Inspect the recorded row before enabling it; do not enable every captured plan indiscriminately.

### Main Objects and Settings

- `sr_plans` stores query hashes, query identifiers, plan hashes, plan text, enable state, and referenced relation/index OIDs.
- `sr_plan.write_mode` enables plan capture and is a superuser setting.
- `sr_plan.enabled` enables the planner hook and is a superuser setting.
- `sr_plan.log_usage` controls the message level for cached-plan use.
- `_p(anyelement)` marks constants that should be treated as parameters when matching plans.
- `show_plan(query_hash, index, format)` displays a saved plan in `text`, `json`, `xml`, or `yaml` form.

An event trigger deletes saved plans when referenced tables or indexes are dropped. Schema changes, statistics changes, and data-distribution changes can still make a once-good forced plan harmful, so monitor and prune `sr_plans`. If `pg_stat_statements` is also preloaded, upstream requires it to appear after `sr_plan` in the preload list. The project is unmaintained; verify compatibility with the exact PostgreSQL major version before use.
