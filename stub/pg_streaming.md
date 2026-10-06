## Usage

Sources:

- [extensions/pg_streaming/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_streaming/Cargo.toml)
- [extensions/pg_streaming/src/worker.rs](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_streaming/src/worker.rs)
- [crates/pg_bgworker/src/supervision.rs](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/crates/pg_bgworker/src/supervision.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/CHANGELOG.md)
- [extensions/pg_streaming/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_streaming/pgbrew.toml)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/LICENSE)
- [extensions/pg_streaming/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_streaming/pgbrew.toml)
- [Extension control file](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_streaming/pg_streaming.control)
- [Pipeline SQL API and worker initialization](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_streaming/src/lib.rs)
- [Pipeline definition types](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_streaming/src/dsl/types.rs)

`pg_streaming` version `0.3.2` is a declarative stream-processing engine whose pipelines, state, offsets, errors, and metrics remain queryable in PostgreSQL. Pipelines connect inputs to processor chains and outputs, then run under coordinator, executor, and timer background workers.

### Core Workflow

Preload the library, restart PostgreSQL, create the extension in its configured database, and define a pipeline:

```ini
shared_preload_libraries = 'pg_streaming'
pg_streaming.database = 'postgres'
pg_streaming.worker_count = 2
```

```sql
CREATE EXTENSION pg_streaming;

SELECT pgstreams.create_pipeline(
    'active_orders',
    $json$
    {
      "input": {"table": {"name":"public.order_inbox", "offset_column":"id", "poll":"1s"}},
      "pipeline": {"processors": [{"filter":"value_json->>'active' = 'true'"}]},
      "output": {"table": {"name":"public.active_orders", "mode":"append"}}
    }
    $json$::jsonb
);

SELECT pgstreams.start('active_orders');
SELECT * FROM pgstreams.status();
SELECT * FROM pgstreams.metrics('active_orders');
SELECT pgstreams.stop('active_orders');
```

Lifecycle functions are `create_pipeline`, `update_pipeline`, `drop_pipeline`, `start`, `stop`, and `restart`. Observability includes `status`, `errors`, `late_events`, `metrics`, `lag`, and `trace`; secret and custom-connector registry functions support connector configuration. The DSL covers table, CDC, Kafka, paginated HTTP, OpenDAL, custom inputs/outputs, and processors such as filter, mapping, aggregate, window, join, dedupe, and CEP.

### Operational Notes

The fixed `pgstreams` schema and background-worker registration require superuser installation and server planning. Adding the library or changing worker count requires a restart; reserve enough `max_worker_processes` capacity for one coordinator, the configured executors, and the timer worker. Pipeline expressions and connector definitions execute inside a privileged database service, so restrict pipeline and secret management. Test delivery guarantees, checkpoints, retries, late data, schema evolution, and failure recovery for each connector before production use.

### Version 0.3.0 Boundary

This is an unsupported proof of concept under Matroid Source Available License 1.0. APIs may change. Version 0.3.0 is the new upgrade baseline: earlier 0.2.0 installations require a rehearsed data migration/recreation, not ordinary ALTER EXTENSION UPDATE. Back up data and dependencies before following that destructive upstream path.

The new `call` sink can assume a configured `set_role` for guarded execution. Modbus TCP and Siemens S7 inputs and write sinks can interact with external devices; restrict connector configuration and verify operational authorization before enabling them.

### Current Release and Upgrade

Extension version 0.3.2 uses `pg_streaming.database` for background workers. Create the extension in that database; a worker now waits instead of repeatedly exiting when it is absent. Changing restart-sensitive worker configuration requires a restart. From extension 0.3.0 onward, install matching files and use ALTER EXTENSION UPDATE. Earlier versions still require the migration described above. Withdrawn repository 0.4.0 bottles must be replaced by 0.5.0; upstream bottles do not imply Pigsty package availability.

### Worker Supervision

Version 0.3.2 adds `pgstreams.worker_status()` and the superuser-only `pgstreams.reset_workers()`. Status reports worker name, state, failures, restarts, PID, time of entry into the current state and last failure. `pg_streaming.max_worker_failures` defaults to 10 consecutive failures; 0 retries indefinitely. Restart delays back off from 5 to 60 seconds, and reaching the limit leaves a failed worker idle until reset.

Install matching libraries, restart PostgreSQL to initialize the new shared memory, and update the extension SQL in the configured worker database. Inspect the failure before resetting; a reset does not repair its cause.

```sql
ALTER EXTENSION pg_streaming UPDATE;
SELECT * FROM pgstreams.worker_status();
```
