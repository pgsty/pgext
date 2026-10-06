## Usage

Sources:

- [README.md](https://github.com/ringerc/postgres_otel_api/blob/7f1b649dd7008de651890996bfc6f671dfac3b2d/README.md)
- [otel_demo_exporter/otel_demo_exporter--0.1.1.sql](https://github.com/ringerc/postgres_otel_api/blob/7f1b649dd7008de651890996bfc6f671dfac3b2d/otel_demo_exporter/otel_demo_exporter--0.1.1.sql)
- [otel_demo_exporter/otel_demo_exporter.c](https://github.com/ringerc/postgres_otel_api/blob/7f1b649dd7008de651890996bfc6f671dfac3b2d/otel_demo_exporter/otel_demo_exporter.c)
- [otel_demo_exporter/otel_demo_exporter.control](https://github.com/ringerc/postgres_otel_api/blob/7f1b649dd7008de651890996bfc6f671dfac3b2d/otel_demo_exporter/otel_demo_exporter.control)

`otel_demo_exporter` is a minimal file exporter for `otel_api` spans. It writes a small JSON projection synchronously and is explicitly a development example, not a production telemetry pipeline.

### Core Workflow

```ini
shared_preload_libraries = 'otel_api,otel_postgres_tracing,otel_demo_exporter'
otel_demo_exporter.output_file = '/var/log/postgresql/otel_spans.jsonl'
```

### Operational Boundaries

Load the API, an instrumentation producer and this exporter; `CREATE EXTENSION` is optional bookkeeping because the SQL has no user-facing objects. Set `otel_demo_exporter.output_file` to a writable server-side file. An empty setting disables output; a relative path is relative to the data directory. Changing it takes effect after a configuration reload and causes backends to reopen the destination.

There is no batching, background dispatch or OTLP/HTTP/gRPC transport. Synchronous output can slow queries and fill disk; arrange access controls, rotation and retention. The older per-directory README uses a superseded module name; the current root README and source establish the `otel_api` loading path.

The project is pre-release: wire and file formats may change. Stock PostgreSQL 14+ supports the GUC/sqlcommenter fallback; protocol-header propagation and some error/parallel-worker behavior require the companion patched PostgreSQL. Do not assume those patches are present on an ordinary server. Append libraries to existing preload entries and restart for cluster-wide enablement. Protect traces because statement text and attributes may contain sensitive data.
