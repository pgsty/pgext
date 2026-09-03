## Usage

Sources:

- [PGXN 0.2.6 README](https://pgxn.org/dist/pg_durable/0.2.6/README.html)
- [0.2.6 user guide](https://api.pgxn.org/src/pg_durable/pg_durable-0.2.6/USER_GUIDE.md)
- [0.2.6 changelog](https://api.pgxn.org/src/pg_durable/pg_durable-0.2.6/CHANGELOG.md)
- [pg_durable control file](https://api.pgxn.org/src/pg_durable/pg_durable-0.2.6/pg_durable.control)
- [0.2.5 to 0.2.6 upgrade SQL](https://api.pgxn.org/src/pg_durable/pg_durable-0.2.6/sql/pg_durable--0.2.5--0.2.6.sql)

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

### Version 0.2.6 Boundaries

- Upstream source installation and published images support PostgreSQL 17 and 18 with `pgrx` 0.16.1. The extension still requires `shared_preload_libraries`, a restart, and a superuser worker role.
- Upgrades through 0.2.4 and 0.2.5 contain replay-breaking workflow changes. Drain or cancel in-flight JOIN, RACE, loop, and `df.wait_for_schedule()` work before upgrading; the 0.2.4 `df.nodes` key migration also takes an `ACCESS EXCLUSIVE` lock.
- `df.start(..., transaction_mode => 'new')` persists an independent start outside the caller transaction. Cluster-wide admission defaults to two concurrent starts and is controlled by `pg_durable.max_new_transaction_starts` and `pg_durable.new_transaction_start_timeout`.
- Variable substitution is resolved once from left to right in 0.2.6, so token-shaped text introduced by a value is not rescanned. It remains raw SQL substitution; never place untrusted input in `{name}` variables. Named step-result substitution through `$name` performs SQL escaping.
- The undocumented `df.ensure_durofut(text)` helper was removed. Drop or rewrite customer-owned dependent objects before upgrading.
- Re-run `df.grant_usage()` after `ALTER EXTENSION ... UPDATE`, because grants on all functions do not automatically include functions added later.
- `df.http()` and `df.http_multipart()` availability and egress policy are compile-time features. Their restrictions do not sandbox arbitrary SQL or other installed extensions.
- The project remains pre-1.0, and upstream's published Docker images are for evaluation and learning rather than production. Read every adjacent upgrade warning instead of assuming an untested multi-version jump is replay-safe.
