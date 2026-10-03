## Usage

Sources:

- [extensions/pg_swarm/pg_swarm.control](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_swarm/pg_swarm.control)
- [README.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/README.md)
- [extensions/pg_swarm/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_swarm/Cargo.toml)
- [extensions/pg_swarm/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_swarm/src/lib.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/LICENSE)
- [extensions/pg_swarm/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_swarm/pgbrew.toml)
- [extensions/pg_swarm/src/node.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_swarm/src/node.rs)
- [extensions/pg_swarm/src/scheduler.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_swarm/src/scheduler.rs)

`pg_swarm` 0.3.0 registers SQL executors and splits jobs into tasks managed by background workers in `pgswarm`, with status, retries and optional result tables.

### Core Workflow

```conf
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

The control requires superuser installation. Preload and restart; the reviewed node and scheduler connect to the postgres database. Create the extension and executor objects there. `register_executor` records a function receiving task ID, JSONB payload, chunk index and chunk count; `submit_job` queues work. Executors run inside a privileged service, so restrict registration and submission. `pg_swarm.workers`, timeout and retry settings control scheduling. Retried tasks require idempotent handling of external side effects; no cross-node exactly-once guarantee is implied. This is an unsupported proof of concept under Matroid Source Available License 1.0. APIs may change. Version 0.3.0 is the new upgrade baseline: earlier 0.2.0 installations require a rehearsed data migration/recreation, not ordinary ALTER EXTENSION UPDATE. Back up data and dependencies before following that destructive upstream path.
