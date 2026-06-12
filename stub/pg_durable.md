Source: [pg_durable v0.2.2 README](https://github.com/microsoft/pg_durable/blob/v0.2.2/README.md), [User Guide](https://github.com/microsoft/pg_durable/blob/v0.2.2/USER_GUIDE.md), [control file](https://github.com/microsoft/pg_durable/blob/v0.2.2/pg_durable.control), [GUC definitions](https://github.com/microsoft/pg_durable/blob/v0.2.2/src/lib.rs).

## Usage

`pg_durable` runs durable, fault-tolerant SQL workflows inside PostgreSQL. A workflow is built from SQL strings, functions, and DSL operators, then submitted with `df.start()`. State is persisted in PostgreSQL so completed steps are not re-executed after crashes or restarts.

`pg_durable` must be loaded through `shared_preload_libraries`, followed by a PostgreSQL restart. Its background worker connects to the database named by `pg_durable.database` and runs under `pg_durable.worker_role`; upstream defaults are `postgres` and `azuresu`, and the worker role must be a superuser.

### Enable and Grant Access

```sql
CREATE EXTENSION pg_durable;

SELECT df.grant_usage('app_role');
```

`CREATE EXTENSION` does not grant usage to `PUBLIC`. Use `df.grant_usage()` for application roles, and rerun it after extension upgrades so newly added functions are covered. The background worker initializes asynchronously after extension creation; retry if `df.*` calls report that the worker is not initialized yet.

### Start and Monitor Workflows

```sql
SELECT df.start('SELECT ''Hello, durable world!''', 'hello-job');

SELECT *
FROM df.list_instances();

SELECT df.status('a1b2c3d4');
SELECT df.result('a1b2c3d4');
SELECT df.cancel('a1b2c3d4', 'No longer needed');
```

`df.start()` returns an instance ID. Use that ID with `df.status()`, `df.result()`, `df.cancel()`, `df.signal()`, and `df.explain()`.

### Compose SQL Steps

```sql
-- Run one step, name its result, then substitute it in the next step.
SELECT df.start(
  'SELECT 100 AS amount' |=> 'total'
  ~> 'SELECT $total * 2 AS doubled',
  'double-total'
);

-- Branch on a SQL condition.
SELECT df.start(
  'SELECT count(*) > 10 FROM orders'
    ?> 'SELECT ''high volume'''
    !> 'SELECT ''low volume''',
  'order-volume'
);

-- Run in parallel and wait for both branches.
SELECT df.start(
  'SELECT refresh_accounts()' & 'SELECT refresh_orders()',
  'parallel-refresh'
);
```

Core operators are `~>` for sequence, `|=>` for naming a result, `&` for join, `|` for race, `?>` and `!>` for conditional branches, and `@>` for loops.

### Timers, Schedules, and Signals

```sql
SELECT df.start(
  @> (
    df.wait_for_schedule('0 * * * *')
    ~> 'SELECT run_hourly_rollup()'
  ),
  'hourly-rollup'
);

SELECT df.start(
  'SELECT create_invoice()' |=> 'invoice'
  ~> df.wait_for_signal('approval', 86400)
  ~> 'SELECT finalize_invoice($invoice.id)',
  'invoice-approval'
);
```

Useful DSL functions include `df.sleep(seconds)`, `df.wait_for_schedule(cron)`, `df.wait_for_signal(name, timeout)`, `df.signal(id, name, data)`, `df.join()`, `df.race()`, `df.if()`, `df.loop()`, and `df.explain()`.

### Configuration and Caveats

- Required preload: add `pg_durable` to `shared_preload_libraries` and restart PostgreSQL.
- `pg_durable.database` must name the database where the extension is created; workflows are not processed in a different database.
- `pg_durable.worker_role` must exist and be a superuser because the worker bypasses RLS to manage all users' instances.
- Connection-related GUCs include `pg_durable.max_management_connections`, `pg_durable.max_duroxide_connections`, `pg_durable.max_user_connections`, and `pg_durable.execution_acquire_timeout`.
- `df.http()` performs outbound HTTP work and is not included in standard grants unless `df.grant_usage(..., include_http => true)` is used.
- Upstream marks v0.2.2 as early development and not production-ready.
