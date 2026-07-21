## Usage

Sources:

- [Official v0.2.3 README](https://github.com/microsoft/pg_durable/blob/v0.2.3/README.md)
- [v0.2.3 user guide](https://github.com/microsoft/pg_durable/blob/v0.2.3/USER_GUIDE.md)
- [v0.2.3 release notes](https://github.com/microsoft/pg_durable/releases/tag/v0.2.3)
- [v0.2.2 to v0.2.3 upgrade SQL](https://github.com/microsoft/pg_durable/blob/v0.2.3/sql/pg_durable--0.2.2--0.2.3.sql)

`pg_durable` runs durable, fault-tolerant SQL workflows inside PostgreSQL. A workflow is a graph of SQL steps, timers, signals, conditions, and parallel branches submitted with `df.start()`. Execution state is checkpointed in PostgreSQL so completed steps are not repeated after a crash, restart, or retry.

### Enable and Grant Access

Preload the worker, select its database and superuser role if the defaults are unsuitable, then restart PostgreSQL:

```conf
shared_preload_libraries = 'pg_durable'
pg_durable.database = 'postgres'
pg_durable.worker_role = 'postgres'
```

Create the extension in `pg_durable.database` and grant an application login role access:

```sql
CREATE EXTENSION pg_durable;
SELECT df.grant_usage('app_role');
```

The worker role must be a superuser because it manages all users' instances while bypassing row-level security. The role that calls `df.start()` must have `LOGIN`, because workflow SQL is executed through a connection authenticated as that captured role.

### Build and Run a Workflow

```sql
SELECT df.start(
    'SELECT 100 AS amount' |=> 'total'
    ~> 'SELECT $total.amount * 2 AS doubled',
    'double-total'
);
```

`df.start()` returns an instance ID. Use it to monitor or control the run:

```sql
SELECT df.status('a1b2c3d4');
SELECT df.result('a1b2c3d4');
SELECT * FROM df.instance_nodes('a1b2c3d4');
SELECT * FROM df.instance_executions('a1b2c3d4', 20);
SELECT df.cancel('a1b2c3d4', 'No longer needed');
```

### DSL Index

- `~>` sequences steps; `|=>` names a result for `$name`, `$name.column`, or `$name.*` substitution.
- `&` / `df.join()` waits for parallel branches; `|` / `df.race()` keeps the first result.
- `?>` and `!>` / `df.if()` select conditional branches; `@>` / `df.loop()` repeats a graph.
- `df.sleep()`, `df.wait_for_schedule()`, and `df.wait_for_signal()` make waits durable.
- `df.signal()`, `df.wait_for_completion()`, `df.explain()`, and the instance-inspection functions operate on running or stored instances.
- `df.setvar()`, `df.getvar()`, `df.unsetvar()`, and `df.clearvars()` manage per-user variables captured when `df.start()` is called.

### Version 0.2.3 Boundaries

- Fresh v0.2.3 installs place provider objects in `_duroxide`; installations upgraded from 0.2.2 or earlier keep `duroxide`. `df.duroxide_schema()` reports the active schema.
- Graphs deeper than 256 levels or larger than 10,000 nodes are rejected. A condition query returning no rows evaluates as false.
- Re-run `df.grant_usage()` after `ALTER EXTENSION ... UPDATE`, because grants on all functions do not automatically include functions added later.
- Variable `{name}` substitution is raw SQL text substitution; never place untrusted input in such variables. Named step-result substitution through `$name` performs SQL escaping.
- `df.http()` availability and egress policy are compile-time features. Its restrictions do not sandbox arbitrary SQL or other installed extensions.
- Upstream labels the project preview, and the published v0.2.3 Docker images are for evaluation and learning rather than production.
