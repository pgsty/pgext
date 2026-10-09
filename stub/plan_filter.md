## Usage

Sources:

- [README.md](https://github.com/pgexperts/pg_plan_filter/blob/c1afc73dacae8688456bdeef72ed6a3030c500ba/README.md)
- [Makefile](https://github.com/pgexperts/pg_plan_filter/blob/c1afc73dacae8688456bdeef72ed6a3030c500ba/Makefile)
- [License](https://github.com/pgexperts/pg_plan_filter/blob/c1afc73dacae8688456bdeef72ed6a3030c500ba/License)

`plan_filter` is the loadable library in pg_plan_filter 1.0.0. It rejects statements or transactions whose estimated plan cost exceeds configured limits. It supports PostgreSQL 14–18 and has no SQL extension to create.

### Core Workflow

```ini
shared_preload_libraries = 'plan_filter'
plan_filter.statement_cost_limit = 100000.0
plan_filter.transaction_cost_limit = 500000.0
plan_filter.filter_select_only = false
```

```sql
SELECT current_setting('plan_filter.module_loaded', true) IS NOT NULL AS active;
```

### Operational Boundaries

Append the library to the existing preload list and restart; `LOAD` is also supported for development and testing. Both cost limits default to zero, disabling enforcement. `plan_filter.statement_cost_limit` caps a single plan and can also reject expensive plain EXPLAIN. `plan_filter.transaction_cost_limit` charges every execution, including repeated prepared statements; rolling back to a savepoint does not refund the budget. Plain EXPLAIN does not consume the transaction budget, while plan-bearing utility commands do.

`plan_filter.filter_select_only` applies both limits only to SELECT when true; SELECT can still have side effects. `plan_filter.module_loaded` is a read-only presence signal. The limit settings require superuser or explicitly granted parameter-setting privileges. Rejection uses SQLSTATE 54001. Cost is an estimate rather than elapsed time, and ordinary roles can lower planner cost settings to evade the threshold. Use this as cooperative protection against accidental load, not as the sole security boundary against hostile sessions.
