## Usage

Sources:

- [extensions/pg_swarm/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_swarm/Cargo.toml)
- [extensions/pg_swarm/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_swarm/src/lib.rs)
- [crates/pg_bgworker/src/supervision.rs](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/crates/pg_bgworker/src/supervision.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/CHANGELOG.md)
- [extensions/pg_swarm/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_swarm/pgbrew.toml)
- [extensions/pg_swarm/pg_swarm.control](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_swarm/pg_swarm.control)
- [README.md](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/README.md)
- [extensions/pg_swarm/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_swarm/Cargo.toml)
- [extensions/pg_swarm/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_swarm/src/lib.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/LICENSE)
- [extensions/pg_swarm/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_swarm/pgbrew.toml)
- [extensions/pg_swarm/src/node.rs](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_swarm/src/node.rs)
- [extensions/pg_swarm/src/scheduler.rs](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_swarm/src/scheduler.rs)

`pg_swarm` 0.3.2 registers SQL executors and splits jobs into tasks managed by background workers in `pgswarm`, with status, retries and optional result tables.

### Core Workflow

```ini
shared_preload_libraries = 'pg_swarm'
```

```sql
CREATE EXTENSION pg_swarm;
CREATE FUNCTION swarm_example(bigint, jsonb, integer, integer)
RETURNS jsonb LANGUAGE SQL AS $$ SELECT $2; $$;
SELECT pgswarm.register_executor('example', 'public.swarm_example');
SELECT pgswarm.submit_job('example', '{"message":"hello"}'::jsonb);
SELECT * FROM pgswarm.list_executors();
```

### Operational Boundaries

The control requires superuser installation. Preload and restart; the node and scheduler use `pg_swarm.database` (default postgres). Create the extension and executor objects in that database. `register_executor` records a function receiving task ID, JSONB payload, chunk index and chunk count; `submit_job` queues work. Executors run inside a privileged service, so restrict registration and submission. `pg_swarm.workers`, timeout and retry settings control scheduling. Retried tasks require idempotent handling of external side effects; no cross-node exactly-once guarantee is implied. This is an unsupported proof of concept under Matroid Source Available License 1.0. APIs may change. Version 0.3.0 is the new upgrade baseline: earlier 0.2.0 installations require a rehearsed data migration/recreation, not ordinary ALTER EXTENSION UPDATE. Back up data and dependencies before following that destructive upstream path.

### Current Release and Upgrade

Extension version 0.3.2 uses `pg_swarm.database` for background workers. Create the extension in that database; a worker now waits instead of repeatedly exiting when it is absent. Changing restart-sensitive worker configuration requires a restart. From extension 0.3.0 onward, install matching files and use ALTER EXTENSION UPDATE. Earlier versions still require the migration described above. Withdrawn repository 0.4.0 bottles must be replaced by 0.5.0; upstream bottles do not imply Pigsty package availability.

### Worker Supervision

Version 0.3.2 adds `pgswarm.worker_status()` and the superuser-only `pgswarm.reset_workers()`. Status reports worker name, state, failures, restarts, PID, time of entry into the current state and last failure. `pg_swarm.max_worker_failures` defaults to 10 consecutive failures; 0 retries indefinitely. Restart delays back off from 5 to 60 seconds, and reaching the limit leaves a failed worker idle until reset.

Install matching libraries, restart PostgreSQL to initialize the new shared memory, and update the extension SQL in the configured worker database. Inspect the failure before resetting; a reset does not repair its cause.

```sql
ALTER EXTENSION pg_swarm UPDATE;
SELECT * FROM pgswarm.worker_status();
```
