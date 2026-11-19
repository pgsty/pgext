## Usage

Sources:

- [Official source](https://github.com/its-the-shrimp/pg_otel_exporter/blob/v0.0.3/README.md)
- [Official source](https://github.com/its-the-shrimp/pg_otel_exporter/blob/v0.0.3/pg_otel_exporter.control)
- [Official source](https://github.com/its-the-shrimp/pg_otel_exporter/releases/tag/v0.0.3)
- [Official source](https://github.com/its-the-shrimp/pg_otel_exporter/blob/v0.0.3/LICENSE)
- [Official source](https://github.com/its-the-shrimp/pg_otel_exporter/blob/v0.0.3/Cargo.toml)

`pg_otel_exporter` 0.0.3 exports statement traces, server logs and sampled PostgreSQL metrics over OTLP/HTTP. Release artifacts and Cargo features support PostgreSQL 14–19; the README’s broader 13–19 line is not supported by those build inputs.

### Enable telemetry

Add the library to the preload list and restart for the metrics worker and auxiliary-process log capture. Create the extension in each database as a superuser. Set a collector endpoint you control before enabling export.

```ini
shared_preload_libraries = 'pg_otel_exporter'
pg_otel_exporter.endpoint = 'http://127.0.0.1:4318'
```

```sql
CREATE EXTENSION pg_otel_exporter;
BEGIN;
CALL traceparent('00-0af7651916cd43dd8448eb8c8defd531-00f067aa0ba902b7-01');
SELECT 1;
COMMIT;
```

### Signals and configuration

`traceparent(text)` attaches subsequent statements in the transaction to a W3C trace; outside a transaction it has no effect. A SQL comment carrying the same context takes precedence. The context clears on commit, abort or prepared transaction. Executor and utility hooks emit spans, including nested statement relationships; transaction-control statements do not produce spans.

`pg_otel_exporter.traces`, `pg_otel_exporter.metrics` and `pg_otel_exporter.logs` select signals. `pg_otel_exporter.sdk_disabled` disables export. `pg_otel_exporter.traces_sampler_arg`, `pg_otel_exporter.metric_export_interval`, `pg_otel_exporter.log_level` and `pg_otel_exporter.bsp_max_export_batch_size` control sampling and batching. Settings are user-settable as documented; manage endpoint and role policy accordingly.

### Data boundary

Spans contain full SQL statement text, database and role names; log records include server messages. These can disclose query literals and sensitive information to the configured collector. Sampling is telemetry, not a committed audit trail. This project is distinct from the separately catalogued `pg_otel` extension.
