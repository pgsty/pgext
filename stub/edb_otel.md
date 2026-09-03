## Usage

Sources:

- [Official EDB OTEL documentation](https://www.enterprisedb.com/docs/pg_extensions/otel/)
- [Official installation guide](https://www.enterprisedb.com/docs/pg_extensions/otel/installing/)
- [Official SQL API guide](https://www.enterprisedb.com/docs/pg_extensions/otel/using/)

`edb_otel` exports PostgreSQL metrics and traces to OpenTelemetry endpoints and exposes an API that other extensions can use for centralized telemetry routing.

### Enablement

EDB OTEL 2.0.0 is distributed through EDB repositories. Configure the collector endpoints, preload the library, restart, and create the extension:

```ini
shared_preload_libraries = 'edb_otel'
edb_otel.metrics_endpoint = 'http://otel-collector:4318'
edb_otel.traces_endpoint = 'http://otel-collector:4318'
edb_otel.enable_tracing = true
```

```sql
CREATE EXTENSION edb_otel;
```

Version 2.0 breaks the 1.0 API; update dependent extension code and SQL callers together.

### Report Metrics and Traces

`edb_otel.report_metric` sends a counter, gauge, up/down counter, or histogram value with optional JSON labels. Tracing uses an explicit span lifecycle.

```sql
SELECT edb_otel.report_metric(
  'orders', 'processed', 1, 1::bigint, '{"region":"us-east"}'
);

SELECT edb_otel.tracing_init();
SELECT edb_otel.start_span('checkout', 'validate-order');
```

Use `edb_otel.tracing_set_attribute`, `edb_otel.tracing_end_span`, and `edb_otel.tracing_get_http_trace_context` with the returned span key.

### Operational Boundaries

Metrics pass through a shared-memory queue; when `edb_otel.metrics_queue_size` is exhausted, new metrics are dropped until the worker drains it. Labels over 1,024 bytes or invalid JSON are also dropped and logged. Trace export uses a synchronous `SimpleSpanProcessor` in each backend, so collector latency or failure affects span completion and failed ended spans are not retried.

