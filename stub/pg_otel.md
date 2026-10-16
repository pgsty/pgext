## Usage

Sources:

- [README.md](https://github.com/neteye-platform/pg_otel/blob/490e8dea7860f71a86950902fef426bc6e02edbf/README.md)
- [Cargo.toml](https://github.com/neteye-platform/pg_otel/blob/490e8dea7860f71a86950902fef426bc6e02edbf/Cargo.toml)
- [src/lib.rs](https://github.com/neteye-platform/pg_otel/blob/490e8dea7860f71a86950902fef426bc6e02edbf/src/lib.rs)
- [sql/pg_otel--0.1.0--0.2.0.sql](https://github.com/neteye-platform/pg_otel/blob/490e8dea7860f71a86950902fef426bc6e02edbf/sql/pg_otel--0.1.0--0.2.0.sql)
- [sql/setup.sql](https://github.com/neteye-platform/pg_otel/blob/490e8dea7860f71a86950902fef426bc6e02edbf/sql/setup.sql)
- [pg_otel.control](https://github.com/neteye-platform/pg_otel/blob/490e8dea7860f71a86950902fef426bc6e02edbf/pg_otel.control)

`pg_otel` 0.2.0 exports query and plan-node spans through an OTLP background worker. It is experimental. Load the library at server startup, restart, then enable tracing explicitly; the default minimum duration of -1 disables instrumentation and export.

### Core Workflow

```ini
shared_preload_libraries = 'pg_otel'
pg_otel.otlp_endpoint = 'http://localhost:4317'
pg_otel.otlp_protocol = 'grpc'
pg_otel.min_duration_ms = 100
```

```sql
CREATE EXTENSION pg_otel;
SELECT pg_otel_dropped_spans();
```

### Operational Boundaries

The inspected Cargo manifest targets PostgreSQL 18 and 19, defaulting to 19; these build features are not a production qualification. The library refuses a late `LOAD`. `CREATE EXTENSION` installs the SQL interface; an existing 0.1.0 database needs `ALTER EXTENSION pg_otel UPDATE` after new binaries and restart. No explicit license was found in the inspected repository.

`pg_otel.otlp_endpoint`, `pg_otel.otlp_protocol` and the timeout configure the collector. `pg_otel.queue_size_kb` reserves shared memory at startup; `pg_otel.max_plan_spans` bounds plan detail. A full queue drops spans, reported by `pg_otel_dropped_spans()`. Use transaction-local `pg_otel.traceparent` to avoid associating pooled connections with a previous request.

Instrumentation adds per-node timing overhead. `pg_otel.query_text` defaults to normalized SQL; `raw` may disclose sensitive text. Collector authentication is a privileged setting. A trace can be incomplete or dropped, and a caller can suppress export through an unsampled parent, so this is not a durable audit log.
