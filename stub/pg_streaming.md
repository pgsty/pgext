## Usage

Sources:

- [Extension control file](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_streaming/pg_streaming.control)
- [Pipeline SQL API and worker initialization](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_streaming/src/lib.rs)
- [Pipeline definition types](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_streaming/src/dsl/types.rs)

`pg_streaming` version `0.2.0` is a declarative stream-processing engine whose pipelines, state, offsets, errors, and metrics remain queryable in PostgreSQL. Pipelines connect inputs to processor chains and outputs, then run under coordinator, executor, and timer background workers.

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
