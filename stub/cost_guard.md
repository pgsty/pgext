## Usage

Sources:

- [Official documentation](https://github.com/vijay750/pg_cost_guard/blob/490986b50898ff3ab1c0625c3c9032f165bee235/README.md)
- [Extension control file](https://github.com/vijay750/pg_cost_guard/blob/490986b50898ff3ab1c0625c3c9032f165bee235/cost_guard.control)
- [Official repository](https://github.com/vijay750/pg_cost_guard)

`cost_guard` Planner-hook guard that rejects statements whose estimated cost exceeds a configured threshold.

### Enablement

Install the files for the intended server, then create `cost_guard` in the target database:

```sql
CREATE EXTENSION cost_guard;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
-- Check query cost without executing
EXPLAIN (FORMAT TEXT, COSTS ON)
SELECT * FROM orders o
JOIN customers c ON o.customer_id = c.id
WHERE o.order_date > '2024-01-01';

-- View cost in JSON format for programmatic analysis
EXPLAIN (FORMAT JSON, COSTS ON)
SELECT * FROM large_table WHERE complex_condition;
```

### Main Objects

The official sources define the extension surface dynamically or through provider tooling; inspect the installed version before granting access.

### Operations and Boundaries

- The verified PostgreSQL-major evidence covers 11, 12, 13, 14, 15, 16, 17, 18; do not infer unlisted majors.
