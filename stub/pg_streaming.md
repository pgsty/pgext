## Usage

Sources:

- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/LICENSE)
- [extensions/pg_streaming/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_streaming/pgbrew.toml)
- [Extension control file](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_streaming/pg_streaming.control)
- [Pipeline SQL API and worker initialization](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_streaming/src/lib.rs)
- [Pipeline definition types](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_streaming/src/dsl/types.rs)

`pg_streaming` version `0.3.0` is a declarative stream-processing engine whose pipelines, state, offsets, errors, and metrics remain queryable in PostgreSQL. Pipelines connect inputs to processor chains and outputs, then run under coordinator, executor, and timer background workers.

### Core Workflow

Preload the library, restart PostgreSQL, create the extension in its configured database, and define a pipeline:

```conf
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
