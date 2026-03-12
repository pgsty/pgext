


## Usage

> [plan_filter: filter statements by their execution plans](https://github.com/pgexperts/pg_plan_filter)

The `pg_plan_filter` module tests statements against specific configured criteria before execution, raising an error if the criteria are violated. This allows administrators to prevent execution of certain queries on production databases.

The only criterion currently supported is the maximum allowed estimated cost of the statement plan.

### Configuration

Add to `postgresql.conf`:

```ini
shared_preload_libraries = 'plan_filter'
plan_filter.statement_cost_limit = 100000.0
```

The `statement_cost_limit` must be set to a non-zero value to enable filtering. The default is `0` (no filtering).

### GUC Parameters

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `plan_filter.statement_cost_limit` | float | `0` | Maximum allowed estimated plan cost. `0` disables filtering |
| `plan_filter.limit_select_only` | bool | `false` | When true, only filter `SELECT` statements |

### Examples

Prevent expensive queries globally:

```ini
plan_filter.statement_cost_limit = 100000.0
```

Limit filtering to SELECT statements only (note: SELECT != READONLY, since SELECT can also modify data):

```ini
plan_filter.limit_select_only = true
```

When the module is running with a non-zero `statement_cost_limit`, it will also prevent `EXPLAIN` on expensive queries. Temporarily bypass the filter:

```sql
BEGIN;
SET LOCAL plan_filter.statement_cost_limit = 0;
EXPLAIN SELECT ...;
COMMIT;
```

Override the limit per user:

```sql
ALTER USER can_run_expensive SET plan_filter.statement_cost_limit = 0;
ALTER USER only_cheap_queries SET plan_filter.statement_cost_limit = 10000;
```

### Caveats

The `statement_cost_limit` cancels plans based on their **estimated cost**. The PostgreSQL planner can return cost estimates unrelated to actual query execution time. Be prepared for false positive cancellations and set the limit generously.
